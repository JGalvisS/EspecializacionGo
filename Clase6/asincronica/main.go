package main

import (
	"fmt"
	"log"
	"os"
)

/*Una empresa que se encarga de vender productos de limpieza necesita:
*1 Implementar una funcionalidad para guardar un archivo de texto con la
información de productos comprados, separados por punto y coma (CSV).
*2 Este archivo debe tener el ID del producto, precio y la cantidad.
*3 Estos valores pueden ser hardcodeados o escritos en duro en una variable.*/
func main(){
crearProdCompradosFile()
addProdComprado(1,36.50,9)
}

//Esta funcion crea el archivo de inventario
func crearProdCompradosFile(){
	//Valido sí el archivo inventario existe
	_, err3 := os.Stat("./inventario.csv")
	//Si el error indica que el archivo no existe, lo creo
	if os.IsNotExist(err3){

		//Creo el archivo inventario 
		f,err := os.Create("./inventario.csv")
		if err != nil{
			log.Fatal(err)
		}

		//Crea los titulos de las columnas
		_, err2 := f.WriteString("ID,Precio,Cantidad;\n")
		if err2 != nil{
			log.Fatal(err2)
		} 
		fmt.Println("El archivo Inventario ha sido creado con exito")
	}else if err3 == nil{
		fmt.Println("El archivo inventario ya existe")
	}
	
}
//Esta funcion agrega un producto
func addProdComprado(id int, precio float64, cantidad int){
	//Abro el archivo, os.O_APPEND para agregar sin sobrescribir, os.O_WRONLY abre en modo escritura
	f, err := os.OpenFile("./inventario.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//Agrego el producto y uso Sprinttf para que sea un string
	_, err = f.WriteString(fmt.Sprintf("%d,%.2f,%d;\n", id, precio, cantidad))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("El producto ha ha sido agregado al inventario con exito")
}
	
