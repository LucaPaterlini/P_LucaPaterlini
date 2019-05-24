// Package coreDatabase contains the primitives to access the database
package coreDatabase

import (
	"github.com/LucaPaterlini/P_LucaPaterlini/API/schema"
	"github.com/globalsign/mgo"
)

var connections map[string]*mgo.Collection

// DatabaseConnect create a connection to the appropriate database depending on debug param
func DatabaseConnect(debug bool) (perksTable *mgo.Collection, err error) {
	var location, database string
	if debug {
		location = schema.MONGODBHOSTSDEBUG + "/" + schema.MONGODBDATABASEDEBUG
		database = schema.MONGODBDATABASEDEBUG
	} else {
		location = schema.MONGODBHOSTS + "/" + schema.MONGODBDATABASE
		database = schema.MONGODBDATABASE
	}

	if val, ok := connections[location]; ok {
		perksTable = val
	} else {
		var mongoSession *mgo.Session
		mongoSession, err = mgo.Dial(location)
		if err != nil {
			return
		}
		perksTable = mongoSession.DB(database).C("perks")
	}
	// checking the indexing
	// setting the index on the test database

	index := mgo.Index{
		Key:    []string{"name", "brand", "value", "created", "expiry"},
		Unique: true,
	}
	err = perksTable.EnsureIndex(index)

	return
}
