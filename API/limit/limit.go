// Package limit provides the primitives to limit the access for each second
// from the same address (comb ipv4/ipv6+port)
package limit

import (
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

// visitor is used as value in the map value field register
// as attributes it have the limiter and the last access
// of the referenced source address
type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// Visitors is used to hold the limiter of the access of each user
// having in common the Refresh rate and the Bucket size for
// all the Visitors of the same
type Visitors struct {
	register           map[string]*visitor
	mtx                sync.RWMutex
	CleanUpRefreshTime int64
	CleanUpExpiry      int64
	R                  rate.Limit
	B                  int
}

// addVisitor is a private method that adds the new entry to
// the register that holds the visitor objects
func (v *Visitors) addVisitor(addr string) *rate.Limiter {
	limiter := rate.NewLimiter(v.R, v.B)
	v.mtx.Lock()
	v.register[addr] = &visitor{limiter, time.Now()}
	v.mtx.Unlock()
	return limiter
}

// getVisitor is a private method that retrieve or create
// the visitor object from the register of accesses
func (v *Visitors) getVisitor(addr string) *rate.Limiter {
	v.mtx.Lock()
	item, exists := v.register[addr]

	if !exists {
		v.mtx.Unlock()
		return v.addVisitor(addr)
	}
	item.lastSeen = time.Now()
	v.mtx.Unlock()
	return item.limiter
}

// cleanupVisitors its a private method that erase the expired
// access logs
func (v *Visitors) cleanupVisitors() {
	for {
		time.Sleep(time.Duration(v.CleanUpRefreshTime) * time.Minute)
		v.mtx.Lock()
		for ip, visitor := range v.register {
			if time.Now().Sub(visitor.lastSeen) > time.Duration(v.CleanUpExpiry)*time.Minute {
				delete(v.register, ip)
			}
		}
		v.mtx.Unlock()
	}
}

// Limit is the public method that act as middleware for http.Handlers
// limiting the number of access to service provided by the handler
func (v *Visitors) Limit(next http.Handler) http.Handler {
	// initialization
	//go v.cleanupVisitors()
	v.register = make(map[string]*visitor)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := v.getVisitor(r.RemoteAddr)
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}