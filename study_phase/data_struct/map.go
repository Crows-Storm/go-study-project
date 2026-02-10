package main

import "fmt"

type user struct {
	name    string
	surname string
}

func main() {
	// string : user struct
	users1 := make(map[string]user)

	// Add key/value pairs to the map.
	users1["Roy"] = user{"Rob", "Roy"}
	users1["Ford"] = user{"Henry", "Ford"}
	users1["Mouse"] = user{"Mickey", "Mouse"}
	users1["Jackson"] = user{"Michael", "Jackson"}

	// Iterate over map
	for k, v := range users1 {
		fmt.Printf("%s: %s\n", k, v)
	}

	// Declare and initialize the map with values.
	users2 := map[string]user{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}

	// Iterate over the map.
	fmt.Printf("\n=> Map literals\n")
	for key, value := range users2 {
		fmt.Println(key, value)
	}

	// delete key
	delete(users1, "Roy")

	u1, found1 := users1["Roy"]
	u2, found2 := users2["Roy"]

	// Display the value and found flag.
	fmt.Printf("\n=> Find key\n")
	fmt.Println("users1 Roy", found1, u1)
	fmt.Println("users2 Roy", found2, u2)

}
