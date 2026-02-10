package main

import "fmt"

type Error interface {
	Error() string
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	return &errorString{text}
}

func main() {
	if _, err := webCall(true); err != nil { // err not nul
		fmt.Println(err)
		return
	}

	if rep, err := webCall(false); err != nil { // err is nul
		fmt.Println(rep, err)
	}
	fmt.Println("Life is good")
}

func webCall(success bool) (string, error) {
	if success {
		return "ok", nil
	}
	return "bad", New("Bad Request")
}
