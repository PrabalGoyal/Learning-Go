// Golang program to show how to declare and define the struct
package main

import "fmt"

// Defining a struct type
type Address struct {
	Name string
	city string
	Pincode int
}

func main() {

	var a Address // Declaring a variable of a `struct` type All the struct fields are initialized with their zero value
	fmt.Println(a)


	a1 := Address{"Rahul", "Delhi", 3623572} // Declaring and initializing a struct using a struct literal

	fmt.Println("Address1: ", a1)

	a2 := Address{Name: "Anita", city: "Bangalore",  	// Naming fields while initializing a struct
								Pincode: 277001}

	fmt.Println("Address2: ", a2)


	a3 := Address{Name: "Delta"} // Uninitialized fields are set to their corresponding zero-value
	fmt.Println("Address3: ", a3)
}
