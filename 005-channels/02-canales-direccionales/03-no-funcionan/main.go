package main

import "fmt"

func main() {
	// unbuffered channel (Canal con búfer)
	ca := make(<-chan int, 2)

	ca <- 24
	ca <- 42

	fmt.Println(<-ca)
	fmt.Println(<-ca)
	fmt.Println("-------")
	fmt.Printf("%T\n", ca)
}
