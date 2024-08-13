package main

import (
	"fmt"
	"log"
	"os"
)

/*Ejercicio - Mayorista
Se necesita un software desarrollado en Go que imprima las estimaciones de lo que generarían,
los productos de cada categoría en un mayorista, en ventas de productos para el hogar.
Para eso se detallarán los pasos para conseguirlo:

1.Funcionalidad para generar un archivo CSV, llamado categorías.csv.
2. Cargar el archivo con las categorías. Ejemplo:
		001	Electrodomésticos	ListaProductos
		002	Muebles				ListaProductos
		003	Herramientas		ListaProductos
		004	Pinturas			ListaProductos
		005	Aberturas			ListaProductos
		006	Construcción		ListaProductos
		007	Automotor 			ListaProductos
		Etcétera…
	Elegir al menos tres categorías.

3. Generar para cada una de estas categorías los productos. 
Estos tendrán como información:
	* Código
	* Nombre
	* PrecioActual
	* CantidadActual
Insertar al menos cuatro productos.
4 Generar otro archivo CSV, llamado estimaciones.csv. Este tendrá los resultados de la suma de todos los productos de cada una de las categorías.
5 Imprimir todos los estimativos por consola. Ejemplo:

Categoría			EstimativoPorCategoría
Construcción				60.700
Pinturas 					40.500
Aberturas					55.300
TotalEstimativo 			156.500

*/


func main(){
	generarCategoriaFile("categorias")

	addCategoria(001,"Electrodomesticos")
	addCategoria(002,"Muebles")
	addCategoria(003,"Herramientas")
	addCategoria(004,"Pinturas")
	addCategoria(005,"Aberturas")

	addPrducto("Electrodomesticos","E01","Microondas",1500.00,10)
	addPrducto("Electrodomesticos","E02","Lavadora",3000.00,5)
	addPrducto("Electrodomesticos","E03","Televisor",2000,15)
	addPrducto("Electrodomesticos","E04","Nevera",3500.00,3)

	addPrducto("Muebles","E01","Silla",150.00,20)
	addPrducto("Muebles","E02","Mesa",300.00,15)
	addPrducto("Muebles","E03","Sofa",280,5)
	addPrducto("Muebles","E04","Cama",350.00,6)

}

//PASO 1 - Generar archivo csv
func generarCategoriaFile(name string){
	filename := "./"+name+".csv"
	//Validacion de que no exista
	_, err  := os.Stat(filename)
	if os.IsNotExist(err) {
		//Creo archivo
		f, err2 := os.Create(filename)
		if err2 != nil {
			log.Fatal(err2)
		}

		//Creo los titulos de las columnas
		_, err1 := f.WriteString("ID,Categoria,Lista_Productos;")
		if err1 != nil {
			log.Fatal(err1)
		}else{ fmt.Println("El archivo ",filename, " fue creado con exito")}
	}else{
		fmt.Print("El archivo ",filename, " ya existe")
	}
} 

//PASO 2 -  Cargar el archivo con las categorías
func addCategoria(id int, categoria string){
	listPrducts := "./"+categoria+".csv"
	//Abro el archivo, os.O_APPEND para agregar sin sobrescribir, os.O_WRONLY abre en modo escritura
	f, err := os.OpenFile("./categorias.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	//Creo esa categoria//

	//valido que no exista esa categoria
	_, err1  := os.Stat(listPrducts)
	if os.IsNotExist(err1) {

		//Agrego las Categorias
		_, err = f.WriteString(fmt.Sprintf("%d,%s,%s;",id,categoria,listPrducts))
		if err != nil {
			log.Fatal(err)
		}

		//Creo archivo categoria
		categoryfile, err2 := os.Create(listPrducts)
		if err2 != nil {
			log.Fatal(err2)
		}

		//Creo los titulos de las columnas
		_,err3 := categoryfile.WriteString("Codigo,Nombre,PrecioActual,CantidadActual;")
		if err3!= nil {
			log.Fatal(err3)
		}
		fmt.Println("La categoria ", listPrducts," se agrego con exito")
	}else if err == nil {
		fmt.Println("La categoria ya existe")
	}
}

//PASO 3 - Generar para cada una de estas categorías los productos
func addPrducto(categoria string, codigo string, nombre string, precio float64, stok float64){
	prodCategory := "./"+categoria+".csv"
	//Valido que exista esa categoria
	_,err := os.Stat(prodCategory)
	if err == nil {
		//Abro el archivo, os.O_APPEND para agregar sin sobrescribir, os.O_WRONLY abre en modo escritura
		f,err1 := os.OpenFile(prodCategory, os.O_APPEND|os.O_WRONLY,0644)
		if err1 != nil {
			log.Fatal(err1)
		}
		//Agrego el producto y uso Sprinttf para que sea un string
		_, err1 = f.WriteString(fmt.Sprintf("%s,%s,%.2f,%.2f;",codigo,nombre,precio,stok))
		if err1 != nil {
			log.Fatal(err1)
		}
		fmt.Println("El producto ",nombre," se guardo con exito en la categoria",categoria)
	}else if os.IsNotExist(err){
		fmt.Println("La categoria no fue encontrada. Crea una nueva categoria para este producto")
	}else{
		log.Fatal(err)
	}
}

