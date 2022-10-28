package main

import "fmt"
import "math/rand"
import "time"

func main(){
	fmt.Println(rand.Intn(200))   // Intn returns, as an int, a non-negative pseudo-random number in
    fmt.Println(rand.Intn(200))  // [0,n) from the default Source. i.e. simply call Intn to get the next random integer
    fmt.Println(rand.Intn(200))
    fmt.Println()

	fmt.Println(rand.Float64())
  
    // By default, it uses the value 1.
    fmt.Println((rand.Float64() * 8) + 7)
    fmt.Println((rand.Float64() * 8) + 7)
    fmt.Println()

	x1 := rand.NewSource(time.Now().UnixNano())
    y1 := rand.New(x1)
      
    fmt.Println(y1.Intn(200))
    fmt.Println(y1.Intn(200))
    fmt.Println()
  
    x2 := rand.NewSource(55)
    y2 := rand.New(x2)
    fmt.Println(y2.Intn(200))
    fmt.Println(y2.Intn(200))
    fmt.Println()
      
    x3 := rand.NewSource(5)
    y3 := rand.New(x3)
    fmt.Println(y3.Intn(200))
    fmt.Println(y3.Intn(200))  //The above method is not safe if the user wants to keep the random numbers secret.
}