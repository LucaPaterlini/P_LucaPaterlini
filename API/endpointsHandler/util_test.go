package endpointsHandler

import (
	"errors"
	"testing"
)
// TestComposeJson check the compose json function on error options
func TestComposeJson(t *testing.T) {
	// checking composition of json for positive response
	param := map[string]interface{}{"data": "hello"}
	expected := `{
	"err": false,
	"data": {
		"data": "hello"
	}
}`
	response := ComposeJson(param, nil)
	if response != expected {
		t.Errorf("TestComposeJson Error Nill expected %s and got %s instead", response, expected)
	}
	// checking composition of json for an answer with error

	expected=`{
	"err": true,
	"data": {
		"errMsg": "funny error message"
	}
}`
	response = ComposeJson(param, errors.New("funny error message"))
	if response != expected {
		t.Errorf("TestComposeJson Error Msg expected %s and got %s instead", response, expected)
	}
}

// TestCheckKeys CheckKeys positive and negative
func TestCheckKeys(t *testing.T) {
	response :=CheckKeys( map[string]interface{}{"hi":"item"}, []string{"hi"})
	if !response{
		t.Error("TestCheckKeys error on positive check")
	}
	response =CheckKeys( map[string]interface{}{"hi":"item"}, []string{"hello"})
	if response{
		t.Error("TestCheckKeys error on negative check")
	}


}
