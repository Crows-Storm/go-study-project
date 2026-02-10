package main

import "fmt"

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method notifies users of different events.
func (u *user) notify() { // method up
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

//func (u *admin) notify() {	// method hiding user.notify
//	fmt.Printf("Sending admin email To %s<%s>\n", u.name, u.email)
//}

type admin struct {
	user  // Embedded Type
	level string
}

func main() {
	// We are now constructing outer type admin and inner type user.
	// This inner type value now looks like a field, but it is not a field.
	// We can access it through the type name like a field.
	// We are initializing the inner value through the struct literal of user.
	ad := admin{
		user: user{
			name:  "Crow S",
			email: "crow@gmail.com",
		},
		level: "superuser",
	}

	// We can access the inner type's method directly.
	ad.user.notify()

	// Because of inner type promotion, we can access the notify method directly through the outer
	// type. Therefore, the output will be the same.
	ad.notify()
}
