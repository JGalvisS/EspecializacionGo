package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*La empresa encargada de vender los productos de limpienza necesita leer ahora el archivo
almacenado. Para esto requiere que se imprima por pantalla mostrando los valores
tabulados, con un titulo (tabulado a ;a izquierda para el ID y a la derecha para el Precio y
Cantidad), el precio la cantidad y abajo del precio se debe visualizar el total (sumando
precio por cantidad ) */

func main(){
	//Leyendo el archivo
	res,err := os.ReadFile("./data.csv")
	if err != nil{
		fmt.Println(err)
		os.Exit(2)//Para que termine la ejecucion del if y sigue a la siguiente instruccion almacenada en esta funcion 
	}
	fmt.Println("Este es el contenido del archivo: ")
	fmt.Println(string (res))//imprimo la respuesta convirtiendo byte a string

	//Dividiendo la respuesta por filas, convirtiendolo en slice de strings 
	data := strings.Split(string(res), ";")
	fmt.Println("Creando un slice del archivo")
	fmt.Println(data)

	

	//sumando el precio por cantidad 
	var total float64
	fmt.Printf("%-10s %10s %10s\n", "ID", "Precio", "Cantidad")//Encabezado de la tabla, el formato %-10s alinea el texto a la izquierda y %-10s lo alinea a la derecha.
	for i := 0;i <len(data)-1; i++ {
		line := strings.Split(string(data[i]), ",")
		if i != 0 {
			//Obteniendo el valor del precio
			//strvonvert permite convertir un string en determinado tipo de dato
			precio, err := strconv.ParseFloat(line[1],64) //usa como bitsize 64 porque es el float que venimos usando
			if err != nil{
				fmt.Println("El precio no se pudo obtener")
			}

			//Obteniendo la cantidad 
			cant, err := strconv.ParseFloat(line[2],64)
			if err != nil{
				fmt.Println("La cantidad no se pudo obtener")
			}
			multiplica := precio * cant

			total = total + multiplica
			fmt.Printf("%-10s %10.2f %10.2f\n", line[0], precio, cant)
		}
	}
	fmt.Println("-------------------------------------")
	fmt.Printf("%-10s %10.2f\n", "Total:", total)

}