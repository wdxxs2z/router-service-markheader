package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/wdxxs2z/router-service-markheader/proxy"
	"github.com/wdxxs2z/router-service-markheader/roundTripper"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
)

var (
	port  = kingpin.Flag("port", "Port to listen").Envar("PORT").Short('p').Required().Int()
	debug = kingpin.Flag("debug", "Port to listen").Envar("DEBUG").Short('d').Bool()
        mark  = kingpin.Flag("mark", "Http header mark").Envar("MARK").Short('m').Required().String()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	httpClient := &http.Client{}
	roundTripper := roundTripper.NewLoggingRoundTripper(*debug)
	proxy := proxy.NewReverseProxy(roundTripper, httpClient, *debug, *mark)

	log.Fatal(http.ListenAndServe(":"+fmt.Sprintf("%v", *port), proxy))
}
