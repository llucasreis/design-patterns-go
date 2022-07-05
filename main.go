package main

import (
	"fmt"

	"github.com/llucasreis/design-patterns-go/builder"
)

func main() {
	// Builder

	builder.RunBuilderSimple()
	builder.RunBuilderMultiple()
	builder.RunBuilderParameter()
	builder.RunBuilderFunctional()

	fmt.Println("-----------------")

}
