package endpointsHandler

import (
	"../coreDatabase"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

//TestHandler_HandlerCreateUpdate check all the http status codes
// using a clean debug debug database and clening up after usage
// 500 has not been tested
func TestHandler_HandlerCreateUpdate(t *testing.T) {
	// drop the debug database
	table, err := coreDatabase.DatabaseConnect(true)
	if err != nil {
		t.Error("TestCreateUpdate Connection Error: " + err.Error())
		return
	}
	_, err = table.RemoveAll(bson.M{})

	// adding the debug param
	m := HandlerStruct{Debug: true}
	params := "?name=Save £20 at Tesco&brand=Tesco&value=20&createdAt=2018-03-01 10:15:53&expiry=2019-03-01 10:15:53&update=true"
	request, _ := http.NewRequest("GET", "/createupdate"+params, nil)

	// request 1 expected positive
	response := httptest.NewRecorder()
	m.HandlerCreateUpdate(response, request)
	if http.StatusOK != response.Code {
		t.Errorf("Response code expected: 200, received: %d", response.Code)
	}
	expected := `{
	"err": false,
	"data": {}
}`
	received := response.Body.String()
	if received != expected {
		t.Errorf("Unexpected Response: expected %s, received %s", expected, received)
	}
	// request 2 expected positive
	response = httptest.NewRecorder()
	m.HandlerCreateUpdate(response, request)
	if http.StatusOK != response.Code {
		t.Errorf("Response code expected: 200, received: %d", response.Code)
	}
	received = response.Body.String()
	if received != expected {
		t.Errorf("Unexpected Response: expected\n %s \nreceived\n %s\n", received, expected)
	}
	// request 3 should fail
	params = "?name=Save £20 at Tesco&brand=Tesco&value=20&createdAt=2018-03-01 10:15:53&expiry=2019-03-01 10:15:53&update=false"
	request, _ = http.NewRequest("GET", "/createupdate"+params, nil)
	response = httptest.NewRecorder()
	expected = `{
	"err": true,
	"data": {
		"errMsg": "/createupdate: Entry Duplicate"
	}
}`

	m.HandlerCreateUpdate(response, request)
	if http.StatusBadRequest != response.Code {
		t.Errorf("Response code expected: 405, received: %d", response.Code)
	}
	received = response.Body.String()
	if received != expected {
		t.Errorf("Unexpected Response: expected %s, received %s", expected, received)
		fmt.Println(received)
	}

	// request 4 response to wrong params
	params = "?name=Save £20 at Tesco&brand=Tesco&value=20&createdAt=2018-03-01 10:15:53&expiry=2019-03-01 10:15:53&update=banan"
	request, _ = http.NewRequest("GET", "/createupdate"+params, nil)
	response = httptest.NewRecorder()
	m.HandlerCreateUpdate(response, request)
	if http.StatusBadRequest != response.Code {
		t.Errorf("Response code expected: 400, received: %d", response.Code)
	}
	expected = `{
	"err": true,
	"data": {
		"errMsg": "/createupdate: wrong set of input parameters"
	}
}`
	received = response.Body.String()
	if received != expected {
		t.Errorf("Unexpected Response: expected %s, received %s", expected, received)
		fmt.Println(received)
	}

	// checking that only GET is a valid method
	for _, method := range []string{"HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE"} {
		request, _ = http.NewRequest(method, "/createupdate"+params, nil)
		response = httptest.NewRecorder()
		m.HandlerCreateUpdate(response, request)
		if http.StatusMethodNotAllowed != response.Code {
			t.Error("Wrong forbidden method response")
		}

	}
	// cleanUp
	_, err = table.RemoveAll(bson.M{})
}

// TestHandlerStruct_HandlerRetrieve checks the handler Retrieve
// endpoint respnse status and bodies using a clean database
// 500 has not been tested
func TestHandlerStruct_HandlerRetrieve(t *testing.T) {
	// drop the debug database
	table, err := coreDatabase.DatabaseConnect(true)
	if err != nil {
		t.Error("TestCreateUpdate Connection Error: " + err.Error())
		return
	}
	_, err = table.RemoveAll(bson.M{})
	// Insert test Items
	for i := 0; i < 4; i++ {
		dateCreated := time.Date(2018, time.Month(i), 4, 0, 0, 0, 0, time.UTC)
		dateExpiry := time.Date(2018+i, time.Month(i+2), 4, 0, 0, 0, 0, time.UTC)
		item := bson.M{"name": "Tesco" + strconv.Itoa(i),
			"brand":   "Tesco",
			"value":   1 + i*2,
			"created": dateCreated,
			"expiry":  dateExpiry,
		}
		err = table.Insert(&item)
	}
	// adding the debug param
	m := HandlerStruct{Debug: true}
	// request 1 check specific name
	params := "?name=Tesco0&active=false"
	request, _ := http.NewRequest("GET", "/retrieve"+params, nil)
	response := httptest.NewRecorder()
	m.HandlerRetrieve(response, request)
	received := response.Body.String()

	if http.StatusOK != response.Code {
		t.Errorf("Response code expected: 200, received: %d", response.Code)
	}
	expected := `{
	"err": false,
	"data": {
		"list": [
			{
				"name": "Tesco0",
				"brand": "Tesco",
				"value": 1,
				"created": "2017-12-04T00:00:00Z",
				"expiry": "2018-02-04T00:00:00Z"
			}
		]
	}
}`
	if received != expected {
		t.Errorf("Unexpected Response: expected %s, received %s", expected, received)
	}

	// request 2  check active perks
	params = "?name=*&active=true"
	request, _ = http.NewRequest("GET", "/retrieve"+params, nil)
	response = httptest.NewRecorder()
	m.HandlerRetrieve(response, request)
	if http.StatusOK != response.Code {
		t.Errorf("Response code expected: 200, received: %d", response.Code)
	}
	expected = `{
	"err": false,
	"data": {
		"list": [
			{
				"name": "Tesco2",
				"brand": "Tesco",
				"value": 5,
				"created": "2018-02-04T00:00:00Z",
				"expiry": "2020-04-04T00:00:00Z"
			},
			{
				"name": "Tesco3",
				"brand": "Tesco",
				"value": 7,
				"created": "2018-03-04T00:00:00Z",
				"expiry": "2021-05-04T00:00:00Z"
			}
		]
	}
}`
	received = response.Body.String()
	if received != expected {
		t.Errorf("Unexpected Response: expected %s, received %s", received, expected)
	}
	// request 3 check any perk
	params = "?name=*&active=false"
	request, _ = http.NewRequest("GET", "/retrieve"+params, nil)
	response = httptest.NewRecorder()
	m.HandlerRetrieve(response, request)
	received = response.Body.String()

	if http.StatusOK != response.Code {
		t.Errorf("Response code expected: 200, received: %d", response.Code)
	}
	expected = `{
	"err": false,
	"data": {
		"list": [
			{
				"name": "Tesco0",
				"brand": "Tesco",
				"value": 1,
				"created": "2017-12-04T00:00:00Z",
				"expiry": "2018-02-04T00:00:00Z"
			},
			{
				"name": "Tesco1",
				"brand": "Tesco",
				"value": 3,
				"created": "2018-01-04T00:00:00Z",
				"expiry": "2019-03-04T00:00:00Z"
			},
			{
				"name": "Tesco2",
				"brand": "Tesco",
				"value": 5,
				"created": "2018-02-04T00:00:00Z",
				"expiry": "2020-04-04T00:00:00Z"
			},
			{
				"name": "Tesco3",
				"brand": "Tesco",
				"value": 7,
				"created": "2018-03-04T00:00:00Z",
				"expiry": "2021-05-04T00:00:00Z"
			}
		]
	}
}`
	if received != expected {
		t.Errorf("Unexpected Response: expected %s, received %s", expected, received)
	}
	// request 4 check wrong parameters
	params = "?name=Tesco0"
	request, _ = http.NewRequest("GET", "/retrieve"+params, nil)
	response = httptest.NewRecorder()
	m.HandlerRetrieve(response, request)
	received = response.Body.String()

	if http.StatusBadRequest != response.Code {
		t.Errorf("Response code expected: 400, received: %d", response.Code)
	}
	expected = `{
	"err": true,
	"data": {
		"errMsg": "/retrieve: wrong set of input parameters"
	}
}`
	if received != expected {
		t.Errorf("Unexpected Response: expected %s, received %s", expected, received)
	}

	// checking that only GET is a valid method
	for _, method := range []string{"HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE"} {
		request, _ = http.NewRequest(method, "/retrieve", nil)
		response = httptest.NewRecorder()
		m.HandlerCreateUpdate(response, request)
		if http.StatusMethodNotAllowed != response.Code {
			t.Error("Wrong forbidden method response")
		}

	}
	// cleanup
	_, _ = table.RemoveAll(bson.M{})
}
