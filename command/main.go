package main

import (
	"fmt"

	"github.com/llucasreis/design-patterns-go/command/internal"
)

func main() {
	internal.RunCommand()
	fmt.Println("----")
	internal.RunCompositeCommands()
	fmt.Println("----")
	internal.RunFunctional()
}
