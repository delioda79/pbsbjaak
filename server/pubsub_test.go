package server

import "testing"
import "strconv"

func TestSubscribe(t *testing.T) {
	srv := NewPubSubServer()

	topic := "topic1"
	sbscr := "subs1"
	res := srv.Subscribe(topic, sbscr)
	if !res {
		t.Errorf("We were expecting to add the subscriber %s to the topic %s", topic, sbscr)
		return
	}

	res = srv.Subscribe(topic, sbscr)
	if res {
		t.Errorf("We were expecting to not have added the subscriber %s to the topic %s", topic, sbscr)
		return
	}

	sbscr = "subs2"
	res = srv.Subscribe(topic, sbscr)
	if !res {
		t.Errorf("We were expecting to add the subscriber %s to the topic %s", topic, sbscr)
		return
	}

	topic = "topic2"
	res = srv.Subscribe(topic, sbscr)
	if !res {
		t.Errorf("We were expecting to add the subscriber %s to the topic %s", topic, sbscr)
		return
	}

	sbscr = "subs1"
	res = srv.Subscribe(topic, sbscr)
	if !res {
		t.Errorf("We were expecting to add the subscriber %s to the topic %s", topic, sbscr)
		return
	}

	topic = ""
	res = srv.Subscribe(topic, sbscr)
	if !res {
		t.Errorf("We were expecting to add the subscriber %s to the topic %s", topic, sbscr)
		return
	}
	sbscr = "subs2"
	res = srv.Subscribe(topic, sbscr)
	if !res {
		t.Errorf("We were expecting to add the subscriber %s to the topic %s", topic, sbscr)
		return
	}
}

func TestPublish(t *testing.T) {
	srv := NewPubSubServer()

	res := srv.Publish("", []byte("msg"))
	if res != 0 {
		t.Errorf("We were expecting 0 subscribers but we received %d", res)
	}

	topic := "topic1"

	for i := 0; i < 4; i++ {
		sbscr := "subs" + strconv.Itoa(i)
		srv.Subscribe(topic, sbscr)
	}

	topic = "topic2"
	for i := 0; i < 5; i++ {
		sbscr := "subs" + strconv.Itoa(i)
		srv.Subscribe(topic, sbscr)
	}

	res = srv.Publish("topic1", []byte("msg"))
	if res != 4 {
		t.Errorf("We were expecting 4 subscribers but we received %d", res)
	}

	res = srv.Publish("topic2", []byte("msg"))
	if res != 5 {
		t.Errorf("We were expecting 5 subscribers but we received %d", res)
	}

	res = srv.Publish("topic3", []byte("msg"))
	if res != 0 {
		t.Errorf("We were expecting 0 subscribers but we received only %d", res)
	}

	for i := 0; i < 4; i++ {
		sbscr := "subs" + strconv.Itoa(i)
		srv.Subscribe("", sbscr)
	}

	res = srv.Publish("", []byte("msg"))
	if res != 4 {
		t.Errorf("We were expecting 4 subscribers but we received %d", res)
	}
}
