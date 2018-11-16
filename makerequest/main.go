package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	messagedef "github.com/delioda79/pbsbjaak/message"
	proto "github.com/golang/protobuf/proto"
)

func main() {
	var topic = flag.String("topic", "", "topic")
	var message = flag.String("message", "", "message")
	var url = flag.String("url", "http://localhost:9090", "URL")
	flag.Parse()

	bts, _ := proto.Marshal(&messagedef.Publish{Topic: *topic, Message: []byte(*message)})

	rsp, err := http.Post(*url, "", bytes.NewBuffer(bts))
	if err != nil {
		log.Fatal(err)
	}

	bd, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bd))
}
