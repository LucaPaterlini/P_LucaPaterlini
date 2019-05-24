package main

import (
	"../coreDatabase"
	"../endpointsHandler"
	"flag"
	"log"
	"net/http"
)

func main() {
	flag.Parse()
	// database connection
	perkstable, err := coreDatabase.DatabaseConnect(false)
	if err!=nil {
		log.Println(err)
		return
	}

	// adding the handlers
	m := endpointsHandler.HandlerStruct{Collection: perkstable}
	http.HandleFunc("/createupdate", m.HandlerCreateUpdate)
	http.HandleFunc("/retrieve", m.HandlerRetrieve)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
