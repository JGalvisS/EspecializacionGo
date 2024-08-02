package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*Una empresa necesita realizar una buena gestion de sus
empleados, para esto realixaremos un pequeño programa que nos
ayudara correctamente dichos empleados:

1. Crear una estructura Person con los campos ID, Name, DateOfBirth.
2 Crear una estructura Employee con los campos, ID,Position y una composición
con la estructura Person.
3.Realizar un metodo a la estructura Employee que se llame PrintEmployee(),lo
que hara es realizar la impresion de los campos empleado
4. Iniciar en la funcion main()tanto para Person como un Employee cargando respectivamente
camppos y por ultimo ejecutar Print Employee()*/

//Agregamos etiquetas a cada uno de los campo para que cuando se parsee los campos lleven esas etiquetas
type Person struct{
	Id 			int			`json:"dni"`
	Name 		string		`json:"name"`
	DateOfBirth string		`json:"date_of_birth"`
}

type Employee struct{
	Id 			int			`json:"id"`
	Position 	string		`json:"position"`
	Person 		Person		`json:"person"`
}

func main()  {
	p1 := Person{115655, "Jose", "10/10/1998"}
	e1 :=Employee{1,"Desarrollador",p1}

	fmt.Printf(e1.PrintEmployee())
	fmt.Println()

	//----------Llamando etiquetas----------
	// Convierto objetos de una estructura en json
	employee, err := json.Marshal(e1)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Println("Imprimiendo el objeto convertido a json")
	fmt.Println(string(employee))
	fmt.Println()

	//Convertir un json en un objeto
	var e2 Employee
	employee2 := "{\"id\":2,\"position\":\"Desarrollador\",\"person\":{\"dni\":1156565,\"name\":\"Sandra\",\"date_of_birth\":\"10/10/1993\"}}"

	if err2 := json.Unmarshal([]byte(employee2), &e2); err2!= nil{
		log.Fatal(err2)
	}
	fmt.Println("Imprimiendo el json convertido en objeto " )
	fmt.Println(e2.PrintEmployee())

}
func (e Employee)PrintEmployee()string{
	//Sprinttf parsea a string 
	return fmt.Sprintf("Employee: %d\nName: %s\nPosition: %s\nId: %d\nDate of birth: %s", e.Id, e.Person.Name, e.Position, e.Person.Id, e.Person.DateOfBirth)

}