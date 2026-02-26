package main

import (
	"context"
	"fmt"
)

type user struct {
	name string
}

type userKey int

func main() {
	u := user{
		name: "Crow",
	}
	const uk userKey = 0

	ctx := context.WithValue(context.Background(), uk, &u)

	if u, ok := ctx.Value(uk).(*user); ok {
		fmt.Println("Found user:", u.name)
	}

	if _, ok := ctx.Value(0).(*user); !ok {
		fmt.Println("User Not Found")
	}

	const uk2 userKey = 0
	// use userKey type to get value, is true
	if u, ok := ctx.Value(uk2).(*user); ok {
		fmt.Println("Found user: ", u.name)
	}
}
