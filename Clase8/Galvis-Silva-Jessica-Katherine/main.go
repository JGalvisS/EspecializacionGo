package main

import (
	"Galvis-Silva-Jessica-Katherine/internal/tickets"
	"fmt"
)

/*Ejecutar al menos una vez cada requerimiento en la función main. Las ejecuciones deben
realizarse de manera concurrente (utilizando diferentes goroutines).*/

//Estructuras que almacenaran las respuestas
type Result struct {
	Value 	int
	Err  	error
}

type PercentResult struct {
	Percent 	float64
	Err   	error
}

func main() {
	destination := "China"
	period := "noche"

	//Canales
	PaxPerDestinationCh := make(chan Result)
	PaxPerPeriodCh := make(chan Result)
	PercentPerDestinationCh := make(chan PercentResult)

	//Gorutines
	go func ()  {
		value, err := tickets.GetTotalTickets(destination)
		PaxPerDestinationCh <- Result{
			Value: value,
			Err: err,
		}
	}()

	go func ()  {
		value, err := tickets.GetCountByPeriod(period)
		PaxPerPeriodCh <-  Result{
			Value: value,
			Err: err,
		}
	}()

	go func() {
		percent, err := tickets.AverageDestination(destination)
		PercentPerDestinationCh <- PercentResult{
			Percent: percent,
			Err: err,
		}
	}()

	//Imprimir resultados
	totalPaxDestination := <-PaxPerDestinationCh
	if totalPaxDestination.Err != nil {
		fmt.Println("Error al obtener los pasajeros por destino debido a:", totalPaxDestination.Err)
	}else{fmt.Println("A ",destination," viajaron en total ",totalPaxDestination.Value," pasajeros")}

	totalPaxPerPeriod := <-PaxPerPeriodCh
	if totalPaxDestination.Err != nil {
		fmt.Println("Error al obtener los pasajeros por periodo:", totalPaxPerPeriod.Err)
	}else{fmt.Println("En el periodo ", period," viajaron un total de ", totalPaxPerPeriod.Value, " pasajeros")}

	percenPerDestination := <-PercentPerDestinationCh
	if percenPerDestination.Err != nil{
		fmt.Println("Error al obtener el porcentaje de pasajeros: ",percenPerDestination.Err)
	}else{fmt.Println("Un ", percenPerDestination.Percent,"porciento de los pasajeros viajaron a ", destination)}
}
