package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

/*Una empresa necesita un sistema para gestionar sus empleados. Para poder hacer esto,
necesitan un servidor que ejecute una API que les permita manipular a los empleados.

Ejercicio 1: Iniciando el proyecto

Debemos crear un repositorio en GitHub para poder subir nuestros avances. Este repositorio
es el que vamos a utilizar para contener todo lo que realicemos durante las distintas prácticas de Go Web.
Primero debemos clonar el repositorio creado. Luego, iniciar nuestro proyecto de Go con el comando go mod init.
El siguiente paso será crear un archivo main.go donde deberán cargar un slice, desde una
función que devuelva una lista de empleados. Este slice se debe cargar cada vez que se inicie la API
para realizar las distintas consultas.
La estructura de los empleados es la siguiente:
Id		 int
Nombre	 string
Activo	 bool

Ejercicio 2: Empleados

Vamos a levantar un servidor utilizando el paquete Gin en el puerto 8080.
Para probar nuestros endpoints, haremos uso de Postman.

Crear una ruta / que nos devuelva una bienvenida al sistema. Ejemplo: “¡Bienvenido a la empresa Gophers!”.
Crear una ruta /employees que nos devuelva la lista de todos los empleados en formato JSON.
Crear una ruta /employees/:id que nos devuelva un empleado por su ID. Manejar el caso de que no encuentre el empleado con ese ID.
Crear una ruta /employeesparams que nos permita crear un empleado a través de los params y lo devuelva en formato JSON.
Crear una ruta /employeesactive que nos devuelva la lista de empleados activos. También debería poder devolver la lista de los empleados no activos.
*/

type Empleado struct{
	Id		 int
	Nombre	 string
	Activo	 bool
}

//ListaEmpleados es mi base de datos harcode y devuelva un slice de empleados
func ListaEmpleados()[]Empleado{
	empleados := []Empleado{
		{Id: 1,Nombre: "Sofia",Activo: true},
		{Id: 2,Nombre: "Carlos", Activo: false},
		{Id: 3,Nombre: "Marta", Activo: false},
		{Id: 4,Nombre: "Roberto", Activo: true},
		{Id: 5,Nombre: "Lorena", Activo: true},
	}
	return empleados
}

func main()  {
	empleados :=ListaEmpleados()
	router := gin.Default()

//Bienvenida
router.GET("/",func(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,gin.H{
		"msg": "¡Bienvenido a la empresa Gophers!",
	})
})

//trae todos los empleados
router.GET("/employees",func(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, empleados)
})

router.GET("/employees/:id", BuscarEmpleado())
router.POST("/employeesparams", CrearEmpleado())

//trae los empleados que se encuentren en el estado solicitado en el contexto 
router.GET("/employeesactive",func(ctx *gin.Context) {
	activoParam := ctx.Query("activo")
	activo, err := strconv.ParseBool(activoParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"error": "Ninguno de los empeleados se encuentra en este estado"})
		return
	}
	estatus := []Empleado{}

	for _, empleado := range empleados{
		if activo == empleado.Activo {
			estatus =append(estatus, empleado)
		}
	}
	ctx.JSON(http.StatusOK,estatus)
})

router.Run()
}

//  CrearEmpleado agrega un empleado recibe desde el contexto los valores id, nombre, activo
// y devuelve a todos los empleados 
func CrearEmpleado()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		
		idParam := ctx.Query("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}
		nombre := ctx.Query("nombre")
		
		activoParam := ctx.Query("activo")
		activo,err := strconv.ParseBool(activoParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Estado Activo es invalido"})
			return
		}

		e := Empleado{
			Id: id,
			Nombre: nombre,
			Activo: activo,
		}
		empleados := append(ListaEmpleados(), e)

		ctx.JSON(http.StatusOK, empleados)
		

	}
}

//BuscarEmpleado busca el empleado por el param id
func BuscarEmpleado()gin.HandlerFunc{
	return func (ctx *gin.Context){
		empleados :=ListaEmpleados()
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		for _, empleado := range empleados{
			if empleado.Id == id{
				ctx.JSON(http.StatusOK, empleado)
			}
		}

		ctx.JSON(http.StatusNotFound,gin.H{"error": "Empleado no encontrado"})
	}
}
