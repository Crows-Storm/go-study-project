package main

import "fmt"

type Animal struct {
	Name     string
	IsMammal bool
}

func (a *Animal) Speak() {
	fmt.Println("UGH!", "My name is: ", a.Name, ", its: ", a.IsMammal, "I am a mammal")
}

type Dog struct {
	Animal
	PackFactor int
}

func (d *Dog) Speak() {
	fmt.Println("I am a dog", "My name is: ", d.Name, ", its: ", d.IsMammal, ", I am a mammal",
		"I am a mammal with a pack factor of", d.PackFactor)
}

type Cat struct {
	Animal
	ClimbFactor int
}

func (c *Cat) Speak() {
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

	an := Animal{
		Name:     "Humans",
		IsMammal: mammal,
	}
	an.Speak()

	// It's all fine until this one. This code will not compile.
	// Here, we try to group the Cat and Dog based on the fact that they are Animals. We are trying
	// to leverage sub-typing in Go. However, Go doesn't have it.
	// Go doesn't encourage us to group types by common DNA.
	// We need to stop designing APIs around this idea that types have a common DNA because if we
	// only focus on who we are, it is very limiting on who can we group with.
	// Sub-typing doesn't promote diversity. We lock types in a very small subset that can be
	// grouped with. But when we focus on behavior, we open up entire world to us.
	animals := []Animal{
		// Create a Dog by initializing its Animal parts and then its specific Dog attributes.
		Dog{
			Animal: Animal{
				Name:     "Fido",
				IsMammal: true,
			},
			PackFactor: 5,
		},

		// Create a Cat by initializing its Animal parts and then its specific Cat attributes.
		Cat{
			Animal: Animal{
				Name:     "Milo",
				IsMammal: true,
			},
			ClimbFactor: 4,
		},
	}
	// Have the Animals speak.
	for _, animal := range animals {
		animal.Speak()
	}
}
