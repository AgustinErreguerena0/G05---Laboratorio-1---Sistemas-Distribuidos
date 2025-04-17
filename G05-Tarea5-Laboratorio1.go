/* Escriba un programa que reciba como argumento un nombre de archivo y muestre por
consola su contenido. Si el archivo no existe se debe mostrar al usuario un mensaje y
terminar.*/

package main

import (
	"fmt"
	"os" // Para leer archivos y manejar errores del sistema
)

// Funcion principal para mostrar el contenido de un archivo seleccionado por el usuario
func main() {
	var nombreArchivo string

	// Se le pide al usuario que ingrese el nombre del archivo a leer
	fmt.Print("Ingrese el nombre del archivo a leer: ")
	fmt.Scan(&nombreArchivo)

	// Se intenta leer todo el contenido del archivo
	contenido, err := os.ReadFile(nombreArchivo)

	// Se verifica si hubo un error
	if err != nil {
		fmt.Printf("Error al abrir el archivo '%s':    %v\n", nombreArchivo, err)
		return
	}

	// Si no hubo errores, se imprime el contenido del archivo en la terminal
	fmt.Println("\nContenido del archivo:")
	fmt.Println(string(contenido))
}
