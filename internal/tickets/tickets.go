package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	Id      int
	Name    string
	Mail    string
	Country string
	Time    string
	Price   int
}

var allTickets []Ticket

// ejemplo 1
func GetTotalTickets(destination string) (result int, err error) {
	for _, ticket := range allTickets {
		if ticket.Country == destination {
			result++
		}
	}
	if result == 0 {
		err = errors.New("No existe el destino o no hay nadie que haya viajado alli\n")
	}
	return
}

func GetCountByPeriod(time string) (result int, err error) {
	for _, ticket := range allTickets {
		if ticket.Time == time {
			result++
		}
	}
	if result == 0 {
		err = errors.New("No hay nadie que viaje en ese horario o está mal escrito el periodo")
	}
	return
}

// ejemplo 3
func AverageDestination(destination string) (result float64, err error) {
	totalTicketsForDestination, err := GetTotalTickets(destination)
	if err != nil {
		fmt.Println(err)
	}
	result = float64(totalTicketsForDestination) / float64(len(allTickets))
	if result == 0.0 {
		err = errors.New("No fue posible calcular el promedio de personas que viajan al destino " + destination)
	}
	return
}

func GetInfo(path string) []Ticket {
	fd, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//fmt.Println("Archivo abierto exitosamente")
	defer func(fd *os.File) {
		err := fd.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(fd)
	//leer el file con csv
	fileReader := csv.NewReader(fd)
	info, err2 := fileReader.ReadAll()
	if err2 != nil {
		fmt.Println(err2)
		return nil
	}
	for _, row := range info {
		id, err := strconv.Atoi(row[0])
		if err != nil {
			fmt.Println(err)
		}
		price, err := strconv.Atoi(row[0])
		if err != nil {
			fmt.Println(err)
		}

		periodName, err := getPeriodName(row[4])
		if err != nil {
			fmt.Println(err)
		}
		ticket := Ticket{
			Id:      id,
			Name:    row[1],
			Mail:    row[2],
			Country: row[3],
			Time:    periodName,
			Price:   price,
		}
		allTickets = append(allTickets, ticket)
	}
	return allTickets
}

func getPeriodName(time string) (period string, err error) {
	hour := strings.Split(time, ":")[0]
	hourInt, err := strconv.Atoi(hour)
	if err != nil {
		println(err)
	}
	if hourInt >= 0 && hourInt < 7 {
		period = "Madrugada"
		return period, nil
	}
	if hourInt >= 7 && hourInt < 13 {
		period = "Mañana"
		return period, nil

	}
	if hourInt >= 13 && hourInt < 20 {
		period = "Tarde"
		return period, nil
	}
	if hourInt >= 20 && hourInt < 24 {
		period = "Noche"
		return period, nil
	}

	err = errors.New("La hora que se encuentra en los registros está mal formateada")
	return period, err

}
