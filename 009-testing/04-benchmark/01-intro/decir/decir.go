// Package decir nos permite saludar
package decir

import "fmt"

// Saludar ejecuta un saludo a un usuario
func Saludar(s string) string {
	return fmt.Sprint("Bienvenido estimado ", s)
}
