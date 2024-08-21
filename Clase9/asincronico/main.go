package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//Crear un router Gin
	router := gin.Default()

	//Captura la sollicitud GET "/hello-word"
	router.GET("hello-word", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "HELLO WORD",
		})
	})

	//Corremos en nuestro servidor sobnre el puerto 8080
	router.Run()
}