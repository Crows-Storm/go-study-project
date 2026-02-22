package main

import (
	"fmt"
	"math/rand"
	"time"
)

type car struct {
}

func (car) String() string {
	return "Vroom"
}

func (car) AddGas() string {
	return "Add Gas"
}

type cloud struct{}

func (cloud) String() string {
	return "Big Data"
}

func (cloud) Push() string {
	return "Push Data"
}

func main() {
	rand.NewSource(time.Now().UnixNano())

	mvs := []fmt.Stringer{
		car{},
		cloud{},
	}

	fmt.Println(mvs)

	for i := 0; i < 10; i++ {
		rn := rand.Intn(2)

		// About Assertion the Automatic Conversion
		if v, ok := mvs[rn].(cloud); ok { // assertion mvs[rn] is cloud
			v.Push() // cloud Unique method
			fmt.Println("Got Lucky: ", v)
			continue
		} else if v, ok := mvs[rn].(car); ok { // assertion mvs[rn] is car
			v.AddGas() // car Unique method
		}
		fmt.Println("Got Unlucky")
	}
}
