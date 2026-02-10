package main

import (
	"fmt"
	"study-project/study_phase/exporting/counters"
	//"study-project/study_phase/exporting/counters"
	// This is a relative path to a physical location on our disk - relative to GOPATH.
	//"github.com/hoanhan101/ultimate-go/go/language/exporting/exporting_1/counters"
)

func main() {
	// Create a variable of the exported type and initialize the value to 10.

	counter := counters.AlertCounter(10)

	// alertCounter is an unexported named type that contains an integer counter for alerts.
	// The first character is in lower-case format so it is considered to be unexported.
	// It is not accessible for other packages, unless they are part of the package counters themselves.
	//counter := counters.alertCounter(10)	// compilation failed

	// However, when we create a variable of the unexported type and initialize the value to 10:
	// counter := counters.alertCounter(10)
	// The compiler will say:
	// - cannot refer to unexported name counters.alertCounter
	// - undefined: counters.alertCounter

	fmt.Printf("Counter: %d\n", counter)
	fmt.Printf("Counter: %d\n", counters.New(20)) // The counter is indirectly derived through the export function

	// Create a value of type User from the users package using struct literal.
	// However, since password is unexported, it cannot be compiled:
	// - unknown users.User field 'password' in struct literal
	u := counters.User{
		Name: "John Doe",
		ID:   10086,

		//password: "xxx",
	}

	fmt.Printf("User: %#v \n", u)
	counters.InitPwd(&u) // The counter is indirectly derived through the export function
	fmt.Printf("User: %#v \n", u)

	// Create a value of type Manager from the users package.
	// During construction, we are only able to initialize the exported field Title. We cannot
	// access the embedded type directly.
	school := counters.Manager{ // student is embedded in Manager
		Title: "Dev Manager",
	}
	// However, once we have the manager value, the exported fields from that unexported type are
	// accessible.
	fmt.Printf("school: %#v \n", school)
	school.Name = "Hoanh"
	school.ID = 101
	fmt.Printf("school: %#v \n", school)

	// Go's rule: Identifiers with a capital first letter can be exported, but lowercase ones cannot. This applies to:
	// Types (type User struct)
	//	 Functions (func DoWork())
	//	 Methods (func (u *User) Save())
	//	 Variables (const MaxSize = 100)
	// 	 Fields (struct fields)
}
