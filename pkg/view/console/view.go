package console

import (
	"github.com/rivo/tview"
)

type iFileWorker interface {
	FileInfo(string)
	Create(string) error
	Write(string, []byte) (int, error)
	Read(string) ([]byte, error)
	Delete(string) error
}

type iFSWorker interface {
	DriveInfo(string)
}

type iJSONWorker interface {
	CreateFile() error
	CreateObject() error
	Read(string) (map[string]any, error)
	Delete(string) error
}

type iXMLWorker interface {
	CreateFile(string) error
	AddData(string, []byte) (int, error)
	Read(string) ([]byte, error)
	Delete(string) error
}

type iZipWorker interface {
	Compress()
	AddFile()
	Info()
	Decompress()
	Delete()
}

type ConsoleView struct {
	app        *tview.Application
	fsWorker   iFSWorker
	fileWorker iFileWorker
	jsonWorker iJSONWorker
	xmlWorker  iXMLWorker
	zipWorker  iZipWorker
}

func NewConsoleView(app *tview.Application) *ConsoleView {
	return &ConsoleView{
		app: app,
	}
}

func (c *ConsoleView) Run() error {
	if err := c.app.SetRoot(list, true).EnableMouse(true).Run(); err != nil {
		return err
	}
}
