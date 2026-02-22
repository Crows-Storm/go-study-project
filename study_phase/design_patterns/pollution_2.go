package main

import "fmt"

type Server struct {
	host string
}

func NewServer(host string) *Server {
	return &Server{host}
}

func (s *Server) Start() error {
	return nil
}
func (s *Server) Stop() error {
	return nil
}
func (s *Server) Wait() error {
	return nil
}

func main() {
	ser := NewServer("localhost")

	ser.Start()
	ser.Stop()
	ser.Wait()

	fmt.Println(ser.host)
	// Guidelines around interface pollution:
	// --------------------------------------
	// Use an interface:
	// - When users of the API need to provide an implementation detail.
	// - When APIs have multiple implementations that need to be maintained.
	// - When parts of the APIs that can change have been identified and require decoupling.
	// Question an interface:
	// - When its only purpose is for writing testable API’s (write usable APIs first).
	// - When it’s not providing support for the API to decouple from change.
	// - When it's not clear how the interface makes the code better.
}
