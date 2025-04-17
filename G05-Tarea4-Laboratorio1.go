/*Escriba una función que convierta de °C a °F y otra de °F a °C. Luego realice un menú para
elegir qué conversión hacer y pida los datos por teclado.*/

package main

import (
	"fmt"
)

// Función que convierte grados Celsius a Fahrenheit
func celsiusAFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

// Función que convierte grados Fahrenheit a Celsius
func fahrenheitACelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

/* Función principal para mostrar un menú con el que el usuario pueda
interactuar y donde se manejen las conversiones de temperatura */

func main() {
	var opcion int
	var valor float64

	// Se le muestran las opciones al usuario
	fmt.Println("Conversor de Temperatura:")
	fmt.Println("1 - Convertir °C a °F")
	fmt.Println("2 - Convertir °F a °C")
	fmt.Print("Seleccione una opción por favor (1 o 2): ")
	fmt.Scan(&opcion)

	// Se procesa la opción elegida por el usuario
	switch opcion {
	case 1:
		fmt.Print("Ingrese la temperatura en °C: ") // Se le pide al usuario que ingrese el valor a convertir
		fmt.Scan(&valor)
		resultado := celsiusAFahrenheit(valor)
		fmt.Printf("%.2f °C equivalen a %.2f °F\n", valor, resultado)

	case 2:
		fmt.Print("Ingrese la temperatura en °F: ") // Se le pide al usuario que ingrese el valor a convertir
		fmt.Scan(&valor)
		resultado := fahrenheitACelsius(valor)
		fmt.Printf("%.2f °F equivalen a %.2f °C\n", valor, resultado)

	default:
		fmt.Println("Opción no válida. Intente nuevamente.") //Error en caso de una opción no válida
	}
}
