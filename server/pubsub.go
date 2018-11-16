package server

// PubSubServer is the PubSub server
type PubSubServer interface {
	Subscribe(topic, subscriber string) bool
	Publish(topic string, message []byte) int
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
func (pbs defaultPubSubServer) Publish(topic string, message []byte) int {
	tp, found := pbs.subscriptions[topic]
	if !found {
		return 0
	}

	return len(tp)
}

// NewPubSubServer returns a new PubSubServer
func NewPubSubServer() PubSubServer {
	return &defaultPubSubServer{
		subscriptions: map[string][]string{},
	}
}
