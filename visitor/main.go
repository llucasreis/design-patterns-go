package main

import (
	"fmt"

	"github.com/llucasreis/design-patterns-go/visitor/internal"
)

func main() {
	internal.RunIntrusive()
	fmt.Println("----")
	internal.RunReflective()
	fmt.Println("----")
	internal.RunClassicDispatch()
}
