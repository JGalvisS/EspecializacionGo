package main

import (
	"fmt"
	"net/http"
)

func main() {
	//importo el paquete con net/http que esta de forma nativa sin bibliotecas
	//HandlerFunc permite manejar las peticiones http con una respuesta, 
	//"/hola" establece la ruta y holahandler es el manejador de esa request
	http.HandleFunc("/hola", holahandler)

	//Escucha y sirve en el puerto 8080 y el nil es porque vamos a usar nuestro propio handler 
	http.ListenAndServe(":8080", nil)
}

//La funcion recibe como parametro un ResponseWriter, eata escribe el cuerpo de la respuesta que se enviará al cliente y definir un status code de la respuesta http
//El segundo parametro es un puntero de http.Request proporciona toda la información sobre la solicitud entrante, por ejemplo el tipo de metodo
func holahandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "GET"{
		w.WriteHeader(http.StatusUnauthorized)
		return//corta la ejecucion
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"Hello word") //Especial para trabajar ResponseWriter
}