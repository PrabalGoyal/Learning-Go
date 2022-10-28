package main

import "fmt"

func main(){
	num1 := 8
	num2 := 10
  

	sum := num1 + num2
	fmt.Printf("%d + %d = %d\n", num1, num2, sum) 	// Adds two variables
  

	difference := num1 - num2
	fmt.Printf("%d - %d = %d\n",num1, num2,  difference) 	// Subtract two variables
  

	product := num1 * num2
	fmt.Printf("%d * %d is %d\n",num1, num2,  product)			// Multiply two variables

	quotient := num1 / num2
  fmt.Printf(" %d / %d = %d\n", num1, num2, quotient)			// Divide two integer variables

  num1 := 11.0
  num2 := 4.0

  
  result := num1 / num2
  fmt.Printf(" %g / %g = %g\n", num1, num2, result)  // Divide two floating point variables

  num := 5		 // increment of num by 1

  num++
  fmt.Println(num)  // 6

  // decrement of num by 1
  num--
  fmt.Println(num)  // 4
}