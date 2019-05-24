package main

import (
	"../endpointsHandler"
	"flag"
	"log"
	"net/http"
)

func main() {
	flag.Parse()
	m := endpointsHandler.HandlerStruct{Debug: false}
	http.HandleFunc("/createupdate", m.HandlerCreateUpdate)
	http.HandleFunc("/retrieve", m.HandlerRetrieve)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
