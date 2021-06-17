package decir

import (
	"fmt"
	"testing"
)

func TestSaludar(t *testing.T) {
	s := Saludar("Juan Diego")
	if s != "Bienvenido estimado Juan Diego" {
		t.Error("Expected", "Bienvenido estimado Juan Diego", "Got", s)
	}
}

func ExampleSaludar() {
	fmt.Println(Saludar("Juan Diego"))
	// Output
	// Bienvenido estimado Juan Diego
}

func BenchmarkSaludar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Saludar("Juan Diego")
	}
}
