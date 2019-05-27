package main

import (
	"../coreDatabase"
	"flag"
	"github.com/LucaPaterlini/P_LucaPaterlini/API/endpointsHandler"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	PORT    = 8080
	IPADDR  = "127.0.0.1"
	RTOUT = 10
	WTOUT = 10
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
	// setting the timeout for the service responses
	srv := &http.Server {
		Handler:mux,
		Addr : IPADDR+":"+string(PORT),
		ReadTimeout: RTOUT *time.Second,
		WriteTimeout: WTOUT * time.Second,
	}

	// Configure Logging
	LogFileLocation := os.Getenv("LOG_FILE_LOCATION")
	f, err := os.OpenFile(LogFileLocation, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	// start the service
	log.Fatal(srv.ListenAndServe())
}
