package handler

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/delioda79/pbsbjaak/server/serverfakes"
)

func TestGet(t *testing.T) {

	srv := &serverfakes.FakePubSubServer{}

	hdlr := PubSubHandler{pbserver: srv}

	r := httptest.NewRequest("GET", "http://localhost", nil)
	r.RemoteAddr = "192.168.0.1"
	w := httptest.NewRecorder()
	hdlr.handleGET(w, r)

	_, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestPost(t *testing.T) {

	srv := &serverfakes.FakePubSubServer{}
	srv.PublishReturns(5)
	hdlr := PubSubHandler{pbserver: srv}

	r := httptest.NewRequest("GET", "http://localhost", nil)
	r.RemoteAddr = "192.168.0.1"
	w := httptest.NewRecorder()
	hdlr.handlePOST(w, r)

	res, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err.Error())
	}

	if string(res) != "5" {
		t.Errorf("We expected 5 but we got %v", string(res))
	}
}
