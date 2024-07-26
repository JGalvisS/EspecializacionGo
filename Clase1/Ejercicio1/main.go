package main

import ( "fmt"	)

func main() {
	//------------------------ASINCRONICO-----------------------//
	meses := []string{"Enero", "Febrero", "Junio", "Agosto", "Abril"}
	fmt.Println("-------ASINCRONICO--------")
	ObtenerEstacion(meses)
	ObtenerEstacionSwitch(meses)

	//------------------EJERCICIO 1 PROFE-----------------//
	// letras de una palabra
	fmt.Println("------CON EL PROFE --------")
	var palabra string = "palabraOpcion1"
	//palabra2 := "palabraOpcion2"

	cant := len(palabra)
	fmt.Println( "la palabra", palabra, "tiene ", cant, "letras")

	for i := 0; i < cant; i++ {
		fmt.Println(string(palabra[i])) // casteo 
	}
	//------------------EJERCICIO 2 PROFE -------------------//
	// A qué mes corresponde

	mesesProfe := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre" }
	mes := 3

	if mes >= 1 && mes <= 12{
		for i := 0; i <= len(mesesProfe); i++ {
			if i== mes{
				fmt.Println(mesesProfe[i-1])
			}
	}}

	fmt.Println(mesesProfe[mes-1])

	//------------MESA DE TRABAJO---------//

	fmt.Println("-------MESA DE TRABAJO---------")

	//Ejercicio 1 - Listado de nombres
//Una profesora de la universidad quiere tener un listado con todos sus estudiantes. 
//Es necesario crear una aplicación que contenga dicha lista. Sus estudiantes son: 
//Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernán, Leandro, Eduardo, Duvraschka.
//Luego de dos clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado, 
//sin modificar el código que escribiste inicialmente. El nombre de la nueva estudiante es Gabriela.

estudiantes := [] string {"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"} 

fmt.Println(estudiantes)

estudiantes = append(estudiantes, "Gabriela")

fmt.Println(estudiantes)

//Ejercicio 2 - Qué edad tiene...
//Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. 
//Según el siguiente map, debemos imprimir la edad de Benjamin.
var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
//Por otro lado, también es necesario:
//Saber cuántos de sus empleados son mayores de 21 años.
//Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
//Eliminar a Pedro del map.

fmt.Println(employees)
fmt.Println(employees["Benjamin"])
var cuantos int = 0
for llave, valor := range employees {
	
	if  valor > 21 {
		fmt.Println(llave, "Es mayor de 21 años")
		cuantos ++
	}
}
fmt.Println("Tiene ", cuantos, "empleados mayores de 21 años")

employees ["Federico"] = 25
fmt.Println("Agregando  un nuevo empleado",employees)

delete(employees,"Pedro")
fmt.Println("Eliminando un empleado ", employees)
}



//-----------------------ASINCRONICO---------------------------------//
// Usando ciclo for
func ObtenerEstacion(meses []string)  {
	for _, mes := range meses{
		if mes == "Enero" || mes == "Febrero" || mes == "Marzo" {
			fmt.Println("Verano")
		}else if mes == "Abril" || mes == "Mayo" || mes == "Junio" {
			fmt.Println("Otoño")
		}else if mes == "Julio" || mes == "Agosto" || mes == "Septiembre" {
			fmt.Println("Invierno")
		}else if mes == "Octubre" || mes == "Noviembre" || mes == "Diciembre"{
			fmt.Println("Primavera")
		}else{
			fmt.Println("No existe ese mes")
		}
	}
}
//Usando swtich
func ObtenerEstacionSwitch (meses []string){
for _, mes := range meses{
	switch mes {
	case  "Enero", "Febrero",  "Marzo":
		fmt.Println("Verano - Switch")
	case "Abril", "Mayo", "Junio":
		fmt.Println("Ontoño -Switch")
	case "Julio", "Agosto", "Septiembre":
		fmt.Println("Invierno -Switch")
	case "Octubre", "Noviembre", "Diciembre":
		fmt.Println("Primavera - Switch")
	default: fmt.Println("No existe el mes Swtch")
	}
}
}