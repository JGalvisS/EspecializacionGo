package main
import (
	"errors"
	"fmt"
)

/*Crear un programa que cumpla los siguientes puntos:
Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
Tener un slice global de Product llamado Products, instanciado con valores.
Dos métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método. El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
Una función getById() a la cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado. 
Ejecutar al menos una vez cada método y función definidos desde main().*/

type Product struct{
	Id 			int
	Name 		string
	Price 		float64
	Description string
	Category 	string
}

var Products = []Product{
	{Id: 1, Name: "Laptop", Price: 1200.50, Description: "Laptop de 15 pulgadas", Category: "Electrónica"},
	{Id: 2, Name: "Mouse", Price: 25.00, Description: "Mouse inalámbrico", Category: "Accesorios"},
}

func main()  {
	newProduct := Product{Id: 3, Name: "Teclado", Price: 45.00, Description: "Teclado mecánico", Category: "Accesorios"}
	fmt.Printf("Nuevo producto: ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s \n",newProduct.Id, newProduct.Name, newProduct.Price, newProduct.Description, newProduct.Category)
	fmt.Println()
	fmt.Println("Guardando nuevo producto en la lista...")
	fmt.Println()
	err := newProduct.Save()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("El producto ", newProduct.Name," fue guardado con éxito.")
    }
	fmt.Println()
	fmt.Println("LISTA DE PRODUCTOS")
	newProduct.GetAll()
	fmt.Println()

	idPorBuscar := 1
	fmt.Printf("Buscando producto con ID: %d\n", idPorBuscar)
	product, err := getById(idPorBuscar)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Producto encontrado: ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n",
			product.Id, product.Name, product.Price, product.Description, product.Category)
	}
}

func (p Product)Save() error{
	originLen := len(Products)
	Products = append(Products,p)
	if len(Products)== originLen+1 {
	return nil
	}
	return errors.New("El producto NO fue guardado con éxito")
}

func (p Product) GetAll() {
		for _, product := range Products {
        fmt.Printf("ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n",
            product.Id, product.Name, product.Price, product.Description, product.Category)
    }	
}

func getById(id int)(Product, error){
	for _, product := range Products{
		if product.Id == id{
			return product, nil
		}
	}
	return Product{},errors.New("El producto no fue encontrado")
}


