package main

import (
	"errors"
	"fmt"
	"log"
)

//------------CON EL PROFE-------------//

const (
	min = "minimo"
	avg = "pronedio"
	max = "maximo"
)

func main() {
	//EJERCICIO :Establecer la mejor calificacion, la menor y el promedio

	// pruebo la funcion minimo: en la variable minFun voy a recibir el valor de la funcion calculoMinimo a traves de operacion
	// y en la variable err voy a recibir el valor de error de la funcion operacion 
	minFun, err := operacion(min)
	if err != nil{ log.Fatal(err)}
	fmt.Println("El minimo es: ", minFun(9,8,9,7,9,6,5,3,7))


	//pruebo la funcion max
	maxFun, err := operacion(max)
	if err != nil{ log.Fatal(err)}
	fmt.Println("El maximo es: ", maxFun(9,8,9,7,9,6,5,3,7))

	//pruebo la funcion promedio
	avgFun, err := operacion(avg)
	if err != nil{ log.Fatal(err)}
	fmt.Println("El promedio es: ", avgFun(9,8,9,7,9,6,5,3,7))

	//pruebo multi operaciones
	minimo, promedio, maximo, err := operacionesMulti()
	if err != nil{ log.Fatal(err)}
	fmt.Println("El minimo es: ", minimo(9,8,9,7,9,6,5,3,7), ". El promedio es: ", promedio(9,8,9,7,9,6,5,3,7), ". El maximo es: ", maximo(9,8,9,7,9,6,5,3,7) )

	//pruebo saltar error
	errorFun, err := operacion("cualquier Cosa")
	if err != nil{ log.Fatal(err)}
	fmt.Println("Esta es la prueba del error: ", errorFun(9,8,9,7,9,6,5,3,7))

}
//esta funcion dependiendo del string llamara a la funcion requerida,
//cada funcion recibe un slice y devuelve un entero y un error
func operacion(calculo string) (func (...int) int, error) {
	switch calculo {
	case min:
		return calculoMinimo, nil
	case max:
		return caluloMaximo, nil
	case avg:
		return calculoPromedio,nil
	default:
		return nil, errors.New("El calculo no existe.")
	}
}
//Esta funcion determina el valor minimo en el slice
func calculoMinimo(values ...int) int{
	min := values[0]

	//Hacer un forEach, al no necitar la llave que en este caso seria el indice usamos _,
	//especificamos que solo usaremos el valor contenido en cada posicion del slice
	for _, valor := range values{
		if valor < min{
			min = valor
		}
	}
	return min
}

//Esta funcion determina el valor maximo en el slice
func caluloMaximo(values ...int) int{
	max := values[0]
	for _, valor := range values{
		if valor > max{
			max = valor
		}
	}
	return max
}

//Esta funcion determina el valor promedio en el slice
func calculoPromedio(values ...int) int {
	sum := 0
	for _, valor := range values{
		sum += valor 
	}
	return sum / len(values)
}

func operacionesMulti()(func(...int)int, func(...int)int, func(...int)int, error){
	return calculoMinimo, calculoPromedio, caluloMaximo, nil
}