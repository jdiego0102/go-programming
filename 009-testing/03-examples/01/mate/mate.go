// Package mate nos permite comprobar que sabes sumar
package mate

// Sum suma una cantidad ilimtada números enteros
func Sum(xi ...int) int {
	var s int
	for _, v := range xi {
		s += v
	}
	return s
}
