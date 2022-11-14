package main

import (
	"fmt"

	"github.com/llucasreis/design-patterns-go/interpreter/internal"
)

func main() {
	internal.RunLexing()
	fmt.Println("-----")
	internal.RunParsing()
}
