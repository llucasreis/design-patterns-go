package internal

import "fmt"

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name:     name,
		Age:      age,
		EyeCount: 2,
	}
}

func RunFunction() {
	p := NewPerson("John", 30)
	fmt.Printf("Person: %+v\n", p)

	fmt.Println()
}
