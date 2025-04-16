//1. Escriba una función SumarPares que reciba un slice de enteros y devuelva la suma de los números pares. Implemente un programa que demuestre su funcionamiento.

package main

import "fmt"

func main() {
	numeros := []int{22, -24, 0, 5, 39, 22, 1, 3, 7, 6}
	resultado := SumarPares(numeros)
	fmt.Println("La suma de los números pares es:", resultado)
}

func SumarPares(numeros []int) int {

	suma := 0
	for _, valor := range numeros {
		if valor%2 == 0 {
			suma += valor
		}
	}
	return suma
}
