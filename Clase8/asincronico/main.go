package main

import "fmt"

type Bootcamp struct {
	Organizacion string
	Lenguaje     string
	Alumnos      []Alumno
}

type Alumno struct {
	Nombre       string
	FechaIngreso string
}

// Esta funcion instancia un Bootcamp sin alumnos
func crearBootcamp(organizacion string, lenguaje string) Bootcamp {
	boot := Bootcamp{
		Organizacion: organizacion,
		Lenguaje:     lenguaje,
	}
	return boot
}

// Este metodo sirve para agregar al slice de Alumnos un alumno nuevo
func (b *Bootcamp) guardaAlumno(alumno Alumno) {
	b.Alumnos = append(b.Alumnos, alumno)
}

//Esta funcion sirve para crear un Alumno
func crearAlumno(nombre string,fechaIgreso string)Alumno {
	alumno:= Alumno{
		Nombre: nombre,
		FechaIngreso: fechaIgreso,
	}
	return alumno
}

func main() {
	bootcamp := crearBootcamp("digital house", "Go")
	fmt.Println("Nuevo bootcamp: ",bootcamp)
	alm1 :=crearAlumno("Jessica","08/14/24")
	alm2 :=crearAlumno("Lolita","08/14/24")
	bootcamp.guardaAlumno(alm1)
	bootcamp.guardaAlumno(alm2)
	fmt.Println("Agregando alumnos al bootcamp: ",bootcamp)
	


}