package internal

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address2 struct {
	StreetAddress2, City, Country string
}

type Person2 struct {
	Name     string
	Address2 *Address2
	Friends  []string
}

func (p *Person2) DeepCopy() *Person2 {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := Person2{}
	_ = d.Decode(&result)
	return &result
}

func RunSerialization() {
	john := Person2{Name: "John",
		Address2: &Address2{
			StreetAddress2: "123 London Road",
			City:           "London",
			Country:        "UK",
		},
		Friends: []string{"Chris", "Matt"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address2.StreetAddress2 = "321 Baker Street"

	fmt.Printf("%+v, %+v\n", john, john.Address2)
	fmt.Printf("%+v, %+v\n", jane, jane.Address2)
}
