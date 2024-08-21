package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

/*Vamos a crear una aplicacion web con el framework Gin  que tenga un endpoint /persona
que responda con los datos  de la persona
1. Crear una estructura persona con los valores:
Nombre
Apellido
Edad
Direccion
Telefono
Activo
2. Crear una persona en formato JSON(hardcodear) y aplicar la logica para que:
a. Imprima por consola
b. Imprima la persona a traves  del enpoint en formato JSON. El endpoint debera ser de metodo GET
*/


type Perosona struct{
	Nombre		string
	Apellido	string
	Edad		int
	Direccion	string
	Telefono	string
	Activo		bool
}

func main() {
	//Primero ejegutamos a gin por defecto
	router := gin.Default()

	//Iniacilizo una Persona
	var p Perosona

	//Simulo que recibo una data de la request 
	jsonData := `{"Nombre":"Jessica", "Apellido":"Galvis", "Edad": 98, "Direccion":"Calle falsa 123", "telefono":"526656", "Activo": true}`

	//Parseo la resquest
	if err := json.Unmarshal([]byte(jsonData), &p); err != nil{
		log.Fatal(err)}
	// Resuelvo el punto a 
	fmt.Println(p)


	//router.GET Para usar el metodo GET  en la ruta /persona
	//ctx permite manejar la entrada y la salida 
	router.GET("/persona",func(ctx *gin.Context) {
		//ctx.JSON devueleve una respuesta en formato JSON que necesita un status code y un objeto que en este caso sera persona
		ctx.JSON(http.StatusOK, gin.H{
			"Persona": p,
		})
	})


	//Por defecto lo corre en el puerto 8080
	router.Run()
}


