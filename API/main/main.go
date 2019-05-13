package main

import (
	"../endpointsHandler"
	"../schema"
	"flag"
	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	gocache "github.com/pmylund/go-cache"
	"github.com/valyala/fasthttp"
	"log"
	"strconv"
	"time"
)

var (
	addr = flag.String("addr",schema.IPADDR+":"+strconv.Itoa(schema.PORT),"TCP address to listen to")
	compress = flag.Bool("compress",false, "Whether to enable transparent response compression ")
	cache = gocache.New(1*time.Minute, 3*time.Minute)
)


func main(){

	// Cors Handler
	withCors := cors.DefaultHandler()

	flag.Parse()
	h := withCors.CorsMiddleware(logPanics(routingHandler))
	if *compress {
		h = fasthttp.CompressHandler(h)
	}
	if err := fasthttp.ListenAndServe(*addr,h);err != nil{
		log.Fatalf("Errror in ListenAndServer: %s",err)
	}

}


// routing
func routingHandler (ctx *fasthttp.RequestCtx){
	ctx.SetContentType("text/json; charset=utf-8")

	switch string(ctx.Path()) {
	//Call n°0
	case "/create":
		middlewareEndpoint(ctx,endpointsHandler.HandlerCreate)

	//Call n°1
	case "/update":
		middlewareEndpoint(ctx,endpointsHandler.HandlerUpdate)

	//Call n°2
	case "/retrive":
		middlewareEndpoint(ctx,endpointsHandler.HandlerRetrieve)

	//Call n°3
	case "/all":
		middlewareEndpoint(ctx,endpointsHandler.HandlerList)

	default:
		ctx.Error(schema.ERRPATH,fasthttp.StatusNotFound)
	}
}
