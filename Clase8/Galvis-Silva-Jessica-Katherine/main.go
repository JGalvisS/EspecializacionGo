package main

import (
	"Galvis-Silva-Jessica-Katherine/internal/tickets"
	"fmt"
)

func main() {
	total, err := tickets.GetTotalTickets("Brazil")
	if err != nil{
		fmt.Println(err)
	}else{fmt.Println(total)}

	paxPorHora, err := tickets.GetCountByPeriod("madrugada")
	if err != nil{
		fmt.Println(err)
	}else{fmt.Println(paxPorHora)}
}
