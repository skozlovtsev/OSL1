package components

import (
	"fmt"

	"osl1/pkg/view/colors"
)

type View struct {
	conent []Container
}

func NewView(content []Container) *View {
	return &View{
		conent: content,
	}
}

func (v *View) Render() error {
	fmt.Print(colors.ClearScreen)
	for _, c := range v.conent {
		c.Render()
	}
	return nil
}

func (v *View) ListenForAction() int {
	for {
		var action int
		fmt.Scanf("%d", &action)

		switch {
		case action > 0 && action < 1:
			return action
		case action < 0:
			fmt.Print(colors.ClearLine)
			continue
		default:
			fmt.Print(colors.ClearLine)
			continue
		}
	}
}
