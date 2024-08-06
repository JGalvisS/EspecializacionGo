package main

import "fmt"
//-----------EJERCICIO 1-----------------
//la Compañia es una estructura o "clase"
type Compania struct {
	Nombre string
	Puesto string
}
//Empleado es otra estructura que tiene una relacion de composicion con Compañia
type Empleado struct {
	Nombre   string
	Apellido string
	Compania Compania
}
// Esta funcion imprime en consola a un empleado
func (e Empleado) Informacion() {
	fmt.Printf("Empleado %s %s\nPuesto: %s\nCompania: %s", e.Nombre, e.Apellido, e.Compania.Puesto, e.Compania.Nombre)
}
//Esta funcion actualiza el puesto del empleado en la compañia
//usando un puntero "*" que permite que cada vez se modifique el espcio en memoria 
func (c *Compania) CambiarPuesto(nuevoPuesto string){
	c.Puesto = nuevoPuesto
}

//------------------EJERCICIO 2 ------------------------
/*Una empresa de redes sociales requiere implementar una estructura usuarios con funciones 
que vayan agregando información a la misma. Para optimizar y ahorrar memoria requieren que 
la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones. 
La estructura debe tener los campos: nombre, apellido, edad, correo y contraseña. 
Y deben implementarse las funciones:
cambiarNombre: permite cambiar el nombre y apellido.
cambiarEdad: permite cambiar la edad.
cambiarCorreo: permite cambiar el correo.
cambiarContraseña: permite cambiar la contraseña.*/

type Usuario struct{
	Nombre 		string
	Apellido 	string
	Edad		int
	Correo		string
	Password	string
}


//Implemando una interfaz
type UsuariosOps interface{
	cambiarNombre(nombre, apellido string)
	cambiarEdad(edad int)
	cambiarCorreo(correo string)
	cambiarPassword(password string)
	detail()
}

func (u *Usuario) cambiarNombre(nombre string, apellido string){
	u.Nombre = nombre
	u.Apellido = apellido
}

func (u *Usuario) cambiarEdad(edad int){
	u.Edad = edad
}

func (u *Usuario) cambiarCorreo(correo string){
	u.Correo = correo
}

func (u *Usuario) cambiarPassword(password string){
	u.Password = password
}

func (u Usuario) detail(){
	fmt.Printf("Ususario: %s %s\nEdad: %d\nCorreo: %s\nPassword: %s\n", 
	u.Nombre, u.Apellido, u.Edad, u.Correo, u.Password)
}

func main() {
	//----------------EJERCICIO 1-------------
	//Inicializo el objeto empleado 
	empleado := Empleado{
		Nombre: "Carlos",
		Apellido: "Ramirez",
		Compania: Compania{
			Nombre: "IT Solution",
			Puesto: "Programador Junior",
		},
	}

	//----Usando las funciones-------
	fmt.Println("---------EJERCICIO 1------------")
	empleado.Informacion()
	fmt.Println()
	fmt.Println()
	empleado.Compania.CambiarPuesto("Tester")
	empleado.Informacion()
	fmt.Println()
	fmt.Println()
	//Tambien puedo hacerlo de esta forma
	empleado.Compania.CambiarPuesto("Programador Front")
	empleado.Informacion()
	fmt.Println()

	//----------EJERCICIO 2------------------
	//Instancio un usuario, usando "&" creo una nueva instancia 
	usuario := &Usuario{
		Nombre: "Enrrique",
		Apellido: "Velez",
		Edad: 52,
		Correo: "enrri98@gmail.com",
		Password: "enrri123456",
	}

	var ops UsuariosOps = usuario
	

	fmt.Println()
	// Usando las funciones en mi interfaz para cambiar la informacion del usuario
	fmt.Println("---------EJERCICIO 2---------")
	ops.detail()
	ops.cambiarNombre("Maria", "Perez")
	ops.cambiarEdad(30)
	ops.cambiarCorreo("mariap85@gmail.com")
	ops.cambiarPassword("mar3085@gmail.com")
	fmt.Println()
	ops.detail()
	
}