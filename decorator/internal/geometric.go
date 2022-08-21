package internal

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle with radius %f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %f", s.Side)
}

type ColoredShape struct {
	Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s with color %s", c.Shape.Render(), c.Color)
}

type TransparentShape struct {
	Shape
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s with transparency %f%%", t.Shape.Render(), (t.Transparency * 100.0))
}

func RunGeometric() {
	circle := Circle{2}
	circle.Resize(2)
	fmt.Println(circle.Render())

	redCircle := ColoredShape{&circle, "red"}
	fmt.Println(redCircle.Render())

	rhsCircle := TransparentShape{&redCircle, 0.5}
	fmt.Println(rhsCircle.Render())
}
