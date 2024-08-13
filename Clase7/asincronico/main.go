package main

/*Calcular precio
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento. 
Para ello requieren realizar un programa que se encargue de calcular el precio total de 
Productos, Servicios y Mantenimientos. Debido a la fuerte demanda, y para optimizar la 
velocidad, requieren que el cálculo de la sumatoria se realice en paralelo mediante tres goroutines.
Se requieren tres estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.
Y se requieren tres funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio por cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio por media hora trabajada. 
Si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.
Las tres se deben ejecutar concurrentemente y, al final, se debe mostrar por pantalla el monto final (sumando el total de las tres).
*/

type Prductos struct{
	nombre string
	precio float64
	cantidad int
}
type Servicios struct{
	nombre string
	precio	float64
	minTrabajados int
}
type Mantenimiento struct{
	nombre string
	precio float64
}