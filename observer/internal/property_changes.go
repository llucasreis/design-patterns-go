package internal

import (
	"container/list"
	"fmt"
)

type Person2 struct {
	Observable
	age int
}

type PropertyChange struct {
	Name  string
	Value any
}

func NewPerson2(age int) *Person2 {
	return &Person2{
		Observable: Observable{new(list.List)},
		age:        age,
	}
}

func (p *Person2) Age() int {
	return p.age
}

func (p *Person2) SetAge(age int) {
	if age == p.age {
		return
	}
	p.age = age
	p.Fire(PropertyChange{"age", age})
}

type TrafficManagement struct {
	o Observable
}

func (t *TrafficManagement) Notify(data any) {
	if pc, ok := data.(PropertyChange); ok && pc.Name == "age" {
		if pc.Value.(int) >= 16 {
			println("Congrats, you can drive now!")
			t.o.Unsubscribe(t)
		}
	}
}

func RunPropertyChanges() {
	p := NewPerson2(15)
	t := &TrafficManagement{p.Observable}
	p.Subscribe(t)

	for i := 16; i <= 20; i++ {
		fmt.Println("Setting age to", i)
		p.SetAge(i)
	}
}
