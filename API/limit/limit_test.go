package limit

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

// okHandler its a private http.Handler function that write Ok in the body of the response
func okHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("OK"));err!=nil{log.Print(err.Error())}
}
//emptyHandler is private empty http.Hanler function
func emptyHandler(_ http.ResponseWriter, _*http.Request){

}

// TestVisitors_Limit tests the number of accesses allowed by the Limiter
func TestVisitors_Limit(t *testing.T) {
	//load test
	limit := Visitors{
		CleanUpRefreshTime: 1,
		CleanUpExpiry:      5,
		R:                  2,
		B:                  3,
	}
	ts := httptest.NewServer(limit.Limit(http.HandlerFunc(okHandler)))
	defer ts.Close()
	// test load multiple ports same ip
	for i := 0; i <4 ; i++ {
		if res, err := http.Get(string(ts.URL)); err != nil {
			t.Error(err.Error())
		} else {
			if res.StatusCode != 200 {
				t.Errorf("Expected: 200, got : %d", res.StatusCode)
			}
		}
	}
	// test load single source address

	ts = httptest.NewServer(limit.Limit(http.HandlerFunc(emptyHandler)))
	ctx, _ := context.WithTimeout(context.Background(), 2 * time.Second)
	req, _ := http.NewRequest("GET", string(ts.URL), nil)
	req = req.WithContext(ctx)
	expectedStatus :=  []int{200,200,200,429}
	var receivedStatus []int


	for i := 0; i <4 ; i++ {
		if res, err := http.DefaultClient.Do(req); err != nil {
			t.Error(err.Error())
		} else {
			receivedStatus = append(receivedStatus,res.StatusCode)
		}
	}
	if !reflect.DeepEqual(expectedStatus,receivedStatus) {
		t.Errorf("expected %v ,received %v",expectedStatus,receivedStatus)
	}




}
