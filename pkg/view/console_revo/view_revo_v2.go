package consolerevo

import (
	"osl1/pkg/view/components"
)

var (
	Main *components.View
	File *components.View
	Json *components.View
	XML  *components.View
	Zip  *components.View
)

func init() {
	Main = components.NewView(
		[]components.Container{},
	)
	File = components.NewView(
		[]components.Container{},
	)
	Json = components.NewView(
		[]components.Container{},
	)
	XML = components.NewView(
		[]components.Container{},
	)
	Zip = components.NewView(
		[]components.Container{},
	)
}
