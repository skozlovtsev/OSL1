package components

// Container Interface
type Container interface {
	Select(int)
	Render()
}

type ElementSelectable interface {
	Element
	SelectableElement
}

// Horizontal Container
type Horizontal struct {
	Conent []ElementSelectable
}

func NewHorizontal() *Horizontal {
	return &Horizontal{}
}

func (h *Horizontal) Select(op int) {
	for _, c := range h.Conent {
		if c.IsSelectable() {
			if c.Op() == op {
				c.Select()
			}
		}
	} // for
}

func (h *Horizontal) Render() {}

// Vertical Container
type Vertical struct {
	Conent []ElementSelectable
}

func NewVertical() *Vertical {
	return &Vertical{}
}

func (v *Vertical) Select(op int) {}

func (v *Vertical) Render() {}
