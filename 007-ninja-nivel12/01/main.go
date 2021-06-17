package main

import (
	"fmt"

	"github.com/go-programming/007-ninja-nivel12/01/dog"
)

type canino struct {
	name string
	age  int
}

func main() {
	p1 := canino{
		name: "Annie",
		age:  dog.Age(5),
	}
	fmt.Println(p1)
}
