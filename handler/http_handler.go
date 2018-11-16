package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/delioda79/pbsbjaak/message"
	"github.com/delioda79/pbsbjaak/server"
	"github.com/micro/protobuf/proto"
)

// PubSubHandler handles the http requests
type PubSubHandler struct {
	pbserver server.PubSubServer
}

func (psh PubSubHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		psh.handleGET(w, r)
		return
	}

	if r.Method == "POST" {
		psh.handlePOST(w, r)
		return
	}
}

func (psh PubSubHandler) handleGET(w http.ResponseWriter, r *http.Request) {
	rqstr := r.RemoteAddr
	if frwrd := r.Header.Get("X-Forwarded-For"); frwrd != "" {
		rqstr = frwrd
	}

	psh.pbserver.Subscribe(r.URL.Path, rqstr)
}

func (psh PubSubHandler) handlePOST(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error while gettig the body: ", err)
		w.WriteHeader(500)
		w.Write([]byte("Something went wrong"))
		return
	}

	pubMsg := &message.Publish{}
	err = proto.Unmarshal(b, pubMsg)
	if err != nil {
		log.Println("Error while gettig the message: ", err)
		w.WriteHeader(500)
		w.Write([]byte("Something went wrong while decoding the message"))
		return
	}
	sbs := psh.pbserver.Publish(pubMsg.GetTopic(), pubMsg.GetMessage())

	w.WriteHeader(200)
	w.Write([]byte(strconv.Itoa(sbs)))
}

// NewPubSubHandler returns a new PubSub HTTP handler
func NewPubSubHandler() PubSubHandler {
	srv := server.NewPubSubServer()
	return PubSubHandler{pbserver: srv}
}
