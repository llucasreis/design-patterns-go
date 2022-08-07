package internal

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address3 struct {
	Suite                int
	StreetAddress3, City string
}

type Employee struct {
	Name   string
	Office Address3
}

func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := Employee{}
	_ = d.Decode(&result)
	return &result
}

var mainOffice = Employee{
	"",
	Address3{0, "123 East Dr", "London"}}

var auxOffice = Employee{
	"",
	Address3{0, "66 West Dr", "London"}}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func RunFactory() {
	john := NewMainOfficeEmployee("John", 123)
	jane := NewAuxOfficeEmployee("Jane", 66)

	fmt.Printf("%+v, %+v\n", john, john.Office)
	fmt.Printf("%+v, %+v\n", jane, jane.Office)
}
