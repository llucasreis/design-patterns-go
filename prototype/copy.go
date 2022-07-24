package prototype

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		StreetAddress: a.StreetAddress,
		City:          a.City,
		Country:       a.Country,
	}
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func RunCopy() {
	john := Person{Name: "John",
		Address: &Address{
			StreetAddress: "123 London Road",
			City:          "London",
			Country:       "UK",
		},
		Friends: []string{"Chris", "Matt"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker Street"

	fmt.Printf("%+v, %+v\n", john, john.Address)
	fmt.Printf("%+v, %+v\n", jane, jane.Address)
}
