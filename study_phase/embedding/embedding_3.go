package main

import "fmt"

// notifier is an interface that defined notification type behavior.
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method notifies users of different events using a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

func (u *admin) notify() {
	fmt.Printf("Sending admin email To %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

func main() {
	// Create an admin user.
	ad := admin{
		user: user{
			name:  "Crow S",
			email: "crow@gmail.com",
		},
		level: "superuser",
	}

	// Send the admin user a notification.
	// We are passing the address of outer type value. Because of inner type promotion, the outer
	// type now implements all the same contract as the inner type.
	sendNotification(&ad.user) // call user notify, because specify call notify of admin.user
	sendNotification(&ad)      // call admin notify, because admin implemented notify method, so user.notify method is hidden

	// ==============================
	// Important: Behavior has now changed after *admin.notify()
	// ===============================

	// Embedding still does not create subtype relationships. user is still user, admin is still admin.

	// However, because admin defines its own notify() method:
	// 1. The notify() method of the inner type user is "hidden" - no longer "method hoisting"
	// 2. admin satisfies the notifier interface through its own implementation, instead of through the embedded user
	// 3. admin now has two ways to call notify():
	// a. Its own implementation: admin.notify()
	// b. user's implementation (hidden): admin.user.notify()

	// This means that the way admin implements the notifier interface has changed:
	// - Before: by "inheriting" user's method (method hoisting)
	// - Now: through its own defined method (method hiding + explicit implementation)

	// Key difference: admin.notify() no longer automatically calls or includes the logic of user.notify(),
	// unless a.user.notify() is explicitly called in admin.notify().
}

// We have our polymorphic function here.
// sendNotification accepts values that implement the notifier interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
