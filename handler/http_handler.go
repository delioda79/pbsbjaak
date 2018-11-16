package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/delioda79/pbsbjaak/server"
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

	psh.pbserver.Subscribe("", rqstr)
}

func (psh PubSubHandler) handlePOST(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error while gettig the body: ", err)
		w.WriteHeader(500)
		w.Write([]byte("Something went wrong"))
		return
	}

	sbs, snt := psh.pbserver.Publish("", b)

	rsp, _ := json.Marshal(map[string]int{"subscribers": sbs, "sent": snt})
	w.WriteHeader(200)
	w.Write(rsp)
}

// NewPubSubHandler returns a new PubSub HTTP handler
func NewPubSubHandler() PubSubHandler {
	srv := server.NewPubSubServer()
	return PubSubHandler{pbserver: srv}
}
