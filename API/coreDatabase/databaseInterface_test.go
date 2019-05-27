package coreDatabase

import (
	"testing"
)

// TestDatabaseConnect checks the function that connect to database
func TestDatabaseConnect(t *testing.T) {
	_, err := TableConnect(false, "perk",[]string{"name", "brand", "value", "created", "expiry"})
	if err != nil {
		t.Error("TestdatabaseConnect Prod Error: " + err.Error())
	}
	_, err = TableConnect(true, "perk",[]string{"name", "brand", "value", "created", "expiry"})
	if err != nil {
		t.Error("TestdatabaseConnect Debug Error: " + err.Error())
	}
}
