package coreFunctions

import (
	"../coreDatabase"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"strconv"
	"testing"
	"time"
)

// TestCreateUpdate checks the function that update or create a new record
func TestCreateUpdate(t *testing.T) {
	// drop the debug database
	table, err := coreDatabase.TableConnect(true, "perk", []string{"name", "brand", "value", "created", "expiry"})
	if err != nil {
		t.Error("TestCreateUpdate Connection Error: " + err.Error())
		return
	}
	_, err = table.RemoveAll(bson.M{})

	// 1 executing the Create
	err = CreateUpdate("Save £20 at Tesco",
		"Tesco", 20, "2018-03-01 10:15:53", "2019-03-01 10:15:53", false, table)
	if err != nil {
		t.Error("TestCreateUpdate: Create:  Error: " + err.Error())
	}
	//initializing first insert tCreated and tExpiry
	var tCreated, tExpiry time.Time
	layout := "2006-01-02 15:04:05"
	if tCreated, err = time.Parse(layout, "2018-03-01 10:15:53"); err != nil {
		t.Error("TestCreateUpdate StrToTimeConv Insert Error:", err.Error())
	}
	if tExpiry, err = time.Parse(layout, "2019-03-01 10:15:53"); err != nil {
		t.Error("TestCreateUpdate StrToTimeConv Insert Error:", err.Error())
	}

	// checking the effect of the create
	var item Item
	err = table.Find(bson.M{"name": "Save £20 at Tesco"}).One(&item)
	if err != nil {
		t.Error("TestCreateUpdate: CheckCreateFind:  Error: " + err.Error())
	}

	if !("Tesco" == item.Brand && 20 == item.Value && tCreated == item.Created &&
		tExpiry == item.Expiry) {
		t.Error("TestCreateUpdate: CheckCreateLook: Error: Cannot retrieve item inserted")
	}
	// checking the response of 2 insertions of the same element
	// create 2
	err = CreateUpdate("Save £20 at Tesco",
		"Tesco", 20, "2018-03-01 10:15:53", "2019-03-01 10:15:53", false, table)
	if err != nil {
		if !mgo.IsDup(err) {
			t.Error(err.Error())
		}
	}
	//
	//// Checking the update functionality
	err = CreateUpdate("Save £20 at Tesco",
		"Tesco", 8, "2055-03-01 1:15:53", "2088-03-01 06:05:01", true, table)
	if err != nil {
		t.Error("TestCreateUpdate: Update:  Error: " + err.Error())
	}

	//initializing first update tCreated and tExpiry
	if tCreated, err = time.Parse(layout, "2055-03-01 1:15:53"); err != nil {
		t.Error("TestCreateUpdate StrToTimeConv Update Error:", err.Error())
	}
	if tExpiry, err = time.Parse(layout, "2088-03-01 06:05:01"); err != nil {
		t.Error("TestCreateUpdate StrToTimeConv Update Error:", err.Error())
	}
	// Checking the effect of update
	err = table.Find(bson.M{"name": "Save £20 at Tesco"}).One(&item)

	if !("Tesco" == item.Brand && 8 == item.Value && tCreated == item.Created &&
		tExpiry == item.Expiry) {
		t.Error("TestCreateUpdate: CheckCreateLook: Error: Cannot retrieve item updated")
	}
	//  clean the debug database
	_ = table.DropCollection()
}

// TestRetrieve checks the retrive against a suite of test queries
func TestRetrieve(t *testing.T) {
	// drop the debug database
	table, err := coreDatabase.TableConnect(true, "perk", []string{"name", "brand", "value", "created", "expiry"})
	if err != nil {
		t.Error("TestRetrieve Connection Error: " + err.Error())
		return
	}
	_ = table.DropCollection()
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
	// query: filter by name
	FullList, err := Retrieve("Tesco0", false, time.Time{}, table)
	if err != nil {
		t.Error("TestRetrieve Retrieve Error: " + err.Error())
		return
	}
	compareItem := Item{Name: "Tesco0", Brand: "Tesco", Value: 1,
		Created: time.Date(2018, 0, 4, 0, 0, 0, 0, time.UTC),
		Expiry:  time.Date(2018, 2, 4, 0, 0, 0, 0, time.UTC),
	}
	if FullList[0] != compareItem {
		t.Error("TestRetrieve Retrieve check Item Error")
	}
	// query: only active offers
	layout := "2006-01-02T15:04:05.000Z"
	str := "2019-05-25T11:45:26.371Z"
	startTime, _ := time.Parse(layout, str)
	FullList, err = Retrieve("*", true, startTime, table)
	if err != nil {
		t.Error("TestRetrieve Retrieve Active Error: " + err.Error())
		return
	}
	if len(FullList) != 2 {
		t.Error("TestRetrieve Retrieve check Active List")
	}

	// query: all the offers
	FullList, err = Retrieve("*", false, time.Time{}, table)
	if err != nil {
		t.Error("TestRetrieve Retrieve Active Error: " + err.Error())
		return
	}
	if len(FullList) != 4 {
		t.Error("TestRetrieve Retrieve All List")
	}
	// cleanup
	_ = table.DropCollection()
}
