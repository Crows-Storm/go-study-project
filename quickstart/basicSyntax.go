package main

import (
	"fmt"
)

func main() {
	// ================ Hello World ================
	fmt.Print("hello world !!!")

	baseType()
	inferredType()
	withoutInitialValue()
	multipleVariableDeclaration()
	differentMultipleVariableDeclaration()
	variableDeclarationInBlock()
	constantsVariable()
	printfFunction()
	arraysVariable()
	createSliceFromAnArray()
	createSliceWithTheMake()
	accessAndChangeAndAppendAndCopySlices()
	arithmeticOperators()
	loopArray()

	fmt.Println(myReturnFunction()) // 13
	res, result := myMultipleReturnFunction()
	fmt.Println(res)    // 13
	fmt.Println(result) // Successfully

	_, result = myMultipleReturnFunction() // use '_' omit some return
	fmt.Println(result)                    // Successfully

	fmt.Println(myMultipleReturnFunction()) // 13 Successfully

	fmt.Println("========== recursion Function ==========")
	function := recursionFunction(0)
	fmt.Println(function)

	fmt.Println("========== recursion factorial_recursion ==========")
	y := factorial_recursion(4)
	fmt.Println(y)

	fmt.Println("========== student ==========")
	var student Student
	student.name = "crow"
	student.age = 18
	fmt.Println(student)
	student.score = 101
	fmt.Println(student)

	fmt.Println("========== register Student ==========")
	success := registerStudent(student)
	fmt.Println(success)

	makeMaps()
}

func makeMaps() {
	fmt.Println("========== make Maps ==========")

	// map: k v pairs, map[{keyType}]{valueType}, hash table sorting mapping
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"} // create and init a map, k:v
	a2 := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
	var b = make(map[string]string) // make an empty map
	fmt.Println(a)
	fmt.Println(a2)
	fmt.Println(b)

	fmt.Println("-------------")
	// put data to map
	a["agent"] = "janus"
	b["agent"] = "janus"
	fmt.Println(a)
	fmt.Println(b)

	// Update Map Elements
	fmt.Println("-------------")
	a["agent"] = "janus2" // if a["agent"] == nil, it will store this elements
	fmt.Println(a["agent"])

	// Access Map Elements
	fmt.Println("-------------")
	c := a["agent"]
	fmt.Println(c)
	fmt.Println(a["agent"])
	fmt.Println(a2["Oslo"])
	fmt.Println(b["agent"])

	// Remove Element from Map
	fmt.Println("-------------")
	delete(a, "brand")      // delete({map}, {key})
	fmt.Println(a["brand"]) // print empty

	// Check For Specific Elements in a Map
	fmt.Println("-------------")
	var a3 = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}
	// result NO1: v NO2: bool(is exist)
	val1, ok1 := a3["brand"] // Checking for existing key and its value
	val2, ok2 := a3["color"] // Checking for non-existing key and its value
	val3, ok3 := a3["day"]   // Checking for existing key and its value
	_, ok4 := a3["model"]    // Only checking for existing key and not its value

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

	// If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.
	fmt.Println("-------------")
	var a4 = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	d := a4 // address reference

	fmt.Println(a4)
	fmt.Println(d)

	// update a4[{key}]
	a4["brand"] = "Ford2"
	fmt.Println(a4)
	fmt.Println(d)

	// Iterate Over Maps
	for s, s2 := range a4 {
		fmt.Println(s, s2) // key, value
	}

	// Iterate Over Maps in a Specific Order
	fmt.Println("-------------")
	// Maps are unordered data structures. If you need to iterate over a map in a specific order, you must have a separate data structure that specifies that order.
	a5 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	var b2 []string // defining the order
	b2 = append(b2, "one", "two", "three", "four")

	b3 := make([]string, len(b2))
	b3 = append(b3, "three", "four", "one", "two") // messy

	for s, i := range a5 {
		fmt.Printf("%v : %v", s, i)
		fmt.Println()
	}

	fmt.Println("-------------")
	for _, s := range b2 {
		fmt.Printf("%v : %v", s, a5[s])
		fmt.Println()
	}

	fmt.Println("-------------")
	for _, s := range b3 {
		fmt.Printf("%v : %v", s, a5[s])
		fmt.Println()
	}
}

func registerStudent(student Student) (success bool) {
	if student.age >= 18 {
		success = true
		return
	}
	success = false
	return
}

type Student struct {
	name  string
	age   int
	score float32
}

func recursionFunction(x int) int {
	if x == 11 { // stop condition
		return x
	}
	fmt.Println(x)
	return recursionFunction(x + 1) // recursion
}

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1) // final  4 * (3 * 2)
	} else {
		y = 1
	}
	return
}

func myMultipleReturnFunction() (res int, result string) {
	fmt.Println("========== my Multiple return Function ==========")
	result = "Successfully"
	res = 13
	return
}

func myReturnFunction() (result int) {
	fmt.Println("========== my return Function==========")
	result += 12
	res := 13
	return res // Explicitly specify
}

func loopArray() {
	fmt.Println("========== loop array ==========")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15} // max index 14

	for i := 0; i < len(numbers); i++ {
		fmt.Println(i) // index number, start at 0
	}

	fmt.Println("-------------")
	for i := range numbers {
		fmt.Println(numbers[i]) // index number, start at 1
	}

	fmt.Println("-------------")
	for i, number := range numbers { // loop index and value, is k : v struct
		i = i
		fmt.Println(number)
	}

}

func arithmeticOperators() {
	fmt.Println("========== arithmetic Operators ==========")
	number := 10
	fmt.Println(number % 2) // result: 0
	number = number + 1
	fmt.Println(number % 2) // result: 1

	number = number - 1
	fmt.Println(number) // 10

	number = number * 2
	fmt.Println(number) // 20

	number = number / 2
	fmt.Println(number) // 10
	number++
	fmt.Println(number) // 11
	number--
	fmt.Println(number) // 10

	number += 5
	fmt.Println(number) // 15

	number -= 5
	fmt.Println(number) // 10

	number *= 5
	fmt.Println(number) // 50

	number /= 5
	fmt.Println(number) // 10

	number %= 5
	fmt.Println(number) // 0

	fmt.Println("-------------")
	number = 10
	number &= 3
	fmt.Println(number) // 2

	number = 10

	//numByte := byte(number)
	//factorByte := byte(3)
	//fmt.Println(numByte)
	//fmt.Println(factorByte)
	number |= 3
	//fmt.Println(string(numByte), " ", string(factorByte))
	fmt.Println(number) // 11

	number = 10
	number ^= 3
	// The same is 0, and the difference is 1
	fmt.Println(number) // 9

	number = 10
	number >>= 2
	fmt.Println(number) // 2

	number = 10
	number <<= 2
	fmt.Println(number) // 40
}

func accessAndChangeAndAppendAndCopySlices() {
	fmt.Println("========== Go Access, Change, Append and Copy Slices ==========")
	prices := []int{10, 20, 30}

	// access
	fmt.Println(prices[0])
	fmt.Println(prices[1])

	fmt.Println("-------------")
	// change
	prices[0] = 12
	fmt.Println(prices[0])
	fmt.Println(prices[1])

	fmt.Println("-------------")
	// append
	prices = append(prices, 30)
	fmt.Println(prices[0])
	fmt.Println(prices[1])
	fmt.Println(prices[2])
	prices2 := []int{40, 50, 60}
	prices2 = append(prices, prices2...)
	fmt.Println(prices2)

	prices2 = prices2[1:5]
	prices3 := append(prices2, 88, 89, 90)
	fmt.Println(prices3)

	fmt.Println("-------------")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10] // numbers[:len(numbers)-10]: index start at 0, index end at len(numbers) - 10
	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}

func createSliceWithTheMake() {
	fmt.Println("========== Create a Slice With The make ==========")
	/*
	   myslice1 = [0 0 0 0 0]
	   length = 5
	   capacity = 10
	   -------------
	   myslice2 = [0 0 0 0 0]
	   length = 5
	   capacity = 5
	*/
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	fmt.Println("-------------")
	// with omitted capacity
	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))
}

func createSliceFromAnArray() {

	fmt.Println("========== create Slice From An Array ==========")
	arr1 := [6]int{1, 2, 3, 4, 5, 6}
	mySlice := arr1[2:4] // From index 2, to index 4 (excluding 4), so mySlice = [3, 4]
	fmt.Printf("myslice = %v\n", mySlice)
	fmt.Printf("myslice = %d\n", len(mySlice))
	fmt.Printf("capacity = %d\n", cap(mySlice))
}

func arraysVariable() {
	fmt.Println("========== arrays Variable ==========")
	stringArray := [32]string{}
	fmt.Println(stringArray)
	stringArra2 := [3]string{"1", "2", "3"}
	fmt.Println(stringArra2)

	var arr1 = [...]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8}

	fmt.Println("-------------")
	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Println("-------------")
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars)

	fmt.Println("-------------")
	fmt.Println(cars[0])
	fmt.Println(cars[1])
	fmt.Println(cars[2])
	func() {
		defer func() {
			if r := recover(); r != nil { // cache a error, if error it will call this func
				fmt.Println("defer cache:", r)
			}
		}()
		fmt.Println(cars[getIndexOutOfBounds(len(cars))]) // Index out of bounds, print: defer cache: runtime error: index out of range [4] with length 4
	}()

	fmt.Println(cars[getIndexOutOfBounds(len(cars))-1])

	fmt.Println("-------------")
	// if an element does not reach the set length, it will be padded with is type default value
	arr3 := [5]int{}
	arr4 := [5]int{1, 2, 3}
	arr5 := [5]int{4, 5, 6, 7, 8}
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)

	arr6 := [5]string{}
	arr7 := [5]string{"Volvo", "Mazda"}
	arr8 := [5]string{"Volvo", "Mazda", "BMW", "Ford", "Mazda"}
	fmt.Println(arr6)
	fmt.Println(arr7)
	fmt.Println(arr8)

	fmt.Println("-------------")
	// Specify an array of fixed length, and add elements to the specified index
	arr9 := [5]int{1: 10, 2: 40}
	fmt.Println(arr9)

	// len() function - returns the length of the slice (the number of elements in the slice)
	// cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

	fmt.Println("-------------")
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	myslice3 := [7]string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice3))
	fmt.Println(cap(myslice3))
	fmt.Println(myslice3)
}

func getIndexOutOfBounds(length int) int {
	// should is length - 1, array starts at index 0
	return length
}

func printfFunction() {
	fmt.Println("========== printf Function ==========")
	fmt.Print("hello world !!!")
	fmt.Print("hello ", "world", " !!!")
	// The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end:
	fmt.Println("Hello world !!!")

	// %v is used to print the value of the arguments
	// %T is used to print the type of the arguments
	var i = "Hello"
	var j = 15
	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)

	/*
		%v	Prints the value in the default format
		%#v	Prints the value in Go-syntax format
		%T	Prints the type of the value
		%%	Prints the % sign
	*/
	var je = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", je)
	fmt.Printf("%#v\n", je)
	fmt.Printf("%v%%\n", je)
	fmt.Printf("%T\n", je)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	/*
		Integer Formatting Verbs
			%b	Base 2
			%d	Base 10
			%+d	Base 10 and always show sign
			%o	Base 8
			%O	Base 8, with leading 0o
			%x	Base 16, lowercase
			%X	Base 16, uppercase
			%#x	Base 16, with leading 0x
			%4d	Pad with spaces (width 4, right justified)
			%-4d	Pad with spaces (width 4, left justified)
			%04d	Pad with zeroes (width 4
	*/
	var intV = 15
	fmt.Println("-------------")
	fmt.Printf("%b\n", intV)
	fmt.Printf("%d\n", intV)
	fmt.Printf("%+d\n", intV)
	fmt.Printf("%o\n", intV)
	fmt.Printf("%O\n", intV)
	fmt.Printf("%x\n", intV)
	fmt.Printf("%X\n", intV)
	fmt.Printf("%#x\n", intV)
	fmt.Printf("%4d\n", intV)
	fmt.Printf("%-4d\n", intV)
	fmt.Printf("%04d\n", intV)

	/*
		String Formatting Verbs
			%s	Prints the value as plain string
			%q	Prints the value as a double-quoted string
			%8s	Prints the value as plain string (width 8, right justified)
			%-8s	Prints the value as plain string (width 8, left justified)
			%x	Prints the value as hex dump of byte values
			% x	Prints the value as hex dump with spaces
	*/
	var txt2 = "Hello"

	fmt.Println("-------------")
	fmt.Printf("%s\n", txt2)
	fmt.Printf("%q\n", txt2)
	fmt.Printf("%8s\n", txt2)
	fmt.Printf("%-8s\n", txt2)
	fmt.Printf("%x\n", txt2)
	fmt.Printf("% x\n", txt2)

	/*
		Boolean Formatting Verbs
			%t	Value of the boolean operator in true or false format (same as using %v)
	*/
	var bi = true
	var bj = false
	var bin = 0
	var bjn = 1

	fmt.Println("-------------")
	fmt.Printf("%t\n", bi)
	fmt.Printf("%t\n", bj)
	fmt.Printf("%t\n", bin == 0)
	fmt.Printf("%t\n", bjn == 0)

	/*
		Float Formatting Verbs
			%e	Scientific notation with 'e' as exponent
			%f	Decimal point, no exponent
			%.2f	Default width, precision 2
			%6.2f	Width 6, precision 2
			%g	Exponent as needed, only necessary digits
	*/

	fmt.Println("-------------")
	var fi = 3.141

	fmt.Printf("%e\n", fi)
	fmt.Printf("%f\n", fi)
	fmt.Printf("%.2f\n", fi)
	fmt.Printf("%6.2f\n", fi)
	// if fi=3.1410 print 3.141
	fmt.Printf("%g\n", fi)
}

const G = 123
const (
	H string = "Hello"
	I int    = 100
)

func constantsVariable() {
	fmt.Println("========== constants Variable ==========")
	const A string = "Hello"
	const B = "World"
	fmt.Println(A, B)

	const C, D = "Hello", "World"
	fmt.Println(C, D)
	const (
		E int  = 10
		F bool = true
	)
	fmt.Println(E)
	fmt.Println(F)

	fmt.Println(G)

	fmt.Println(H)
	fmt.Println(I)

}

func variableDeclarationInBlock() {
	fmt.Println("========== variable Declaration In Block ==========")
	var (
		a int
		b = 12
		c = "Hello world !!!"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}

func differentMultipleVariableDeclaration() {
	fmt.Println("========== different multiple Variable Declaration ==========")
	var a, b = 12, "Hello"
	var c, d = 13, "World"
	fmt.Println(a, c)
	//log.Fatalf("Hello world !!! {}", "from log")
	fmt.Println(b, d)
}

func multipleVariableDeclaration() {
	fmt.Println("========== multiple Variable Declaration ==========")
	var a, b, c, d int

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
	fmt.Println("d =", d)

	a = 1
	b = 2
	c = 3
	d = 4

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
	fmt.Println("d =", d)

	var e, f int = 12, 13
	fmt.Println("e =", e)
	fmt.Println("f =", f)
}

func withoutInitialValue() {
	fmt.Println("========== without Initial Value ==========")
	var str string     // default ""
	var integer int    // default 0
	var double float32 // default 0
	var boolean bool   // default false

	fmt.Println(str)
	fmt.Println(integer)
	fmt.Println(double)
	fmt.Println(boolean)

	str = "hello world"
	integer = 100
	double = 3.14
	boolean = true

	fmt.Println(str)
	fmt.Println(integer)
	fmt.Println(double)
	fmt.Println(boolean)
}

func inferredType() {
	fmt.Println("========== inferred type ==========")
	str := "Hello world !!!"
	fmt.Println(str)
	int := 100
	fmt.Println(int)
	float := 0.23333333333333
	float2 := 0.23
	fmt.Println(float)
	fmt.Println(float2)
	bool := true
	fmt.Println(bool)
}

func baseType() {
	fmt.Println("========== base type ==========")
	var str = "Hello world !!!"
	fmt.Println(str)
	var int = 100
	fmt.Println(int)
	var float = 0.23333333333333
	var float2 = 0.23
	fmt.Println(float)
	fmt.Println(float2)
	var bool = true
	fmt.Println(bool)
}
