/*
Ejercicio 10: Desarrolle un programa que tenga una variable global x iniciada en 0 (cero) y una función incrementar() que incremente en 5 a la variable x. La gorutina principal debe lanzar 100 gorutinas que invoquen a la función incrementar() y luego imprimir el valor de x. Ejecute su programa usando la bandera -race para detectar si hay una carrera de datos. Además, el valor final de x debe ser 500, pero es posible que observe que a veces es 490 o 495 u otros valores. Usando WaitGroup y Mutexes, corrija su programa para que imprima el valor correcto y no tenga una carrera de datos.
*/

package main

import (
	"fmt"
	"sync"
)

var x int = 0 // Variable global iniciada en 0
var waitGroup sync.WaitGroup // WaitGroup para esperar a que todas las gorutinas terminen
var mutex sync.Mutex // Mutex para proteger el acceso a la variable x

func incrementar() {
	defer waitGroup.Done() // Indico que la gorutina ha terminado
	
	mutex.Lock() // Bloqueo el mutex antes de modificar x
	defer mutex.Unlock() // Desbloqueo el mutex al final de la función
	x += 5
}

func main() {
	// Lanzo las 100 gorutinas
	for i := 0; i < 100; i++ {
		waitGroup.Add(1) // Incrementar el contador del WaitGroup
		go incrementar() // Lanzo la gorutina que llama a incrementar
	}
	waitGroup.Wait() // Esperar a que todas las gorutinas terminen

	// Imprimo el valor final de x
	fmt.Printf("Valor final de x: %d\n", x)
}