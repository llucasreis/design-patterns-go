package main

import (
	"fmt"

	"github.com/llucasreis/design-patterns-go/chain_of_responsibility/internal"
)

func main() {
	internal.RunMethodChain()
	fmt.Println("-------------------------------")
	internal.RunCommandQuery()
}
