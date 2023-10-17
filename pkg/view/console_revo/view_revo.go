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
		{"<--..", func() {}},
		{"Show drive Info", func() {}},
		{"Operations with files", func() {}},
		{"Operations with JSON", func() {}},
		{"Operations with XML", func() {}},
		{"Operations with ZI", func() {}},
	}}

	fileMenu = &Menu{Options: []Option{
		{"<--..", func() {}},
		{"Create new file", func() {}},
		{"Write to file", func() {}},
		{"Read from file", func() {}},
		{"Delete file", func() {}},
	}}

	jsonMenu = &Menu{Options: []Option{
		{"<--..", func() {}},
		{"Create new JSON", func() {}},
		{"Write to JSON", func() {}},
		{"Read from JSON", func() {}},
		{"Delete JSON", func() {}},
	}}

	xmlMenu = &Menu{Options: []Option{
		{"<--..", func() {}},
		{"Create new XML", func() {}},
		{"Write to XML", func() {}},
		{"Read from XML", func() {}},
		{"Delete XML", func() {}},
	}}

	zipMenu = &Menu{Options: []Option{
		{"<--..", func() {}},
		{"Create new archive", func() {}},
		{"Write to archive", func() {}},
		{"Decompress archive", func() {}},
		{"Delete archive", func() {}},
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
}

type Option struct {
	Text string
	Func func()
}

func (o Option) Render() {}

type Menu struct {
	Func     func()
	Selected int
	Options  []Option
}

func (m *Menu) Render() {
	ansi.ClearScreen()
	for i, op := range m.Options {
		if i == m.Selected {
			ansi.SetReversed(true)
			defer ansi.SetReversed(false)
		}
		fmt.Printf("%s\n", op.Text)
	}
}

func (m *Menu) Up() {
	if m.Selected > 0 {
		m.Selected--
	}
}

func (m *Menu) Down() {
	if m.Selected < len(m.Options)-1 {
		m.Selected++
	}
}

func (m *Menu) Select() {
	m.Options[m.Selected].Func()
}
