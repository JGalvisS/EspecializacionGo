package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

/*
Verbo utilizado: GET, POST, PUT, etc
Fechat y hora: podemos utilizar el paquete time
URL de la consulta: localhost:8080/products
Tamaño de bytes: tamaño de la consulta*/

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verb := ctx.Request.Method // Request.Method trae el metodo que viene por contexto en la solicitud
		time := time.Now() // Time es un paquete que nos ayuda  a usar las variables de tiempo actual o en ejecucion 
		url := ctx.Request.URL //Request.URL trae la url que viene por contexto 
		ctx.Next() //Cargue los valores y siga con la ejecucion cuando regrese con la respuesta cargo el tamaño de los valores
		//NOTA: Se busca que la respueste pese lo menos posible.
		var size = 0
		if ctx.Writer != nil{
		size = ctx.Writer.Size() //Witer me permite escribirle una determinada respuesta al cliente y tiene un metodo Size obtiene el tamaño en bytes 
		}
		fmt.Printf("verb: %s\nTime: %v\nURL: %v\nSize: %d\n", verb,time,url,size)//%v es para escribir un generico 
		ctx.Next()
	}
}