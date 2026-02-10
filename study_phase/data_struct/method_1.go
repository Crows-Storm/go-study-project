package main

import "fmt"

type user struct {
	name  string
	email string
}

func main() {
	crow := user{"Crow", "bill@email.com"}
	crow.notify()
	fmt.Println(crow.email)
	crow.changeEmail("crow@hotmail.com")
	fmt.Println(crow.email)

	users := []user{
		{"bill", "it@email.com"},
		{"Crow", "it@email.com"},
	}

	// We are ranging over this slice of values, making a copy of each value and call notify to
	// make another copy.
	for _, u := range users {
		u.notify()
	}

	// Iterate over the slice of users switching semantics. This is not a good practice.
	for _, u := range users {
		u.changeEmail(u.name + "@wontmatter.com")
	}
}

// this a method, not function
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// changeEmail implements a method with a pointer receiver: u of type pointer user
// Using the pointer receiver, the method operates on shared access.
func (u *user) changeEmail(email string) {
	u.email = email
	fmt.Printf("Changed User Email To %s\n", email)
}
