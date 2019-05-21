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
	addr     = flag.String("addr", schema.IPADDR+":"+strconv.Itoa(schema.PORT), "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression ")
	cache    = gocache.New(1*time.Minute, 3*time.Minute)
)

func main() {

	// Cors Handler
	withCors := cors.DefaultHandler()

	flag.Parse()
	h := withCors.CorsMiddleware(logPanics(routingHandler))
	if *compress {
		h = fasthttp.CompressHandler(h)
	}
	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
		log.Fatalf("Errror in ListenAndServer: %s", err)
	}

}

// routing
func routingHandler(ctx *fasthttp.RequestCtx) {
	// support only Get Requests
	if !ctx.IsGet() {
		ctx.Error(schema.ERRPATH, fasthttp.StatusNotFound)
		return
	}
	ctx.SetContentType("text/json; charset=utf-8")
	switch string(ctx.Path()) {

	case "/createupdate":
		middlewareEndpoint(ctx, endpointsHandler.HandlerCreateUpdate)

	case "/retrieve":
		middlewareEndpoint(ctx, endpointsHandler.HandlerRetrieve)

	default:
		ctx.Error(schema.ERRPATH, fasthttp.StatusNotFound)
	}
}