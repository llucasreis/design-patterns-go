package internal

import "fmt"

func (p *Person2) CanVote() bool {
	return p.age >= 18
}

type ElectoralRoll struct{}

func (e *ElectoralRoll) Notify(data any) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congrats, you can vote now!")
		}
	}
}

func RunPropertyDependencies() {
	p := NewPerson2(0)
	er := &ElectoralRoll{}
	p.Subscribe(er)

	for i := 10; i < 20; i++ {
		fmt.Println("Setting age to", i)
		p.SetAge(i)
	}
}
