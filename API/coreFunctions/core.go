// Package coreFunctions contains the core functions
package coreFunctions

import (
	"github.com/LucaPaterlini/P_LucaPaterlini/API/schema"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

// CreateUpdate create or update the collection with the new record depending on the update param
// using the database and collection for production if debug is false otherwise it utilizes the
// settings for debug
func CreateUpdate(name, brand string, value int64, created, expiry string, update bool, collection *mgo.Collection) (err error) {
	var tCreated, tExpiry time.Time

	// converting created and expiry from str to time
	layout := "2006-01-02 15:04:05"
	if tCreated, err = time.Parse(layout, created); err != nil {
		return
	}
	if tExpiry, err = time.Parse(layout, expiry); err != nil {
		return
	}
	// preparing the id of the item
	item := bson.M{
		"name":    name,
		"brand":   brand,
		"value":   value,
		"created": tCreated,
		"expiry":  tExpiry,
	}

	if update {
		err = collection.Update(bson.M{"name": name}, bson.M{"$set": item})
	}
	if !update || (err != nil && err.Error() == "not found") {
		err = collection.Insert(&item)

	}
	return
}

// Retrieve returns a list of records filtered by the name and or if they are active
func Retrieve(name string, active bool, starttime time.Time, collection *mgo.Collection) (FullList []schema.Item, err error) {
	// debug fmt.Println("Parameters",name,active,starttime,debug)
	// prepare the query
	query := bson.M{}
	if name != "*" {
		query["name"] = name
	}
	if active {
		query["created"] = bson.M{"$lt": starttime}
		query["expiry"] = bson.M{"$gt": starttime}
	}
	// retrieve the data
	err = collection.Find(query).All(&FullList)
	//debug :fmt.Println(FullList)
	return
}
