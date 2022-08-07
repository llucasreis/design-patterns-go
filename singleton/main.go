package main

import (
	"fmt"

	"github.com/llucasreis/design-patterns-go/singleton/internal"
)

func main() {
	db := internal.GetSingletonDatabase()

	population := db.GetPopulation("Seoul")

	fmt.Println("Real Population: ", population)

	db2 := internal.GetDummyDatabase()

	population2 := db2.GetPopulation("Seoul")

	fmt.Println("Dummy Population: ", population2)

	totalPopulation := internal.GetTotalPopulation(db, []string{"Seoul", "Tokyo"})

	fmt.Println("Real Total population: ", totalPopulation)

	totalPopulation2 := internal.GetTotalPopulation(db2, []string{"Seoul", "Tokyo"})

	fmt.Println("Dummy Total population: ", totalPopulation2)
}
