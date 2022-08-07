package main

import (
	"fmt"

	"github.com/llucasreis/design-patterns-go/singleton/internal"
)

func main() {
	db := internal.GetSingletonDatabase()

	population := db.GetPopulation("Seoul")

	fmt.Println(population)
}
