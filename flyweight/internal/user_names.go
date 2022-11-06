package internal

import (
	"fmt"
	"strings"
)

type User struct {
	FullName string
}

func NewUser(fullName string) *User {
	return &User{FullName: fullName}
}

var allNames []string

type User2 struct {
	names []uint8
}

func NewUser2(fullName string) *User2 {
	getOrAdd := func(s string) uint8 {
		for i, name := range allNames {
			if name == s {
				return uint8(i)
			}
		}
		allNames = append(allNames, s)
		return uint8(len(allNames) - 1)
	}

	result := User2{}
	parts := strings.Split(fullName, " ")
	for _, p := range parts {
		result.names = append(result.names, getOrAdd(p))
	}
	return &result
}

func (u *User2) FullName() string {
	var parts []string
	for _, i := range u.names {
		parts = append(parts, allNames[i])
	}
	return strings.Join(parts, " ")
}

func RunUserNames() {
	john := NewUser("John Doe")

	fmt.Println(john.FullName)

	fmt.Println("----")

	john2 := NewUser2("John Doe")
	fmt.Println(john2.FullName())
}
