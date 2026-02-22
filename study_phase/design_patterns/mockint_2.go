package main

import (
	"fmt"
)

type PubSub struct {
	host string
}

func NewPubSub(host string) *PubSub {
	ps := PubSub{host: host}
	return &ps
}

func (ps *PubSub) Publish(key string, v interface{}) error {
	fmt.Println("Actual PubSub: Publish")
	return nil
}

func (ps *PubSub) Subscribe(key string) error {
	fmt.Println("Actual PubSub: Subscribe")
	return nil
}

// =======================================================================
type publisher interface {
	Publish(key string, v interface{}) error
	Subscribe(key string) error
}

type mock struct {
}

func (m *mock) Publish(key string, v interface{}) error {
	fmt.Println("Mock PubSub: Publish")
	return nil
}

func (m *mock) Subscribe(key string) error {
	fmt.Println("Mock PubSub: Subscribe")
	return nil
}

func main() {
	pubs := []publisher{
		NewPubSub("localhost"),
		&mock{},
	}

	for _, pub := range pubs {
		pub.Publish("key", "value")
		pub.Subscribe("key")
	}
}
