package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*Creando un servidor
Un supermercado  necesita un sistema para gestionar los productos  frescos que  tienen
publicados en su página web. Para poder  hacer esto, necesitan un  servidor  que ejecute una API
que les permita  manipular  los productos  cargados desde distintos clientes. Los campos que
conforman  un producto  son:
Nombre Tipo  de dato
JSON
Tipo  de dato
Go
Descripción  |  Ejemplo
id number int Identi fi cador en conj unto de  datos  | 15
name string string Nombre  caracteri zado | Cheese - St. Andre
quantity number int Canti dad al macenada | 60
code_value string string Códi go al fanuméri co caracterís tico | S73191A
is_published boolean bool El  producto s e encuentra  publ i cado o no |  True
expiration string string Fecha de venci mi ento  | 12/04/2022
price number float64 Preci o del  producto | 50.15

Vamos a levantar un  servidor  utilizando el paquete  Gin en el puerto  8080. Para probar  nuestros
endpoints  haremos  uso de Postman.
● Crear una ruta  /products/search que nos permita  buscar por  parámetro  los productos
cuyo precio  sea mayor  a un  valor priceGt. */

type Product struct{
	Id			int
	Name		string
	Quantity	int
	Code 		string
	IsPublic 	bool
	Expiration 	string
	Price 		float64
}

var productList = []Product{}

func main() {
	LoadProducts("./products.json", &productList)
	router := gin.Default()

router.GET("/products/search",SearchProduct())
router.Run()
	
}

//LoadProducts recibe como parametro la ruta donde debe buscar la data de productos 
//y un puntero del espacio de memoria donde la almacenara para poder usar esta data
func LoadProducts(path string, list *[]Product)  {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal( "Error: No se logro cargar los productos almacenados en este ruta. ", err)
	}
	//Convirto el json a productos que se almacenaran en la lista que le diga
	if err = json.Unmarshal(file, &list) ; err != nil{
		log.Fatal("Error: No se logro cargar los productos almacenados en este ruta. ",err)
	}
	/*NOTA: Esta sintaxis de if solo funciona si la funcion solo devuelve un valor unico*/
}

//SearchProduct recibe por contexto un valor float64 y devuelve todos los produtos cuyo valor sea superior 
func SearchProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		priceParam :=ctx.Query("priceGt")
		priceGt, err := strconv.ParseFloat(priceParam,64)
		if err != nil {
			log.Fatal("Error: El valor de busqueda no es valido",err)
			ctx.JSON(http.StatusBadRequest,gin.H{
				"error": "El valor de busqueda no es valido",
			})
		}
		//Aqui almacenare los productos con valor superior 
		list := []Product{}

		for _, product := range productList{
			if product.Price > priceGt{
				list = append(list,product)
			}
		}
		if len(list) > 0 {
			ctx.JSON(http.StatusAccepted, list)
		}else{
			ctx.JSON(http.StatusAccepted, gin.H{
				"msj":"No encontre productos mayores al precio ingresado",
			})
		}

		
	}
}