package internal

import "fmt"

type Person struct {
	name, position string
}

type personMod func(*Person)
type PersonBuilder struct {
	actions []personMod
}

func (pb *PersonBuilder) Called(name string) *PersonBuilder {
	pb.actions = append(pb.actions, func(p *Person) {
		p.name = name
	})
	return pb
}

func (pb *PersonBuilder) WorkAsA(position string) *PersonBuilder {
	pb.actions = append(pb.actions, func(p *Person) {
		p.position = position
	})
	return pb
}

func (pb *PersonBuilder) Build() *Person {
	p := &Person{}
	for _, action := range pb.actions {
		action(p)
	}
	return p
}

func RunBuilderFunctional() {
	pb := PersonBuilder{}
	p := pb.
		Called("John").
		WorkAsA("Software Engineer").
		Build()

	fmt.Printf("Person: %+v\n", p)

	fmt.Println()
}

// func main() {
// 	pb := PersonBuilder{}
// 	p := pb.
// 		Called("John").
// 		WorkAsA("Software Engineer").
// 		Build()

// 	fmt.Printf("%+v\n", p)
// }
