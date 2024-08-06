package main

import (
	"errors"
	"fmt"
)

/*Algunas tiendas de e-commerce necesitan realizar una funcionalidad en Go
para administrar productos y retornar el valor del precio total. La empresa tiene tres tipos de productos:
Pequeño, Mediano y Grande. Pero se espera que sean muchos más.
Los costos adicionales para cada uno son:
Pequeño: solo tiene el costo del producto.
Mediano: el precio del producto + un 3% por mantenerlo en la tienda.
Grande: el precio del producto + un 6% por mantenerlo en la tienda y un adicional de $2500 de costo de envío.
El porcentaje de mantenerlo en la tienda se calcula sobre el precio del producto.
Se requiere una función factory que reciba el tipo de producto y el precio,
y retorne una interfaz Producto que tenga el método Precio.
Se debe poder ejecutar el método Precio y que el método devuelva el precio total basándose
en el costo del producto y los adicionales, en caso de que los tenga.*/

//Creo dos tiendas
type calzadoBucaramanga struct{
	referencia	string
	precio		float64
	tipo		string
}

type PanaleraAG struct{
	nombre		string
	precio		float64
	tipo		string
}


type Producto interface{
	Precio() (float64,error)
}

func main(){

	//Almaceno los prodctos en variables
	var calzadoProd Producto
	var panaleraProd Producto
	var err error

	// Creo productos usando la función factory
	calzadoProd, err = ProductoFactory("calzadoBucaramanga","ARF657",1000.00,"mediano")
	if err != nil {
		fmt.Println(err)
	}


	panaleraProd, err = ProductoFactory("PanaleraAG","Coche para Bebe",5000.00,"grande")
	if err != nil {
		fmt.Println(err)
	}


	// Obtengo el precio de los productos usando la interfaz
	precioCalzado, err := calzadoProd.Precio()
	if err != nil {
		fmt.Println("Error al obtener el precio de calzado:", err)
		return
	}
	fmt.Printf("Precio del calzado: %.2f\n", precioCalzado)

	precioPananlera, err := panaleraProd.Precio()
	if err != nil {
		fmt.Println("Error al obtener el precio del panalera:", err)
		return
	}
	fmt.Printf("Precio del panalera: %.2f\n", precioPananlera)
}
//Estas metodo manejan el precio con envio calzadoBucaramanga
func (t calzadoBucaramanga)Precio()(float64,error){
	switch t.tipo{
	case "pequeño":
		return t.precio,nil
	case "mediano":
		porcentaje := (t.precio / 100)* 3
		return t.precio +porcentaje,nil
	case "grande":
		flete := 2500
		porcentaje := (t.precio / 100)* 6
		return t.precio + porcentaje+ float64(flete),nil
	}
	return 0.00, errors.New("El tamaño no se encuentra disponible")
}
//Estas metodo manejan el precio con envio PanaleraAG
func (t PanaleraAG)Precio()(float64,error){
	switch t.tipo {
	case "pequeño":
		return t.precio,nil
	case "mediano":
		porcentaje := (t.precio / 100)* 3
		return t.precio +porcentaje,nil
	case "grande":
		flete := 2500
		porcentaje := (t.precio / 100)* 6
		return t.precio + porcentaje+ float64(flete),nil
	}
	return 0.00, errors.New("El tamaño no se encuentra disponible")
}
//Fabrica de productos
func ProductoFactory(tienda string,ref string, precio float64, tipo string)(Producto, error){
	switch tienda{
	case "calzadoBucaramanga":
		return &calzadoBucaramanga{referencia: ref,precio: precio, tipo: tipo},nil
	case "PanaleraAG":
		return &PanaleraAG{nombre: ref,precio: precio, tipo: tipo},nil
	default:
		return nil, errors.New("Tienda desconocida")
	}
}
