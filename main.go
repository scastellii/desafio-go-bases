package main

import (
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	// total, err := tickets.GetTotalTickets("Brazil")
	tickets.GetInfo("tickets.csv")
	destination := "Finland"
	period := "Noche"

	//Requerimiento 1
	totalTickets, err := tickets.GetTotalTickets(destination)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Viajaron a %s: %d personas\n", destination, totalTickets)
	}

	//Requermiento 2
	result, err := tickets.GetCountByPeriod(period)
	if err != nil {
		fmt.Println(err)
	} else {
		text := fmt.Sprintf("Viajaron en el período %s: %d personas", period, result)
		fmt.Println(text)
	}

	//Requereimiento 3
	average, err := tickets.AverageDestination(destination)
	if err != nil {
		fmt.Println(err)
	} else {
		text := fmt.Sprintf("Viajaron al destino %s: un promedio de %g personas", destination, average)
		fmt.Println(text)
	}
}
