/*
Ejercicio 11: Escriba un programa que mediante el uso de un mutex global, escriba dos funciones donde: la función a() debe bloquear el mutex, invocar a la función b() y desbloquear el mutex; la función b() debe bloquear el mutex, imprimir “Hola mundo” y desbloquear el mutex. La gorutina principal debe invocar a la función a(). Explica que sucede al ejecutar el programa.
*/

package main
import (
	"fmt"
	"sync"
)

var mutex sync.Mutex // Creo el mutex global para proteger el acceso a la sección crítica

func a() {
	mutex.Lock() // Bloqueo el mutex antes de llamar a b()
	defer mutex.Unlock() // Desbloqueo el mutex al final de la función a
	b() // Llamo a la función b()
}

func b() {
	mutex.Lock() // Bloqueo el mutex antes de imprimir
	defer mutex.Unlock() // Desbloqueo el mutex al final de la función b
	fmt.Println("Hola mundo") // Imprimo "Hola mundo"
}

func main() {
	a() // Llamo a la función a desde la gorutina principal
}
// Al ejecutar el programa, se observa que la ejecución acaba con un "Deadlock" (bloqueo mutuo) porque la función a() bloquea el mutex y luego llama a b(), que también intenta bloquear el mismo mutex. Esto provoca un bloqueo, porque que ambas funciones esperan liberar el mutex, pero ninguna puede hacerlo porque están esperando la liberación del otro. 