package tickets

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Write code here to run before tests
	GetInfo("/Users/scastelli/GolandProjects/desafio-go-bases/tickets.csv")
	// Run tests
	exitVal := m.Run()
	// Write code here to run after tests
	// Exit with exit value from tests
	os.Exit(exitVal)
}

func TestGetTotalTickets(t *testing.T) {
	t.Run("Cantidad de viajes a finlandia", func(t *testing.T) {
		//Arrange
		destination := "Finland"
		//Act
		result, _ := GetTotalTickets(destination)
		//Assert
		correctResult := 8
		if result != correctResult {
			t.Fatal("Viajaron "+string(rune(correctResult))+" personas y el resultado fue: ", result)
		}
	})
	t.Run("Cantidad de viajes a pais inexistente", func(t *testing.T) {
		//Arrange
		destination := "AA"
		//Act
		_, err := GetTotalTickets(destination)
		//Assert
		if err == nil {
			t.Fatal("No existe el destino " + destination)
		}
	})
}

func TestAverageDestination(t *testing.T) {
	t.Run("Promedio de viajes a Finlandia", func(t *testing.T) {
		//Arrange
		destination := "Finland"
		//Act
		result, _ := AverageDestination(destination)
		//Assert
		correctResult := 0.008
		if result != correctResult {
			t.Fatal("Viajaron un promedio de "+fmt.Sprintf("%v", correctResult)+" personas al destino "+destination+" y el resultado fue: ", result)
		}
	})
	t.Run("Promedio de viajes a pais inexistente", func(t *testing.T) {
		//Arrange
		destination := "AA"
		//Act
		_, err := AverageDestination(destination)
		//Assert
		if err == nil {
			t.Fatal("No existe el pais " + destination)
		}
	})
}

func TestGetCountByPeriod(t *testing.T) {
	t.Run("Cantidad de viajes nocturnos", func(t *testing.T) {
		//Arrange
		period := "Noche"
		//Act
		result, _ := GetCountByPeriod(period)
		//Assert
		correctResult := 151
		if result != correctResult {
			t.Fatal("Viajaron "+string(rune(correctResult))+" personas en el período"+period+"y el resultado fue: ", result)
		}
	})
	t.Run("Cantidad de viajes a periodo inexistente", func(t *testing.T) {
		//Arrange
		period := "AA"
		//Act
		_, err := GetCountByPeriod(period)
		//Assert
		if err == nil {
			t.Fatal("No existe el periodo " + period)
		}
	})
}
