package internal

import "fmt"

type User interface {
	SayHello()
}

type developer struct {
	name   string
	salary int
}

type manager struct {
	name   string
	salary int
}

func (d *developer) SayHello() {
	fmt.Printf("Hello, my name is %s and I'm a developer \n", d.name)
}

func (m *manager) SayHello() {
	fmt.Printf("Hello, my name is %s and I'm a manager \n", m.name)
}

func NewUser(name string, salary int) User {
	if salary >= 10000 {
		return &manager{name: name, salary: salary}
	}
	return &developer{name: name, salary: salary}
}

// not exposing the struct, just the interface
func RunInterface() {
	u := NewUser("John", 5000)
	u.SayHello()

	u2 := NewUser("Jane", 10000)
	u2.SayHello()

	fmt.Println()
}
