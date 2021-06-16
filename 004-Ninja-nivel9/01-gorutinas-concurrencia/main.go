package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("Número de CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Número de Gorutinas: %v\n", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hola desde primer Gorutina...")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hola desde la segunda Gorutina...")
		wg.Done()
	}()

	fmt.Printf("Número de CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Número de Gorutinas: %v\n", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Va a finalizar main...")

	fmt.Printf("Número de CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Número de Gorutinas: %v\n", runtime.NumGoroutine())
}
