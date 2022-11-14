package main

import (
	"fmt"

	"github.com/llucasreis/design-patterns-go/memento/internal"
)

func main() {
	internal.RunMemento()
	fmt.Println("----")
	internal.RunUndoRedo()
}
