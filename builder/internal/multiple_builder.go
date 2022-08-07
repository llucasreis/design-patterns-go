package internal

import "fmt"

type Employee struct {
	// address
	StreetAddress, Postcode, City string

	// job
	CompanyName, Position string
	AnnualIncome          int
}

type EmployeeBuilder struct {
	employee *Employee
}

func (eb *EmployeeBuilder) Lives() *EmployeeAddressBuilder {
	return &EmployeeAddressBuilder{*eb}
}

func (eb *EmployeeBuilder) Works() *EmployeeJobBuilder {
	return &EmployeeJobBuilder{*eb}
}

func NewEmployeeBuilder() *EmployeeBuilder {
	return &EmployeeBuilder{
		employee: &Employee{},
	}
}

type EmployeeAddressBuilder struct {
	EmployeeBuilder
}

type EmployeeJobBuilder struct {
	EmployeeBuilder
}

func (eab *EmployeeAddressBuilder) At(streetAddress string) *EmployeeAddressBuilder {
	eab.employee.StreetAddress = streetAddress

	return eab
}

func (eab *EmployeeAddressBuilder) In(city string) *EmployeeAddressBuilder {
	eab.employee.City = city

	return eab
}

func (eab *EmployeeAddressBuilder) WithPostcode(postcode string) *EmployeeAddressBuilder {
	eab.employee.Postcode = postcode

	return eab
}

func (ejb *EmployeeJobBuilder) At(companyName string) *EmployeeJobBuilder {
	ejb.employee.CompanyName = companyName

	return ejb
}

func (ejb *EmployeeJobBuilder) AsA(position string) *EmployeeJobBuilder {
	ejb.employee.Position = position

	return ejb
}

func (ejb *EmployeeJobBuilder) Earning(annualIncome int) *EmployeeJobBuilder {
	ejb.employee.AnnualIncome = annualIncome

	return ejb
}

func (eb *EmployeeBuilder) Build() *Employee {
	return eb.employee
}

func RunBuilderMultiple() {
	eb := NewEmployeeBuilder()
	eb.
		Lives().
		At("123 London Road").
		In("London").
		WithPostcode("SW12BC").
		Works().
		At("Fabrikam").
		AsA("Engineer").
		Earning(120000)

	employee := eb.Build()

	fmt.Printf("Employee: %+v\n", employee)

	fmt.Println()
}

// func main() {
// 	pb := NewPersonBuilder()
// 	pb.
// 		Lives().
// 		At("123 London Road").
// 		In("London").
// 		WithPostcode("SW12BC").
// 		Works().
// 		At("Fabrikam").
// 		AsA("Engineer").
// 		Earning(120000)

// 	person := pb.Build()

// 	fmt.Printf("Person: %+v\n", person)
// }
