//2. Desarrolle un programa que lea una línea de texto desde la entrada estándar y cuente e imprima cuántas palabras tiene. Busque ayuda en los paquetes strings y fmt.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingrese una línea de texto: ")
	if scanner.Scan() {
		texto := scanner.Text()
		palabras := strings.Fields(texto)
		cantidad := len(palabras)
		fmt.Printf("El texto ingresado tiene %d palabras.\n", cantidad)
	}
}
