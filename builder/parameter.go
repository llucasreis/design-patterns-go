package builder

import (
	"fmt"
	"strings"
)

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email
}

func (eb *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("invalid email address")
	}
	eb.email.from = from
	return eb
}

func (eb *EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		panic("invalid email address")
	}
	eb.email.to = to
	return eb
}

func (eb *EmailBuilder) Subject(subject string) *EmailBuilder {
	eb.email.subject = subject
	return eb
}

func (eb *EmailBuilder) Body(body string) *EmailBuilder {
	eb.email.body = body
	return eb
}

type build func(*EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendEmailImpl(&builder.email)
}

func sendEmailImpl(email *email) {
	fmt.Printf("Sending email %+v\n", email)
}

func RunBuilderParameter() {
	SendEmail(func(eb *EmailBuilder) {
		eb.
			From("foo@bar.com").
			To("bar@barz.com").
			Subject("Hello").
			Body("Hello, world!")
	})

	fmt.Println()
}

// func main() {
// 	SendEmail(func(eb *EmailBuilder) {
// 		eb.
// 			From("foo@bar.com").
// 			To("bar@barz.com").
// 			Subject("Hello").
// 			Body("Hello, world!")
// 	})
// }
