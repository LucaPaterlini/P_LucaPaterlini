package main

import (
	"reflect"
	"testing"
)

// TestParseQueryStringToDict positive check
func TestParseQueryStringToDict(t *testing.T) {
	str := "text=hello&fun=true&num=5656"
	expected := map[string]interface{}{"text": "hello", "fun": true, "num": int64(5656)}
	response := ParseQueryStringToDict(str)
	if !reflect.DeepEqual(expected, response) {
		t.Error("TestParseQueryStringToDict Dict Parsing Error: expected:", expected, "response:", response)
	}
}
