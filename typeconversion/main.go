package main

import "fmt"

func main() {
	var num1 int = 845
	var num2 float64 = float64(num1)
	var num3 int64 = int64(num1)
	var num4 uint = uint(num1)
	fmt.Println(num1, num2, num3, num4)
}
