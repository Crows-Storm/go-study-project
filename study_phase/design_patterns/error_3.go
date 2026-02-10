package main

import (
	"fmt"
	"reflect"
)

type UnmarshalTypeError struct {
	Value string // description of JSON value
	Type  reflect.Type
}

func (e UnmarshalTypeError) Error() string {
	return "json: cannot unmarshal " + e.Value + " into Go value of type " + e.Type.String()
}

type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "json: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "json: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "json: Unmarshal(nil " + e.Type.String() + ")"
}

type user struct {
	Name int
}

func main() {
	var u user
	//err := Unmarshal([]byte(`{"name":"bill"}`), u) // Run with a value and pointer.
	//if err != nil {
	//	// This is a special type assertion that only works on the switch.
	//	switch e := err.(type) {
	//	case *UnmarshalTypeError: // Type Assertion
	//		fmt.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n", e.Value, e.Type)
	//	case *InvalidUnmarshalError: // Type Assertion
	//		fmt.Printf("InvalidUnmarshalError: [%s]\n", e.Error())
	//		fmt.Printf("InvalidUnmarshalError: Type[%v]\n", e.Type)
	//	default:
	//		fmt.Println(err)
	//	}
	//	return
	//}

	err2 := Unmarshal([]byte(`{"name":"bill"}`), &u) // Run with a value and pointer.
	if err2 != nil {
		// This is a special type assertion that only works on the switch.
		switch e := err2.(type) {
		case *UnmarshalTypeError: // Type Assertion
			fmt.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n", e.Value, e.Type)
		case *InvalidUnmarshalError: // Type Assertion
			fmt.Printf("InvalidUnmarshalError: [%s]\n", e.Error())
			fmt.Printf("InvalidUnmarshalError: Type[%v]\n", e.Type)
		default:
			fmt.Println(err2)
		}
		return
	}

	fmt.Println("Name:", u.Name)
}

// Unmarshal simulates an unmarshal call that always fails.
// Notice the parameters here: The first one is a slice of byte and the second one is an empty
// interface. The empty interface basically says nothing, which means any value can be passed into
// this function.
// We are going to reflect on the concrete type that is stored inside this interface and we are
// going to validate that if it is a pointer or not nil. We then return different error types
// depending on these.
func Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}
	return &UnmarshalTypeError{string(data), reflect.TypeOf(v)}
}

// There is one flaw when using type as context here. In this case, we are now going back to the
// concrete. We walk away from the decoupling because our code is now bounded to these concrete
// types. If the developer who wrote the json package makes any changes to these concrete types,
// that's gonna create a cascading effect all the way through our code. We are no longer protected
// by the decoupling of the error interface.

// This sometime has to happen. Can we do something different not to lose the decoupling. This is
// where the idea of behavior as context comes in.
