package main

import (
	"flag"
	"github.com/LucaPaterlini/P_LucaPaterlini/API/coreDatabase"
	"github.com/LucaPaterlini/P_LucaPaterlini/API/endpointsHandler"
	"log"
	"net/http"
)

func main() {
	flag.Parse()
	// database connection
	perksTable, err := coreDatabase.DatabaseConnect(false)
	if err!=nil {
		log.Println(err)
		return
	}

	// adding the handlers
	m := endpointsHandler.HandlerStruct{Collection: perksTable}
	http.HandleFunc("/createupdate", m.HandlerCreateUpdate)
	http.HandleFunc("/retrieve", m.HandlerRetrieve)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
