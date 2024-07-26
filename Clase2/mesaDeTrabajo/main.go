package main

import (
	"errors"
	"fmt"
	"log"
)

//Ejercicio 1 - Calcular salario
//Una empresa marinera necesita calcular el salario de sus empleados basándose en la
//cantidad de horas trabajadas por mes y la categoría.
//Categoría C: su salario es de $1.000 por hora.
//Categoría B: su salario es de $1.500 por hora, más un 20 % de su salario mensual.
//Categoría A: su salario es de $3.000 por hora, más un 50 % de su salario mensual.
//Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados
//por mes y la categoría, y que devuelva su salario.
const (
	cA = "Categoria A"
	cB = "Categoria B"
	cC = "Categoria C"
)

func main()  {
	//Pruebo categoria C
	totalC, err := salario(cC, 60)
	if err != nil{ log.Fatal(err)}
	fmt.Println("El salario del empleado con catergoria C es: $",totalC)

	//Pruebo categoria B
	totalB, err := salario(cB, 60)
	if err != nil{ log.Fatal(err)}
	fmt.Println("El salario del empleado con catergoria B es: $",totalB)


	//Pruebo categoria A
	totalA, err := salario(cA, 60)
	if err != nil{ log.Fatal(err)}
	fmt.Println("El salario del empleado con catergoria B es: $",totalA)


	//Pruebo error
	totalerr, err := salario("cualquier cosa", 60)
	if err != nil{ log.Fatal(err)}
	fmt.Println("El salario del empleado con catergoria B es: $",totalerr)
	
}
//Esta funcion sirve para calcular el valor del minuto de trabajo
func valorMinuto(valorHora int) float64  {
	return float64(valorHora) / 60
}
//Esta cuncion calcula el valor del salario de cada empleado
func salario(categoria string,minutos int)(int, error){
	switch categoria {
	case cC:
		valorHora := 1000
		valorMin := valorMinuto(valorHora)
		salario := valorMin * float64(minutos)
		return int(salario), nil
	case cB:
		valorHora := 1500
		valorMin := valorMinuto(valorHora)
		salario := valorMin * float64(minutos)
		extra := salario * 0.20
		return int(salario + extra), nil
	case cA :
		valorHora := 3000
		valorMin := valorMinuto(valorHora)
		salario := valorMin * float64(minutos)
		extra := salario * 0.50
		return int(salario + extra), nil
	default:
		return  0, errors.New("La categoria no existe.")
	}
}
