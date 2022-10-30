package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2) // buffered channel
	go PushMessages(ch)
	PrintMessage(ch)

	// SAMPLE OUTPUT
	// Successfully wrote 0 to channel
	// Successfully wrote 1 to channel
	// Read value 0 from channel
	// Successfully wrote 2 to channel
	// Read value 1 from channel
	// Successfully wrote 3 to channel
	// Read value 2 from channel
	// Successfully wrote 4 to channel
	// Read value 3 from channel
	// Successfully wrote 5 to channel
	// Read value 4 from channel
	// Read value 5 from channel
}

func PrintMessage(ch <-chan int /*Receive only channel*/) {
	for msg := range ch {
		fmt.Println("Read value", msg, "from channel")
		time.Sleep(2 * time.Second)
	}
}

func PushMessages(ch chan<- int /*Send only channel*/) {
	for i := 0; i < 6; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to channel")
	}
	// close channel and clean memory
	close(ch)
}
