package internal

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpression(e *DoubleExpression3)
	VisitAdditionExpression(e *AdditionExpression3)
}

type Expression3 interface {
	Accept(ev ExpressionVisitor)
}

type DoubleExpression3 struct {
	value float64
}

func (d *DoubleExpression3) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(d)
}

type AdditionExpression3 struct {
	left, right Expression3
}

func (a *AdditionExpression3) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(a)
}

type ExpressionPrinter struct {
	sb strings.Builder
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{strings.Builder{}}
}

func (ep *ExpressionPrinter) VisitDoubleExpression(e *DoubleExpression3) {
	ep.sb.WriteString(fmt.Sprintf("%g", e.value))
}

func (ep *ExpressionPrinter) VisitAdditionExpression(e *AdditionExpression3) {
	ep.sb.WriteRune('(')
	e.left.Accept(ep)
	ep.sb.WriteRune('+')
	e.right.Accept(ep)
	ep.sb.WriteRune(')')
}

func (ep *ExpressionPrinter) String() string {
	return ep.sb.String()
}

type ExpressionEvaluator struct {
	result float64
}

func (ee *ExpressionEvaluator) VisitDoubleExpression(e *DoubleExpression3) {
	ee.result = e.value
}

func (ee *ExpressionEvaluator) VisitAdditionExpression(e *AdditionExpression3) {
	e.left.Accept(ee)
	left := ee.result
	e.right.Accept(ee)
	right := ee.result
	ee.result = left + right
}

func RunClassicDispatch() {
	// 1 + (2 + 3) = 6
	e := &AdditionExpression3{
		left: &DoubleExpression3{1},
		right: &AdditionExpression3{
			left:  &DoubleExpression3{2},
			right: &DoubleExpression3{3},
		},
	}

	ep := NewExpressionPrinter()
	e.Accept(ep)

	ee := &ExpressionEvaluator{}
	e.Accept(ee)

	fmt.Printf("%s = %g\n", ep.String(), ee.result)
}
