package main

import (
	"errors"
	"fmt"
	"os"
)

//EJERCICIO 1
/*En la funcion main, definir una variable salary y asignarle un valor de
tipo int. Crear in errors.New() con mensaje "salario muy bajo" y lanzarlo
en caso de que salary sea nenor o igual a 10.000. La validacion debe ser
hecha con la funcion Is() dentro de main.*/

//EJERCICIO 2 - Datos de clientes
/* Un estudio contable necesita acceder a los datos de sus empleados para poder
realizar distintas liquidaciones. Para ello, cuentan con todo el detalle necesario
en un archivo TXT.
Desarrollar el código necesario para leer los datos de un archivo llamado “customers.txt”.
Sin embargo, debemos tener en cuenta que la empresa no nos ha pasado el archivo a leer por
el programa. Dado que no contamos con el archivo necesario, se obtendrá un error. En tal caso,
el programa deberá arrojar un panic al intentar leer un archivo que no existe, mostrando el mensaje
“El archivo indicado no fue encontrado o está dañado”.Más allá de eso, deberá siempre imprimirse por
consola “Ejecución finalizada”.*/
var errSalary =  errors.New("salario muy bajo")
func esSalarioBajo(salario int)error{
	if salario <= 10000{
		return errSalary
	}
	return nil
}

func main(){
	//EJERCICIO 1
	salary :=9000
	err:=esSalarioBajo(salary)
	fmt.Println("ejecutando esSalarioBajo: ",err)
	// Is():en la ejecucion compara la respuesta almacendada en err es 
	// igual al  target (el segundo parametro) retorna un bool
	if errors.Is(err,errSalary){
		fmt.Println("Validando con Is() un error: ", err)
	}

	//EJERCICIO 2

	fmt.Println("Iniciando...")
	readFile("customers.txt")
	fmt.Println("Ejecucion finalizada")
}
// readFile sirve para abrir un archivo .txt o .xlsx
func readFile(routeName string){
	// defer: posponer la ejecución de una función hasta que la función que
	// la contiene haya terminado.
	// recover: permite capturar un panic y evitar la terminación del programa 
	defer func ()  {
		err := recover()
		if err != nil{
			fmt.Println(err)
		}
	}()
	
	_,err := os.Open(routeName)
	if err != nil {
		//panic: detiene todo abruptamente y comienza a desempaquetar la pila de llamadas (es decir, ejecuta las funciones diferidas). 
		//Sino se captura el pánico con recover, el programa terminará abruptamente.
		panic("El archivo indicado no fue encontrado o está dañado")
	}
}
