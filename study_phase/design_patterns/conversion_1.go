package main

import "fmt"

type Mover interface {
	Move()
}

type Locker interface {
	Lock()
	Unlock()
}

type MoveLocker interface {
	Mover
	Locker
}

type bike struct{}

func (bike) Move() {
	fmt.Println("Move the bike")
}

func (bike) Lock() {
	fmt.Println("Locking the bike")
}

func (bike) Unlock() {
	fmt.Println("Unlocking the bike")
}

func main() {
	var ml MoveLocker
	var m Mover

	// However, we cannot go in the other direction, like so:
	// ml = m
	//ml = m
	ml, y := m.(bike)                             // assertion
	fmt.Println("Does m has value of bike?: ", y) // false

	// The compiler will say:
	// cannot use m (type Mover) as type MoveLocker in assignment: Mover does not
	// implement MoveLocker (missing Lock method).

	// --------------
	// Type assertion
	// --------------

	// Interface type Mover does not declare methods named lock and unlock. Therefore, the compiler
	// can't perform an implicit conversion to assign a value of interface type Mover to an
	// interface value of type MoveLocker. It is irrelevant that the concrete type value of
	// type bike that is stored inside of the Mover interface value implements the MoveLocker interface.

	// We can perform a type assertion at runtime to support the assignment.

	// Perform a type assertion against the Mover interface value to access a COPY of the concrete type
	// value of type bike that was stored inside of it. Then assign the COPY of the concrete type
	// to the MoveLocker interface.

	// This is the syntax for type assertion.
	// We are taking the interface value itself, dot (bike). We are using bike as an parameter.
	// If m is not nil and there is a bike inside of m, we will get a copy of it since we are using value semantic.
	// Or else, a panic occurs.
	// b is having a copy of bike value.

	// Large-range interfaces can be converted to small-range interfaces
	// but small-range cannot convert to Large-range interfaces
	// because maybe small interfaces not included big interfaces with some interface
	ml = bike{}
	m = ml

	b := m.(bike)

	b, ok := m.(bike)

	fmt.Println("Does m has value of bike?: ", ok)

	ml = b
}
