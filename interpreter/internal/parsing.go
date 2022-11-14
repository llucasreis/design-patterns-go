package internal

import (
	"fmt"
	"strconv"
)

type Element interface {
	Value() int
}

type Integer struct {
	value int
}

func NewInteger(value int) *Integer {
	return &Integer{value: value}
}

func (i *Integer) Value() int {
	return i.value
}

type Operation int

const (
	Addition Operation = iota
	Subtraction
)

type BinaryOperation struct {
	Type        Operation
	Left, Right Element
}

func (b *BinaryOperation) Value() int {
	switch b.Type {
	case Addition:
		return b.Left.Value() + b.Right.Value()
	case Subtraction:
		return b.Left.Value() - b.Right.Value()
	default:
		panic("Unsupported operation")
	}
}

func Parse(tokens []Token) Element {
	result := BinaryOperation{}
	haveLHS := false
	for i := 0; i < len(tokens); i++ {
		t := &tokens[i]
		switch t.Type {
		case Int:
			n, _ := strconv.Atoi(t.Text)
			integer := Integer{n}
			if !haveLHS {
				result.Left = &integer
				haveLHS = true
			} else {
				result.Right = &integer
			}
		case Plus:
			result.Type = Addition
		case Minus:
			result.Type = Subtraction
		case Lparen:
			j := i
			for ; j < len(tokens); j++ {
				if tokens[j].Type == Rparen {
					break
				}
			}
			var subexp []Token
			for k := i + 1; k < j; k++ {
				subexp = append(subexp, tokens[k])
			}

			element := Parse(subexp)
			if !haveLHS {
				result.Left = element
				haveLHS = true
			} else {
				result.Right = element
			}
			i = j
		}
	}
	return &result
}

func RunParsing() {
	input := "(13+4)-(12+1)"
	tokens := Lex(input)
	result := Parse(tokens)
	fmt.Printf("%s = %d\n", input, result.Value())
}
