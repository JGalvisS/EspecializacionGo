package main

import "fmt"

/*Una empresa de inteligencia artificial necesita tener una funcionalidad para crear matrices en Go.
Para ello  requieren una estructura Matrix que tenga los metodos:
Set: recibe una matriz de flotantes e inicializa los valores en la estructura Matrix
Print: imprime por pantalla la matriz de una forma mas visible (con los saltos de linea entre filas).
La estructura Matrix debe contener los valores de la matriz, la dimension del alto, la dimension del ancho,
si es cuadratica y cual es el valor maximo*/

type Matrix struct{
	valores []float64
	alto int
	ancho int
}

func main(){
	unaMatrix := Matrix{
		valores: []float64{},
		alto: 3,
		ancho: 3,
	}
	fmt.Println("----------inicializa la matrix-------------")
	fmt.Println(unaMatrix)

	fmt.Println("----------usando Set-------------")
	valores := []float64{1,2,3,4,5,6,7,8,9}
	unaMatrix.Set(valores...)
	fmt.Println(unaMatrix)
	fmt.Println("----------usando Print-------------")
	unaMatrix.Print()
	unaMatrix.valorMax()
	fmt.Println("Es cuadratica?",unaMatrix.cuadratica())

}

// Esta funcion recibe una matrix que tiene un puntero es decir usara el mismo espacio de memoria, 
//la linea "valores ...float" indica que recibe por parametro slice de floatantes
func (m *Matrix) Set(valores ...float64){
	//validacion de que coincida con el ancho y alto deseado 
	if len(valores) != m.alto*m.ancho{
		fmt.Println("La cantidad de valores no coincide con el alto y con el ancho")
		return//para cortar el flujo
	}
	m.valores = append(m.valores, valores...)
}

func (m Matrix) Print(){//no uso puntero porque no estoy modificando ningun valor 
	for fila := 0; fila <m.ancho; fila++{
		fmt.Println(m.valores[fila*m.ancho : fila*m.ancho+m.ancho])
	}
	/* Como recorre for
	cuando fila = 0 *ancho3 = 0 [0:3] luego vuelve y multiplica fila = 0* amcho3 = 0 +ancho3 =3
	cuando fila = 1 *ancho3 = 3 [3:6] luego vuelve y multiplica fila = 1* amcho3 = 3 +ancho3 =6
	cuando fila = 2 *ancho3 = 6 [6:9] luego vuelve y multiplica fila = 2* amcho3 = 6 +ancho3 =9
	*/
}
func (m Matrix) valorMax(){
	max := m.valores[0]
	for _, elemento := range m.valores{
		if elemento > max {
			max = elemento
		}
	}
	fmt.Println("El valor maximo es: ",max )
}

func(m Matrix) cuadratica()bool{
	if m.alto == m.ancho && m.alto != 0 && m.ancho != 0{
		return true
	}
	return false
}