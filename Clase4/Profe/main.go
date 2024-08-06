package main

import "fmt"

func main() {
	var a int
	a = 10
	fmt.Println("Valor de  a es: ",a," el espacio de memoria: ",&a)
	calc(&a)
	fmt.Println("La func calc devuelve a convertido ",a)

	//un puntero que apunta al mismo espacio de memoria que calc
	var ptr *int
	ptr = &a
	fmt.Println(ptr)
	*ptr =90
	fmt.Println("\nPuntero ptr: ",*ptr," el espacio de memoria: ",ptr)


}

//usando un punturo por primera vez, el valor apuntara a la varianble que se le pase por parametro
func calc(valor *int){
	fmt.Println("Lo que llega por parametro a la func calc ",valor," y el valor dereferenciado ",*valor)
	//Cambia el valor en espacio de memoria especifico
	*valor =20
}