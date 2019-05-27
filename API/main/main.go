package main

import (
	"../coreDatabase"
	"flag"
	"github.com/LucaPaterlini/P_LucaPaterlini/API/endpointsHandler"
	"log"
	"net/http"
)

const (
	PORT    = 8080
	IPADDR  = "127.0.0.1"
	ERRPATH = "{\"error\":\"Unsupported Path\"}"
)

func main() {
	flag.Parse()
	// database connection
	perksTable, err := coreDatabase.TableConnect(false,"perk",[]string{"name", "brand", "value", "created", "expiry"})
	if err!=nil {
		log.Println(err)
		return
	}

	// adding the handlers
	m := endpointsHandler.HandlerStruct{Collection: perksTable}
	mux := http.NewServeMux()
	mux.HandleFunc("/createupdate", m.HandlerCreateUpdate)
	mux.HandleFunc("/retrieve", m.HandlerRetrieve)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
