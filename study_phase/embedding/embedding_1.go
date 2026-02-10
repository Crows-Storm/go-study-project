package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

type admin struct {
	person user // NOT Embedding
	level  string
}

func main() {
	ad := admin{
		person: user{
			name:  "Crow",
			email: "crow@gmail.com",
		},
		level: "superuser",
	}

	// We call notify through the person field through the admin type value.
	ad.person.notify()
}
