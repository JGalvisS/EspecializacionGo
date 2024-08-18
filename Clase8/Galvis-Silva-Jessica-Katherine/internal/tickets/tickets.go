package tickets

import (
	"errors"
	"os"
	"strings"
)

/*Una aerolínea pequeña tiene un sistema de reservas de pasajes a diferentes países. Este
retorna un archivo con la información de los pasajes sacados en las últimas 24 horas. La
aerolínea necesita un programa para extraer información de las ventas del día y, así,
analizar las tendencias de compra.
El archivo en cuestión es del tipo valores separados por coma (CSV), donde los campos
están compuestos por: id, nombre, email, país de destino, hora del vuelo y precio.*/

type Ticket struct {
	Id		string
	Nombre 	string
	Email	string
	Destino	string
	Hora	string
	precio	int
}

// Funcion recibe un string país de destino y calcula cuantas personas viajan a ese país
func GetTotalTickets(destination string) (int, error) {
	var amountPax int
	res, err := os.ReadFile("./tickets.csv")
	if err != nil{
		amountPax = 0
		return amountPax, errors.New("La base de datos no pudo ser leída")
	}
	data := strings.Split(string(res),"\n")
	for i := 0; i < len(data); i++ {
		line:= strings.Split(string(data[i]), ",")
		if len(line)>= 4 && line[3] == destination {
			amountPax ++
		}
	}
	return amountPax, nil
}

//funcion que recibe un periodo de tiempo y calcula cuántas personas viajan en ese tiempo:
//madrugada (0 → 6), mañana (7 → 12), tarde (13 → 19), y noche (20 → 23).
func GetCountByPeriod(time string) (int, error) {
	var amountPax int
	period:=[]string{}
switch time {
	case "madrugada":
		period = append(period,"0","1","2","3","4","5","6")
	case "mañana":
		period = append(period,"7","8","9","10","11","12")
	case "tarde":
		period = append(period, "13","14","15","16","17","18","19")
	case "noche":
		period = append(period, "20","21","22","23","24")
	default:
		return 0,errors.New("Este periodo de tiempo no es computable.")
}
	res, err := os.ReadFile("./tickets.csv")
	if err != nil{
		amountPax = 0
		return amountPax, errors.New("La base de datos no pudo ser leída")
	}
	data := strings.Split(string(res),"\n")
	for i := 0; i < len(data); i++ {
		line:= strings.Split(string(data[i]), ",")
		if len(line) >=5{
			sectionTime:=line[4]
			hour:=strings.Split(string(sectionTime),":")
			for _,hora := range period{
				if hora ==hour[0] {
					amountPax ++
				}
			}
			
		}
		
	}
	return amountPax,nil
}

//funcion que recibe un destino string y calcula el porcentaje de personas que viajan a ese destino.
/*NOTA: Como el reporte dice que se genera cada 24 horas no le veo sentido al parametro total int*/
func AverageDestination(destination string) (float64, error) {
	var Paxpercent float64
	res, err := os.ReadFile("./tickets.csv")
	if err != nil{
		Paxpercent = 0
		return Paxpercent, errors.New("La base de datos no pudo ser leída")
	}
	data := strings.Split(string(res),"\n")
	amountPaxPerDay := len(data)
	amountPaxPerDestination,errTT := GetTotalTickets(destination)
	if errTT != nil{
		Paxpercent = 0
		return Paxpercent, errors.New("No se logro obtener la cantidad de pasajeros que viajaron")
	}
	Paxpercent = float64(float64(amountPaxPerDestination) / float64(amountPaxPerDay) * 100)
	return Paxpercent,nil
}
