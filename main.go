package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/delioda79/pbsbjaak/handler"
)

func main() {
	var addr = flag.String("port", ":9090", "http service address")
	flag.Parse()

	hdlr := handler.NewPubSubHandler()

	http.Handle("/", hdlr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
