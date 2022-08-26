package main

import (
	"fmt"

	"github.com/llucasreis/design-patterns-go/template_method/internal"
)

func main() {
	internal.RunTemplateMethod()
	fmt.Println("------------------------------")
	internal.RunFunctional()
}
