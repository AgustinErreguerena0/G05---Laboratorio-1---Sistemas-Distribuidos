/*

3. Implemente una estructura Alumno con nombre, una lista de notas y un método Promedio()
que devuelve el promedio de notas. Escriba un programa que permita obtener el promedio
de varios alumnos. Sugerencia: no es necesario que cargue los datos de los alumnos, puede
definirlos al crearlos.

*/

package main

import "fmt"

type Alumno struct {
	nombre string
	notas  []float64
}

func (a Alumno) Promedio() float64 {
	var suma float64
	for _, nota := range a.notas {
		suma += nota
	}
	if len(a.notas) == 0 {
		return 0
	} else {
		return suma / float64(len(a.notas))
	}
}

func main() {
	alumnos := []Alumno{
		{nombre: "Juan", notas: []float64{7.5, 8.0, 9.0}},
		{nombre: "María", notas: []float64{6.0, 7.5, 8.0, 9.0}},
		{nombre: "Carlos", notas: []float64{10.0, 9.5, 8.5}},
	}

	for _, alumno := range alumnos {
		promedio := alumno.Promedio()
		fmt.Printf("El promedio de %s es: %.2f\n", alumno.nombre, promedio)
	}
}
