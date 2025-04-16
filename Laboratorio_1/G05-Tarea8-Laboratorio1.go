/*
Simular un sistema de monitoreo donde 1 goroutina,
envían un ping a una lista de 3 nodos (nodo-1, nodo-2, nodo-3) cada 2 segundos;
cada ping simula una latencia aleatoria entre 100 y 500 milisegundos, la gorutina debe guardar el nodo con menor
latencia de respuesta de cada ronda en un slice. Luego de 10 rondas se debe imprimir los resultados y terminar.
Consulta: se debe proteger el slice con mutex. -> NO HIZO FALTA EL MUTEX SEGUN MI RAZONAMIENTO, CON LOS CONOCIMIENTOS BASICOS SE CUMPLIO CON LA CONSIGNA (SUPONGO).
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Canal struct {
	latencia time.Duration
	nodo     string
}

// cada ping simula una latencia aleatoria entre 100 y 500 milisegundos,
func ping() time.Duration {
	latencia := time.Duration(rand.Intn(400)+100) * time.Millisecond
	time.Sleep(latencia)
	return latencia
}

// Simular un sistema de monitoreo donde 1 goroutina -> EN ESTE CASO EL MAIN
func main() {
	ch := make(chan Canal, 3)
	var resultados []string
	nodos := [3]string{"nodo-1", "nodo-2", "nodo-3"}

	for ronda := 0; ronda < 10; ronda++ {

		//envían un ping a una lista de 3 nodos (nodo-1, nodo-2, nodo-3) cada 2 segundos;
		time.Sleep(2 * time.Second)
		fmt.Printf("Ronda %v\n", ronda+1)
		for i := 0; i < 3; i++ {
			go func(i int) {
				latencia := ping()
				ch <- Canal{latencia: latencia, nodo: nodos[i]}
			}(i)
		}

		p_lug := <-ch
		s_lug := <-ch
		t_lug := <-ch
		fmt.Printf("Primer lugar %v, su latencia %v\n", p_lug.nodo, p_lug.latencia)
		fmt.Printf("Segundo lugar %v, su latencia %v\n", s_lug.nodo, s_lug.latencia)
		fmt.Printf("Tercer lugar %v, su latencia %v\n", t_lug.nodo, t_lug.latencia)
		// la gorutina debe guardar el nodo con menor latencia de respuesta de cada ronda en un slice.
		resultados = append(resultados, p_lug.nodo)

	}
	//Luego de 10 rondas se debe imprimir los resultados y terminar.
	fmt.Println("\nResumen de ganadores por ronda:")
	for i := 0; i < len(resultados); i++ {
		fmt.Printf("Ronda %d: %s\n", i+1, resultados[i])
	}
}
