package main

import (
	"fmt"
)

//---------ASINCRONICO---------//
//declarar una estructura
type Persona struct{
	Nombre string
	Edad int
}

type Estudiante struct{
	Nombre string
	Apellido string
	DNI string
	Fecha string
}
func main (){
	fmt.Println("----EJERCICIO PERSONA-----")
p := Persona {
	Nombre: "Gabriel",
	Edad: 15,
}
p.Descripcion()

fmt.Println("\n\n-----EJERCICIO ESTUDIANTE-----")
e:= Estudiante{
	Nombre: "Maria",
	Apellido: "Jaramillo",
	DNI: "A58d88C",
	Fecha: "29/09/2024",
}

e.Detalle()
}

//Esta funcion imprime los datos de la instancia de persona
func (p *Persona )Descripcion (){
	fmt.Printf("%s tiene %d años de edad.",p.Nombre, p.Edad)
}
func  (e *Estudiante)Detalle(){
	fmt.Printf("Detalle alumno %s %s \nDNI: %s \nFecha de Ingreso: %s ", e.Nombre,e.Apellido,e.DNI,e.Fecha)
}