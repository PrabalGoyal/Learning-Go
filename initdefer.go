package main

import "fmt"

func main() {
	defer func() { fmt.Println("Exiting main function") }()
	fmt.Println("Inside Main")
}
func init() {
	fmt.Println("Inside Init")
}
