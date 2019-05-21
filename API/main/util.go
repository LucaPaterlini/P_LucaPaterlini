package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	gocache "github.com/pmylund/go-cache"
	"github.com/valyala/fasthttp"
	"log"
	"net/url"
	"strconv"
	"strings"
)

// logPanics log all the panic messages
func logPanics(function fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", ctx.RemoteAddr(), x)
			}
		}()
		function(ctx)
	}
}

// middlewareEndpoint ac as middleware between the request and the execution of the andpoint
// and the handling of the response
func middlewareEndpoint(ctx *fasthttp.RequestCtx, f func(map[string]interface{}) string) {
	query, err := url.QueryUnescape(ctx.QueryArgs().String())
	if err == nil {
		d := ParseQueryStringToDict(query)
		key := string(ctx.Path()) + ctx.QueryArgs().String()
		if _, found := cache.Get(key); !found {
			cache.Set(key, f(d), gocache.DefaultExpiration)
		}
		response, _ := cache.Get(key)
		_, err = fmt.Fprint(ctx, fmt.Sprintf("%v", response))
	}
	if err == nil {
		ctx.SetStatusCode(fasthttp.StatusOK)
	} else {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}
}

// Parsing the query to a dictionary
func ParseQueryStringToDict(a string) map[string]interface{} {
	d := make(map[string]interface{})
	if len(a) < 1 {
		return d
	}
	for _, t := range strings.Split(a, "&") {
		g := strings.Split(t, "=")
		if len(g) < 2 {
			continue
		}
		k, v := g[0], g[1]
		if govalidator.IsInt(v) {
			d[k], _ = strconv.ParseInt(v, 10, 64)
		} else if govalidator.IsFloat(v) {
			d[k], _ = strconv.ParseFloat(v, 64)
		} else if b, err := govalidator.ToBoolean(v); err == nil {
			d[k] = b
		} else {
			d[k] = v
		}
	}
	return d
}
