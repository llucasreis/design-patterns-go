package main

import (
	"fmt"

	"github.com/llucasreis/design-patterns-go/prototype"
)

func main() {
	// Builder

	// builder.RunBuilderSimple()
	// builder.RunBuilderMultiple()
	// builder.RunBuilderParameter()
	// builder.RunBuilderFunctional()

	// fmt.Println("-----------------")

	// // Factory

	// factory.RunFunction()
	// factory.RunInterface()
	// factory.RunGenerator()

	fmt.Println("-----------------")

	// Prototype

	prototype.RunCopy()
	prototype.RunSerialization()
	prototype.RunFactory()
}
