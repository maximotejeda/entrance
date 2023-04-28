package main

import (
	"entrance/external/v0.1/proxy"
	"entrance/external/v0.1/proxy/auth"
	"log"
	"net/http"

	"github.com/maximotejeda/helpers/middlewares"
)

func main() {
	// main will redirect each request to it correspondent place
	auth := auth.NewAuthProxy("http://localhost:8083/") //that information will be taken from k8s

	//create a new multiplexer
	mu := proxy.NewMultiplexer()
	// we can visit from the exterior without knowing version
	mu.Add(`(GET|POST|PUT|OPTIONS|HEAD) (/?[v0-9\.]{3,6})?/auth/.*`, middlewares.LoggerMiddleware(auth))

	//http.Handle("/auth/", mw.LoggerMiddleware(auth))
	log.Fatal(http.ListenAndServe(":8080", mu))
}
