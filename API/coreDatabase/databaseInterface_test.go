package coreDatabase

import (
	"testing"
)

// TestDatabaseConnect checks the function that connect to database
func TestDatabaseConnect(t *testing.T) {
	// production database
	_, err := TableConnect(false, "perk",[]string{"name", "brand", "value", "created", "expiry"})
	if err != nil {
		t.Error("TestdatabaseConnect Prod Error: " + err.Error())
	}
	/// debug database
	_, err = TableConnect(true, "perk",[]string{"name", "brand", "value", "created", "expiry"})
	if err != nil {
		t.Error("TestdatabaseConnect Debug Error: " + err.Error())
	}
	// check with no keys to check as index
	_, err = TableConnect(true, "perk",[]string{})
	if err != nil {
		t.Error("TestdatabaseConnect Debug Error: " + err.Error())
	}

}
