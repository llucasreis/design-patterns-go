package main

import (
	"fmt"

	"github.com/llucasreis/design-patterns-go/proxy/internal"
)

func main() {
	internal.RunProtection()
	fmt.Println("----")
	internal.RunVirtual()
}
