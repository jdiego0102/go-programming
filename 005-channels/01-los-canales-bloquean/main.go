package main

import "fmt"

func main() {
	// unbuffered channel (Canal sin búfer)
	ca := make(chan int)

	ca <- 24
	fmt.Println(<-ca)
}
