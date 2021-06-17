package main

import "fmt"

func main() {
	// unbuffered channel (Canal con búfer)
	ca := make(chan int, 1)

	ca <- 24

	fmt.Println(<-ca)
}
