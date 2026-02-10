package main

import "fmt"

type notifier interface {
	notify()
}

type printer interface {
	print()
}
type user struct {
	name  string
	email string
}

func (u user) print() {
	fmt.Printf("My name is %s and my email is %s\n", u.name, u.email)
}

// notify implements the notifier interface with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u *user) String() string {
	return fmt.Sprintf("My name is %q and my email is %q", u.name, u.email)
}

func main() {
	u := user{"Crow", "crow@hotmail.com"}

	sendNotification(&u)
	fmt.Println(u)
	fmt.Println(&u)

	entities := []printer{
		// When we store a value, the interface value has its own copy of the value.
		// Changes to the original value will not be seen.
		u, // just u

		// When we store a pointer, the interface value has its own copy of the address.
		// Changes to the original value will be seen.
		&u, // eques (*u)
	}

	// Change the name and email on the user value.
	u.name = "Crow"
	u.email = "corw@gmail.com"

	// Iterate over the slice of entities and call print against the copied interface value.
	for _, e := range entities {
		e.print()
	}
}

func sendNotification(n notifier) {
	n.notify()
}
