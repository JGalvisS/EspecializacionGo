package main

import "fmt"

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

type Prducto struct{
	nombre string
	precio float64
	cantidad int
}
type Servicio struct{
	nombre string
	precio	float64
	minTrabajados int
}
type Mantenimiento struct{
	nombre string
	precio float64
}
//Funcion que recibe un slice de producto y devuelve el precio total (precio por cantidad).
func SumarProductos(productos...Prducto) float64 {
	var precioTotal float64
	for _, producto := range productos {
		precioTotal += producto.precio * float64(producto.cantidad)
	}
	return precioTotal
}

//Funcion que recibe un slice de servicio y devuelve el precio total (precio por media hora trabajada.
//Si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
func SumaServicios(servicios...Servicio)float64 {
	var precioTotal float64
	for _, servicio := range servicios {
		if servicio.minTrabajados <= 30 {
			precioTotal += servicio.precio
		}else {//le agrego 29 a minTrabajados para asegurar que complete un nuevo bloque y divido entre para contar los bloques ya que usualmente go redondea hacia abajo 
			precioTotal += servicio.precio * float64((servicio.minTrabajados+29)/30)
		}
	}
	return precioTotal
}
//Funcion que recibe un slice de mantenimiento y devuelve el precio total.
func SumarMantenimientos(mantenimientos...Mantenimiento) float64 {
	var precioTotal float64
	for _, mantenimiento := range mantenimientos {
		precioTotal += mantenimiento.precio
	}
	return precioTotal
}
func main()  {
 productos := []Prducto{
	{"trapeador",12.00,6},
	{"jabon", 8.50,9},
	{"suavisante",10.00,10},
}

servicios := []Servicio{
	{"Limpieza", 50.00, 45},
	{"Pintura", 100.00, 90},
	{"Reparación", 75.00, 30},
}

mantenimiento:= []Mantenimiento{
	{"Nevera",180.00},
	{"Computador",120.00},
}


// Declarar los canales para recibir los resultados
productosCh := make(chan float64)
serviciosCh := make(chan float64)
mantenimientoCh := make(chan float64)

// Lanzar las goroutines
go func() {
	productosCh <- SumarProductos(productos...)
}()

go func() {
	serviciosCh <- SumaServicios(servicios...)
}()

go func() {
	mantenimientoCh <- SumarMantenimientos(mantenimiento...)
}()

// Recoger los resultados desde los canales
totalProductos := <-productosCh
totalServicios := <-serviciosCh
totalMantenimiento := <-mantenimientoCh

// Calcular el monto total
total := totalProductos + totalServicios + totalMantenimiento

// Imprimir los resultados
fmt.Println("Total productos: ", totalProductos)
fmt.Println("Total servicios: ", totalServicios)
fmt.Println("Total mantenimientos: ", totalMantenimiento)
fmt.Println("Monto total: ", total)

}

/*CONCEPTOS
Tus conclusiones son correctas. Aquí te confirmo cada una:

1. Los canales sirven para recoger las respuestas de las funciones de forma paralela:
Correcto. Los canales permiten que las goroutines (funciones que se ejecutan de manera concurrente)
envíen y reciban datos. En el caso que discutimos, los canales recogen las respuestas de las funciones 
SumarProductos, SumaServicios, y SumarMantenimientos de manera paralela, ya que estas funciones se ejecutan en goroutines.

2.Al usar el comando "go", esta función toma un hilo de ejecución y se ejecuta mientras se puede ejecutar 
el resto del contenido en main:
Correcto. Cuando usas go func() {...}(), esa función se ejecuta en una goroutine separada, que es como 
un "hilo" ligero de ejecución. Esto permite que el resto del código en main continúe ejecutándose sin 
esperar a que esa función termine.

3.Los resultados del canal deben ser almacenados en una variable para trabajar con ellos:
Correcto. Cuando una goroutine envía un resultado a un canal, ese resultado se debe recibir 
y almacenar en una variable para poder trabajar con él. Esto se hace con la sintaxis resultado := <-canal, 
que lee (o recibe) el valor enviado al canal y lo almacena en la variable resultado.
*/