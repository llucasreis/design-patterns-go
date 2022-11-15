package main

import (
	"fmt"

	"github.com/llucasreis/design-patterns-go/observer/internal"
)

func main() {
	internal.RunObserver()
	fmt.Println("----")
	internal.RunPropertyChanges()
	fmt.Println("----")
	internal.RunPropertyDependencies()
}
