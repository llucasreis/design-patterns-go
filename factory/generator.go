package factory

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// NewEmployeeFactory high-order function to create a user
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{
			Name:         name,
			Position:     position,
			AnnualIncome: annualIncome,
		}
	}
}

type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (ef *EmployeeFactory) Create(name string) *Employee {
	return &Employee{
		Name:         name,
		Position:     ef.Position,
		AnnualIncome: ef.AnnualIncome,
	}
}

// NewEmployeeFactory2 This method allows to change fields of the struct
func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{
		Position:     position,
		AnnualIncome: annualIncome,
	}
}

func RunGenerator() {
	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 8000)

	developer := developerFactory("John")
	manager := managerFactory("Jane")

	fmt.Printf("Developer: %+v\n", developer)
	fmt.Printf("Manager: %+v\n", manager)

	bossFactory := NewEmployeeFactory2("Boss", 100000)
	boss := bossFactory.Create("James")

	fmt.Printf("Boss: %+v\n", boss)

	fmt.Println()
}
