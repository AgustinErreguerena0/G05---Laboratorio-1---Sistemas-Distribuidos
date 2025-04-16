# G05---Laboratorio-1---Sistemas-Distribuidos
## FECHA CIERRE miércoles, 16 de abril de 2025, 23:00
## NOMBRE DE LOS ARCHIVOS:
**G05-Tarea#-Laboratorio1**

"#" indicar número de la tarea
## TAREAS ASIGNADAS:
### PulikoskiME
10. Desarrolle un programa que tenga una variable global x iniciada en 0 (cero) y una función
incrementar() que incremente en 5 a la variable x. La gorutina principal debe lanzar 100
gorutinas que invoquen a la función incrementar() y luego imprimir el valor de x. Ejecute su
programa usando la bandera -race para detectar si hay una carrera de datos. Además, el valor
final de x debe ser 500, pero es posible que observe que a veces es 490 o 495 u otros
valores. Usando WaitGroup y Mutexes, corrija su programa para que imprima el valor
correcto y no tenga una carrera de datos.
11. Escriba un programa que mediante el uso de un mutex global, escriba dos funciones donde:
la función a() debe bloquear el mutex, invocar a la función b() y desbloquear el mutex; la
función b() debe bloquear el mutex, imprimir “Hola mundo” y desbloquear el mutex. La
gorutina principal debe invocar a la función a(). Explica que sucede al ejecutar el programa.
