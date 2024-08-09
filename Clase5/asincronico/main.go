package main

import "fmt"

func main() {
	fmt.Println("La funcion esta iniciando")
	ejemplo()
	fmt.Println("La funcion esta finalizando")
}
//tiene un slice de string y devuelve el valor 3 de ese slice
func ejemplo(){
/*// asi se genera un panic sin manejo
	animales:=[]string{
		"vaca",
		"gato",
		"perro",
		"paloma",
	}
	fmt.Println("El animal que vuela es: ", animales[5])*/

	//manejo de panic
	panic("Ocurrio un error")
}