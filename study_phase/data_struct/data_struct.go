package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	slice1 := make([]string, 5)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"

	fmt.Printf("\n=> Printing a slice\n")
	fmt.Println(slice1)

	slice2 := make([]string, 5, 8)
	slice2[0] = "Apple"
	slice2[1] = "Orange"
	slice2[2] = "Banana"
	slice2[3] = "Grape"
	slice2[4] = "Plum"

	fmt.Printf("\n=> Length vs Capacity\n")
	inspectSlice(slice2)

	fmt.Printf("\n=> Idea of appending\n")

	var data []string

	// Capture the capacity of the slice.
	lastCap := cap(data)

	// Append ~100k strings to the slice.
	for record := 1; record <= 102400; record++ {
		data = append(data, fmt.Sprintf("Rec: %d", record))

		if lastCap != cap(data) {
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100
			lastCap = cap(data)

			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n", &data[0], record, cap(data), capChg)
		}
	}

	slice3 := slice2[2:4]

	fmt.Printf("\n=> Slice of slice (before)\n")
	inspectSlice(slice2)
	inspectSlice(slice3)

	slice3[0] = "CHANGED"

	fmt.Printf("\n=> Slice of slice (after)\n")
	inspectSlice(slice2)
	inspectSlice(slice3)

	slice4 := make([]string, len(slice2))
	copy(slice4, slice2)

	fmt.Printf("\n=> Copy a slice\n")
	inspectSlice(slice4)

	x := make([]int, 7)

	// Random starting counters.
	for i := 0; i < 7; i++ {
		x[i] = i * 100
	}

	// Set a pointer to the second element of the slice.
	twohundred := &x[1]

	x = append(x, 800)

	x[1]++

	// By printing out the output, we can see that we are in trouble.
	fmt.Printf("\n=> Slice and reference\n")
	fmt.Println("twohundred:", *twohundred, "x[1]:", x[1])

	fmt.Printf("\n=> UTF-8\n")

	s := "World means world"

	var buf [utf8.UTFMax]byte

	for i, r := range s {
		// Capture the number of bytes for this rune/code point.
		rl := utf8.RuneLen(r)

		// Calculate the slice offset for the bytes associated with this rune.
		si := i + rl

		copy(buf[:], s[i:si])

		// Display the details.
		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
	}
}

func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], slice[i])
	}
}
