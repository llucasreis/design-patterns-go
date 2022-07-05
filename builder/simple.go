package builder

import "fmt"

type User struct {
	id        int
	nickname  string
	firstName string
	lastName  string
	email     string
}

type UserBuilder struct {
	user *User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		user: &User{},
	}
}

func (ub *UserBuilder) Id(id int) *UserBuilder {
	ub.user.id = id

	return ub
}

func (ub *UserBuilder) Nickname(nickname string) *UserBuilder {
	ub.user.nickname = nickname

	return ub
}

func (ub *UserBuilder) FirstName(firstName string) *UserBuilder {
	ub.user.firstName = firstName

	return ub
}

func (ub *UserBuilder) LastName(lastName string) *UserBuilder {
	ub.user.lastName = lastName

	return ub
}

func (ub *UserBuilder) Email(email string) *UserBuilder {
	ub.user.email = email

	return ub
}

func (ub *UserBuilder) Build() *User {
	return ub.user
}

func RunBuilderSimple() {
	ub := NewUserBuilder()

	ub.
		Id(1).
		Nickname("john_doe").
		FirstName("John").
		LastName("Doe").
		Email("foo@bar.com")

	user := ub.Build()

	fmt.Printf("User: %+v\n", user)

	fmt.Println()
}
