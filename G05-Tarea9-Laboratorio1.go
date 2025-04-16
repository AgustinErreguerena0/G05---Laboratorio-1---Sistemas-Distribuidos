/*Simular un middleware donde un único publicador envía cada 1 segundo un mensaje (por ejemplo: "evento-X") a 3 suscriptores.
Cada suscriptor está representado por una goroutine que escucha su propio canal y muestra los eventos recibidos.
El sistema debe permitir que todos los suscriptores reciban el mismo mensaje simultáneamente y en igual orden.*/

package main

import (
	"fmt"
	"time"
)

// Simular un middleware donde un único publicador envía cada 1 segundo un mensaje (por ejemplo: "evento-X")
func publicador(mensaje string, chans []chan string) {
	x := 1
	for {
		evento := fmt.Sprintf("%v_%v", mensaje, x)
		for i := 0; i < len(chans); i++ {
			chans[i] <- evento
		}
		x++
		time.Sleep(1 * time.Second)
		fmt.Println("")
	}
}

func suscriptor(numero_sus int, ch <-chan string, chSin1, chSin2 chan bool) {
	for {
		//que escucha su propio canal y muestra los eventos recibidos.
		mensaje := <-ch
		//El sistema debe permitir que todos los suscriptores reciban el mismo mensaje simultáneamente
		// y en igual orden. -> ESTO FUE LOS MAS COMPLEJO
		switch numero_sus {
		case 1:
			fmt.Printf("Suscriptor %d, Mensaje: %s\n", numero_sus, mensaje)
			chSin1 <- true // El suscriptor 1 habilita al suscriptor 2
		case 2:
			<-chSin1
			fmt.Printf("Suscriptor %d, Mensaje: %s\n", numero_sus, mensaje)
			chSin2 <- true // El suscriptor 2 habilita al suscriptor 3
		case 3:
			<-chSin2
			fmt.Printf("Suscriptor %d, Mensaje: %s\n", numero_sus, mensaje)
		}
	}
}

func main() {
	// Canales de sincronización entre suscriptores
	chSin1 := make(chan bool)
	chSin2 := make(chan bool)

	// Canales de mensajes
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	chans := []chan string{ch1, ch2, ch3}

	// a 3 suscriptores. Cada suscriptor está representado por una goroutine
	go suscriptor(1, ch1, chSin1, chSin2)
	go suscriptor(2, ch2, chSin1, chSin2)
	go suscriptor(3, ch3, chSin1, chSin2)

	// Lanzar publicador
	go publicador("evento", chans)

	select {}
}
