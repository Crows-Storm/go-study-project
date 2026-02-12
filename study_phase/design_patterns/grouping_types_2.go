package main

import "fmt"

type Speaker interface {
	Speak()
}

type Animal struct {
	Name     string
	IsMammal bool
}

type Dog struct {
	Animal
	PackFactor int
}

func (d Dog) Speak() {
	fmt.Println("Woof!", "I am a dog", "My name is: ", d.Name, ", its: ", d.IsMammal, ", I am a mammal",
		"I am a mammal with a pack factor of", d.PackFactor)
}

func (d Dog) DrugSearch() {
	fmt.Println("Woof!", "I am a dog", "My name is: ", d.Name, ", i'm Drug Search!!!")
}

func (d Dog) Speak() {
	fmt.Println("Woof!", "I am a dog", "My name is: ", d.Name, ", its: ", d.IsMammal, ", I am a mammal",
		"I am a mammal with a pack factor of", d.PackFactor)
}

type Cat struct {
	Animal
	ClimbFactor int
}

func (c Cat) Speak() {
	fmt.Println("Meow!",
		"My name is:", c.Name,
		", it is:", c.IsMammal,
		"I am a mammal with a climb factor of", c.ClimbFactor)
}

func main() {

	mammal := true

	dog := Dog{
		Animal: Animal{
			Name:     "Wang Cai",
			IsMammal: mammal,
		},
		PackFactor: 5,
	}
	dog.Speak()

	cat := Cat{

		Animal: Animal{
			Name:     "Garfield",
			IsMammal: mammal,
		},
		ClimbFactor: 6,
	}
	cat.Speak()

	animals := []Speaker{
		// Create a Dog by initializing its Animal parts and then its specific Dog attributes.
		Dog{
			Animal:     Animal{"Fido", true},
			PackFactor: 5,
		},

		// Create a Cat by initializing its Animal parts and then its specific Cat attributes.
		Cat{
			Animal:      Animal{"Milo", true},
			ClimbFactor: 4,
		},
	}
	// Have the Animals speak.
	for _, animal := range animals {
		animal.Speak()
	}
}
