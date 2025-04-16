/*
Simular el acceso concurrente a un log compartido donde 10 goroutines,
 cada una representando un nodo, intentan registrar un evento crítico cada 0.5 segundos
  (por ejemplo: "nodo-3: temperatura alta" o "nodo-7: pérdida de conexión");
   cada evento debe escribirse en un slice de strings protegido por sync.
   Mutex para garantizar la integridad del log.
*/

/*
El log compartido es un slice que representa un registro histórico de eventos importantes (críticos) 
que distintos nodos van escribiendo.  
*/
package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

func main() {
	// Creo log para que los nodos registren sus eventos
	var log []string
	// Creo mutex para que los nodos ingresen sus eventos sin conflictos
	var mutex sync.Mutex
	// Creo un WaitGroup para esperar a que todas las goroutines terminen
	var wg sync.WaitGroup

	// Array de string de posibles eventos críticos
	eventos := []string{
		" CPU en sobrecarga (98%)",
		" fallo de conexión con el servidor primario",
		" temperatura crítica (92°C) - activado ventilación forzada",
		" disco SSD con sectores defectuosos ",
		" memoria RAM al 99% - killing procesos",
		" reinicio inesperado (kernel panic)",
		" pérdida de paquetes en interfaz eth0 (35%)",
		" intrusión detectada (IP 192.168.1.15)",
		" configuración corrupta en /etc/network/interfaces",
		" GPU en throttling térmico (105°C)",
	}

	// Crear 10 goroutines que representan nodos
	for i := 1; i <= 10; i++ {
		// añadimos goroutine a el waitgroup
		wg.Add(1)
		go nodo(i, eventos, &log, &mutex, &wg)
	}

	// El programa principal espera 3 segundos antes de finalizar
	time.Sleep(3 * time.Second)

	// Esperar a que todas las goroutines terminen
	wg.Wait()

	// Imprimir el log final después de que todas las goroutines terminaran
	fmt.Println("------ LOG FINAL ------")
	for _, entrada := range log {
		fmt.Println(entrada)
	}
}

// Función nodo, con id, con array de posibles eventos
// log tiene un puntero porque queremos acceder al log ubicado en main
// lo mismo pasa con mutex. y waitGroup
func nodo(id int, eventos []string, log *[]string, mutex *sync.Mutex, wg *sync.WaitGroup) {
	// decrementa el contador del WaitGroup en 1, indicando que una goroutine terminó su ejecución
	defer wg.Done()

	// Creador de números aleatorios propio para esta goroutine
	r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(id)))

	// Cada goroutine solo ejecuta durante 3 segundos
	fin := time.After(3 * time.Second)

	for {
		select {
		case <-fin:
			// Si se cumplió el tiempo límite de 3 segundos, terminamos la goroutine
			return
		default:
			// Delay de 0.5 segundos
			time.Sleep(500 * time.Millisecond)

			// Se elige el evento con el generador local
			evento := eventos[r.Intn(len(eventos))]

			// Proteger el acceso al log
			mutex.Lock()
			// Dentro de log (del main), concatenamos el evento crítico del nodo correspondiente
			*log = append(*log, fmt.Sprintf("nodo-%d: %s", id, evento))
			// Libero mutex
			mutex.Unlock()
		}
	}
}
