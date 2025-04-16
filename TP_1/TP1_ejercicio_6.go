/*
Diseñar un sistema en anillo donde cinco nodos, representados por goroutines
, se envían mensajes de heartbeat cada 1 segundo entre sí de manera cíclica a través de canales. El sistema debe funcionar por 1 minuto y terminar.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Creo canales
	canal1 := make(chan string)
	canal2 := make(chan string)
	canal3 := make(chan string)
	canal4 := make(chan string)
	canal5 := make(chan string)

	//Canal "hecho" se asegura de que los nodos detecten que termino su periodo de escritura
	hecho := make(chan bool)
	
	//uso un waitgroup para esperar que el canal hecho termine
	var wg sync.WaitGroup

	//adiciono un canal al waitgroup (que seria "hecho")
	wg.Add(1)

	//inicio goroutines de dentro del anillo
	go nodo(0, canal1, canal2, hecho)
	go nodo(1, canal2, canal3, hecho)
	go nodo(2, canal3, canal4, hecho)
	go nodo(3, canal4, canal5, hecho)
	go nodo(4, canal5, canal1, hecho)
	
	//Escribo dentro del primer canal para iniciar el ciclo en forma de anillo
	canal1 <- "Inicio del heartbeat"

	//Uso time.After(1 * time.Minute) para terminar el sistema al pasar un minuto.
	go func() {
		time.Sleep(1 * time.Minute)
		//Cierro canal "hecho"
		close(hecho)
		//Termina la espera del waitgroup
		wg.Done()
	}()
	
	//wg espera que termine el minuto
	wg.Wait()

	


}
func nodo(id int, entrada <-chan string, salida chan<- string, done <-chan bool ) {
	//Para cada nodo
	for{
		// uso select para permitir a las goroutines esperar en multiples canales al mismo tiempo
		// se ejecuta cuando una de las operaciones del canal esta lista
		select {
			//El mensaje es de entrada
			case msg := <-entrada:  // entrada <- heartbeat ESCRIBE
			//Nodo recibe el heartbeat
			fmt.Println("Nodo", id, "recibió:", msg) //fmt.Println("heartbeat recibido desde nodo X")
			//Sleep por 1 segundo
			time.Sleep(1 * time.Second) //   time.Sleep(Time.second)
			//Nodo escribe su heartbeat a otro canal
			salida <- fmt.Sprintf("heartbeat desde nodo %d", id)  //heartbeat <- salida (esto RECIBE, NO ESCRIBE)
			//Nodo se da por terminado y se pasa al siguiente nodo
			case <-done:		 // done = true (NO ES UNA VARIABLE, ES UN CANAL, TENES QUE VALIDAR QUE EL CANAL TENGA EL VALOR QUE NECESITAS)
				return
		}
	}
}
