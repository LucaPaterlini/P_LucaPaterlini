package coreDatabase

import (
	"testing"
)


// TestDatabaseConnect checks the function that connect to database
func TestDatabaseConnect(t *testing.T) {
	_,err := DatabaseConnect(false)
	if err != nil {
		t.Error("TestdatabaseConnect Prod Error: " + err.Error())
	}
	_,err = DatabaseConnect(true)
	if err != nil {
		t.Error("TestdatabaseConnect Debug Error: " + err.Error())
	}
}
