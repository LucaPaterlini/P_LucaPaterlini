// Package coreFunctions contains the core functions
package coreFunctions

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)
import "../schema"
import "../coreDatabase"

// CreateUpdate create or update the collection with the new record depending on the update param
// using the database and collection for production if debug is false otherwise it utilizes the
// settings for debug
func CreateUpdate(name,brand string, value int64,created,expiry string,update bool,debug bool) (err error) {
	var perksTable *mgo.Collection
	// establishing the connection
	perksTable, err = coreDatabase.DatabaseConnect(debug)
	if err != nil { return}
	// add the new item
	item := bson.M{"name": name,
		"brand": brand,
		"value": value,
		"created": created,
		"expiry":expiry,
	}
	if update {
		err = perksTable.Update(bson.M{"name": name},bson.M{"$set":item})

	} else{
		err = perksTable.Insert(&item)
	}
	return
}

func Retrieve(name string) (FullList []schema.Item, err error, debug bool) {
	var perksTable *mgo.Collection
	// establishing the connection
	perksTable, err = coreDatabase.DatabaseConnect(debug)
	if err != nil { return}
	// retrieve the data
	err = perksTable.Find(bson.M{"name": name}).All(&FullList)
	return
}


func List() (FullList []schema.Item, err error, debug bool) {
	var perksTable *mgo.Collection
	// establishing the connection
	perksTable, err = coreDatabase.DatabaseConnect(debug)
	if err != nil { return}
	// retrieve the data
	err = perksTable.Find(bson.M{}).All(&FullList)
	return
}
