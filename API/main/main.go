package main

import (
	"../coreDatabase"
	"../limit"
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
	ReadTimeout = 10
	WriteTimeout = 10
	LimitCleanUpRefreshTime = 1
	LimitClenaUpExpiry = 5
	LimitRefresh  = 2
	LimitBucket    = 3


)

func main() {
	flag.Parse()
	// database connection
	perksTable, err := coreDatabase.TableConnect(false,"perk",[]string{"name", "brand", "value", "created", "expiry"})
	if err!=nil {
		log.Println(err)
		return
	}

	// adding the handlers and them paths
	m := endpointsHandler.HandlerStruct{Collection: perksTable}
	mux := http.NewServeMux()
	mux.HandleFunc("/createupdate", m.HandlerCreateUpdate)
	mux.HandleFunc("/retrieve", m.HandlerRetrieve)
	// instantiating the Visitor class witch attributed are used to rate the access
	// to the endpoints for each source
	l := limit.Visitors{
		CleanUpRefreshTime: LimitCleanUpRefreshTime,
		CleanUpExpiry:      LimitClenaUpExpiry,
		R:                  LimitRefresh,
		B:                  LimitBucket,
	}


	// setting the timeout for the service responses and the limit control
	srv := &http.Server {
		// appending the middleware
		Handler:l.Limit(mux),
		Addr : IPADDR+":"+string(PORT),
		ReadTimeout: ReadTimeout *time.Second,
		WriteTimeout: WriteTimeout * time.Second,
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
