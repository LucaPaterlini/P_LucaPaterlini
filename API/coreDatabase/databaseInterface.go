// Package coreDatabase contains the primitives to access the database
package coreDatabase

import (
	"github.com/globalsign/mgo"
	"sync"
)

var connections sync.Map

// DatabaseConnect create a connection to the appropriate database depending on debug param
func TableConnect(debug bool, table string, indexkyes []string) (perksTable *mgo.Collection, err error) {
	var location, database string
	if debug {
		location = MONGODBHOSTSDEBUG + "/" + MONGODBDATABASEDEBUG
		database = MONGODBDATABASEDEBUG
	} else {
		location = MONGODBHOSTS + "/" + MONGODBDATABASE
		database = MONGODBDATABASE
	}

	if val, ok := connections.Load(location); ok {
		perksTable = val.(*mgo.Collection)
	} else {
		var mongoSession *mgo.Session
		mongoSession, err = mgo.Dial(location)
		if err != nil {
			return
		}
		perksTable = mongoSession.DB(database).C(table)
	}
	// checking the indexing the test database
	if len(indexkyes) > 0 {

		index := mgo.Index{
			Key:    indexkyes,
			Unique: true,
		}
		err = perksTable.EnsureIndex(index)
	}
	return
}
