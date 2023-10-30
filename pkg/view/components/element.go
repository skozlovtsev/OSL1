package components

import (
	"fmt"
	"osl1/pkg/view/colors"
)

// Element Interface
type Element interface {
	Content() string
	IsSelectable() bool
}

// Selectable Interface
type SelectableElement interface {
	Op() int
	Select()
	Back()
}

// Raw Text Element
type Text struct {
	content string
	esc     string
}

func NewText(content string, esc string) *Text {
	return &Text{
		content: content,
		esc:     esc,
	}
}

func (t *Text) Content() string {
	return fmt.Sprintf(
		"%s%s%s",
		t.esc,
		t.content,
		colors.ResetColor,
	)
}

func (t *Text) IsSelectable() bool {
	return false
}

// Option Element
type Option struct {
	op         int
	content    string
	esc        string
	selectFunc func()
	back       func()
}

func NewOption(op int, content string, esc string, selectFunc func(), back func()) *Option {
	return &Option{
		op:         op,
		content:    content,
		esc:        esc,
		selectFunc: selectFunc,
		back:       back,
	}
}

func (o *Option) Op() int {
	return o.op
}

func (o *Option) Conent() string {
	return fmt.Sprintf(
		"%b) %s%s%s",
		o.op,
		o.esc,
		o.content,
		colors.ResetColor,
	)
}

func (o *Option) IsSelectable() bool {
	return true
}

func (o *Option) Select() {
	o.selectFunc()
}

func (o *Option) Back() {
	o.back()
}

// Input Element
type Input struct {
}
