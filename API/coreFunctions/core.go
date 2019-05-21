// Package coreFunctions contains the core functions
package coreFunctions

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)
import "../schema"
import "../coreDatabase"

// CreateUpdate create or update the collection with the new record depending on the update param
// using the database and collection for production if debug is false otherwise it utilizes the
// settings for debug
func CreateUpdate(name,brand string, value int64,created,expiry string,update bool,debug bool) (err error) {
	var perksTable *mgo.Collection
	var tCreated,tExpiry time.Time
	// establishing the connection
	perksTable, err = coreDatabase.DatabaseConnect(debug)
	if err != nil { return}
	// converting created and expiry from str to time
	layout := "2006-01-02 15:04:05"
	if tCreated, err = time.Parse(layout, created); err!=nil{return }
	if tExpiry, err = time.Parse(layout, expiry); err!=nil{return }
	// add the new item
	item := bson.M{"name": name,
		"brand": brand,
		"value": value,
		"created": tCreated,
		"expiry":tExpiry,
	}
	if update {
		err = perksTable.Update(bson.M{"name": name},bson.M{"$set":item})

	} else{
		err = perksTable.Insert(&item)
	}
	return
}

// Retrieve returns a list of records filtered by the name and or if they are active
func Retrieve(name string,active bool, debug bool) (FullList []schema.Item, err error) {
	// prepare the query
	query:=bson.M{}
	if name!=""{query["name"]=name}
	if active{
		now := time.Now()
		query["created"] = bson.M{"$lt":now}
		query["expiry"] =  bson.M{"$gt":now}
	}
	var perksTable *mgo.Collection
	// establishing the connection
	perksTable, err = coreDatabase.DatabaseConnect(debug)
	if err != nil { return}
	// retrieve the data
	err = perksTable.Find(query).All(&FullList)
	return
}