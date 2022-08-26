package internal

import (
	"fmt"
	"strings"
)

type OutputFormat int

const (
	Markdown OutputFormat = iota
	HTML
)

type ListStrategy interface {
	Start(builder *strings.Builder)
	End(builder *strings.Builder)
	AddListItem(builder *strings.Builder, item string)
}

type MarkdownListStrategy struct{}

func (m *MarkdownListStrategy) Start(builder *strings.Builder) {}

func (m *MarkdownListStrategy) End(builder *strings.Builder) {}

func (m *MarkdownListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString(" * " + item + "\n")
}

type HTMLListStrategy struct{}

func (h *HTMLListStrategy) Start(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}

func (h *HTMLListStrategy) End(builder *strings.Builder) {
	builder.WriteString("</ul>\n")
}

func (h *HTMLListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString("<li>" + item + "</li>\n")
}

type TextProcessor struct {
	builder      strings.Builder
	listStrategy ListStrategy
}

func NewTextProcessor(listStrategy ListStrategy) *TextProcessor {
	return &TextProcessor{strings.Builder{}, listStrategy}
}

func (t *TextProcessor) SetOutputFormat(fmt OutputFormat) {
	switch fmt {
	case Markdown:
		t.listStrategy = &MarkdownListStrategy{}
	case HTML:
		t.listStrategy = &HTMLListStrategy{}
	}
}

func (t *TextProcessor) AppendList(items []string) {
	s := t.listStrategy
	s.Start(&t.builder)
	for _, item := range items {
		s.AddListItem(&t.builder, item)
	}
	s.End(&t.builder)
}

func (t *TextProcessor) Reset() {
	t.builder.Reset()
}

func (t *TextProcessor) String() string {
	return t.builder.String()
}

func RunStrategy() {
	tp := NewTextProcessor(&MarkdownListStrategy{})
	tp.AppendList([]string{"Item 1", "Item 2", "Item 3"})
	fmt.Println(tp.String())

	tp.Reset()
	tp.SetOutputFormat(HTML)
	tp.AppendList([]string{"Item 1", "Item 2", "Item 3"})
	fmt.Println(tp.String())
}
