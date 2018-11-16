package server

import (
	"bytes"
	"net/http"
)

// PubSubServer is the PubSub server
type PubSubServer interface {
	Subscribe(topic, subscriber string) bool
	Publish(topic string, message []byte) (int, int)
}

type defaultPubSubServer struct {
	subscriptions map[string][]string
}

/**
Subscribe subscribes a subscriber to a topic and return true if it is a new subscriber for
 the topic
 **/
func (pbs *defaultPubSubServer) Subscribe(topic, subscriber string) bool {
	tp, found := pbs.subscriptions[topic]
	if !found {
		pbs.subscriptions[topic] = []string{subscriber}
		return true
	}

	for _, sbs := range tp {
		if sbs == subscriber {
			return false
		}
	}

	pbs.subscriptions[topic] = append(pbs.subscriptions[topic], subscriber)
	return true
}

/**
Publish sends `message` to all the subscriber for `topic`
**/
func (pbs defaultPubSubServer) Publish(topic string, message []byte) (subscribers int, sent int) {
	tp, found := pbs.subscriptions[topic]
	if !found {
		return 0, 0
	}

	sent = 0
	for _, sbs := range tp {
		_, err := http.Post(sbs, "", bytes.NewBuffer(message))
		if err == nil {
			sent++
		}
	}
	return len(tp), sent
}

// NewPubSubServer returns a new PubSubServer
func NewPubSubServer() PubSubServer {
	return &defaultPubSubServer{
		subscriptions: map[string][]string{},
	}
}
