package main

import "fmt"

type student struct {
	name    string
	surname string
	age     int
}

type sstudent struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {

	var s1 = student{} // create struct and init all default value
	fmt.Println(s1)

	s2 := student{"Jack", "jj", 18} // no comma is needed at the end

	//sx := student{"Jack",
	//	"jj",
	//	18	// need add ','
	//}
	fmt.Println(s2)

	s3 := struct {
		flag    bool
		counter int16
		pi      float32
	}{true, 12, 3.141592}
	fmt.Println(s3)

	fmt.Println("Flag", s3.flag)
	fmt.Println("Counter", s3.counter)
	fmt.Println("Pi", s3.pi)

	var s4 sstudent
	s4 = s3

	fmt.Printf("%+v\n", s4)
}
