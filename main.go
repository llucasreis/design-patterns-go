package main

import (
	"fmt"

	"github.com/llucasreis/design-patterns-go/builder"
	"github.com/llucasreis/design-patterns-go/factory"
)

func main() {
	// Builder

	builder.RunBuilderSimple()
	builder.RunBuilderMultiple()
	builder.RunBuilderParameter()
	builder.RunBuilderFunctional()

	fmt.Println("-----------------")

	// Factory

	factory.RunFunction()
	factory.RunInterface()
	factory.RunGenerator()

}
