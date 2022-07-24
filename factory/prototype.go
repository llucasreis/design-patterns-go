package factory

import "fmt"

type UserEmployee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

func NewUserEmployee(role int) *UserEmployee {
	switch role {
	case Developer:
		return &UserEmployee{
			Position:     "Developer",
			AnnualIncome: 60000,
		}
	case Manager:
		return &UserEmployee{
			Position:     "Manager",
			AnnualIncome: 8000,
		}
	default:
		return nil
	}
}

func NewPrototype() {
	developer := NewUserEmployee(Developer)
	manager := NewUserEmployee(Manager)

	developer.Name = "John"
	manager.Name = "Jane"

	fmt.Printf("Developer: %+v\n", developer)
	fmt.Printf("Manager: %+v\n", manager)
}
