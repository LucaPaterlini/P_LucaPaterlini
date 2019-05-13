package coreFunctions

import (
	"../coreDatabase"
	"../schema"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"testing"
)


// TestCreateUpdate checks the function that update or create a new record
func TestCreateUpdate(t *testing.T) {
	// drop the debug database
	table,err := coreDatabase.DatabaseConnect(true)
	if err != nil {
		t.Error("TestCreateUpdate Connection Error: " + err.Error())
		return
	}
	_ = table.DropCollection()
	// executing the Create
	err = CreateUpdate("Save £20 at Tesco",
		"Tesco",20,"2018-03-01 10:15:53","2019-03-01 10:15:53", false,true)

	if err != nil {
		t.Error("TestCreateUpdate: Create:  Error: " + err.Error())
	}
	// checking the effect of the create
	var item schema.Item
	err = table.Find(bson.M{"name": "Save £20 at Tesco"}).One(&item)
	if err != nil {
		t.Error("TestCreateUpdate: CheckCreateFind:  Error: " + err.Error())
	}
	if !("Tesco"==item.Brand && 20==item.Value && "2018-03-01 10:15:53"==item.Created &&
		"2019-03-01 10:15:53"==item.Expiry){
		t.Error("TestCreateUpdate: CheckCreateLook: Error: Cannot retrieve item inserted")
	}
	// Checking the update functionality
	err = CreateUpdate("Save £20 at Tesco",
		"Tesco",8,"2055-03-01 1:15:53","2088-03-01 6:5:1", true,true)
	if err != nil {
		t.Error("TestCreateUpdate: Update:  Error: " + err.Error())
	}

	// Checking the effect of update
	err = table.Find(bson.M{"name": "Save £20 at Tesco"}).One(&item)
	if !("Tesco"==item.Brand && 8==item.Value && "2055-03-01 1:15:53"==item.Created &&
		"2088-03-01 6:5:1"==item.Expiry){
		fmt.Println(item)
		t.Error("TestCreateUpdate: CheckCreateLook: Error: Cannot retrieve item updated")
	}
}
