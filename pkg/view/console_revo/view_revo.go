package consolerevo

import (
	"fmt"
	"os"
	"osl1/pkg/workers"

	ansi "github.com/bit101/go-ansi"
	tty "github.com/mattn/go-tty"
)

const (
	txtPostfix = ".txt"

	jsonPostfix = ".json"

	xmlPostfix = ".xml"

	zipPostfix = ".zip"
)

var (
	fileWorker *workers.FileWorker

	jsonWorker workers.JSONWorker

	xmlWorker workers.XMLWorker

	zipWorker workers.ZipWorker

	MainMenu *Menu

	fileMenu *Menu

	jsonMenu *Menu

	xmlMenu *Menu

	zipMenu *Menu
)

func init() {
	wd, _ := os.Getwd()

	fileWorker = workers.NewFileWorker(wd)

	jsonWorker = workers.NewJSONWorker(fileWorker)

	xmlWorker = workers.NewXMLWorker(fileWorker)

	zipWorker = workers.NewZipWorker(fileWorker)

	MainMenu = &Menu{Options: []Option{
		{"<--..", func() {}, true},
		{"Show drive Info", func() {}, false},
		{"Operations with files", func() {}, false},
		{"Operations with JSON", func() {}, false},
		{"Operations with XML", func() {}, false},
		{"Operations with ZI", func() {}, false},
	}}

	fileMenu = &Menu{Options: []Option{
		{"<--..", func() {}, true},
		{"Create new file", func() {}, false},
		{"Write to file", func() {}, false},
		{"Read from file", func() {}, false},
		{"Delete file", func() {}, false},
	}}

	jsonMenu = &Menu{Options: []Option{
		{"<--..", func() {}, true},
		{"Create new JSON", func() {}, false},
		{"Write to JSON", func() {}, false},
		{"Read from JSON", func() {}, false},
		{"Delete JSON", func() {}, false},
	}}

	xmlMenu = &Menu{Options: []Option{
		{"<--..", func() {}, true},
		{"Create new XML", func() {}, false},
		{"Write to XML", func() {}, false},
		{"Read from XML", func() {}, false},
		{"Delete XML", func() {}, false},
	}}

	zipMenu = &Menu{Options: []Option{
		{"<--..", func() {}, true},
		{"Create new archive", func() {}, false},
		{"Write to archive", func() {}, false},
		{"Decompress archive", func() {}, false},
		{"Delete archive", func() {}, false},
	}}
}

type Event rune

type Listener struct {
	EventChan chan Event
}

func (l *Listener) Listen() error {
	c, err := tty.Open()
	if err != nil {
		return err
	}
	defer c.Close()

	for {
		r, err := c.ReadRune()
		if err != nil {
			return err
		}
		l.EventChan <- Event(r)
	}
}

type Object interface {
	Render()
	IsSelectable() bool
}

type Container interface {
	Object
}

type Text struct {
	Text string
}

func (t *Text) Render() {}

func (t *Text) IsSelectable() bool {
	return false
}

type Option struct {
	Text     string
	Func     func()
	Selected bool
}

func (o Option) Render() {}

func (o Option) IsSelectable() bool {
	return true
}

type Menu struct {
	Func     func()
	Selected int
	Options  []Option
}

func (m *Menu) Render() {
	ansi.ClearScreen()
	for _, op := range m.Options {
		op.Render()
		fmt.Println()
	}
}

func (m *Menu) Up() {
	if m.Selected > 0 {
		m.Options[m.Selected].Selected = false
		m.Selected--
		m.Options[m.Selected].Selected = true
	}
}

func (m *Menu) Down() {
	if m.Selected < len(m.Options)-1 {
		m.Options[m.Selected].Selected = false
		m.Selected++
		m.Options[m.Selected].Selected = true
	}
}

func (m *Menu) Select() {
	m.Options[m.Selected].Func()
}
