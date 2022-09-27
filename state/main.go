package main

import (
	"fmt"

	"github.com/llucasreis/design-patterns-go/state/internal"
)

func main() {
	internal.RunClassic()
	fmt.Println("------------------------")
	internal.RunHandmadeStateMachine()
	fmt.Println("------------------------")
	internal.RunSwitchStateMachine()
}
