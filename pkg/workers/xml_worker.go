package workers

import (
	"encoding/xml"
	"os"
)

type XMLWorker struct {
	fw *FileWorker
}

func NewXMLWorker(fw *FileWorker) XMLWorker {
	return XMLWorker{
		fw: fw,
	}
}

func (w XMLWorker) Create(path string) error {
	return w.fw.Create(path)
}

func (w XMLWorker) Write(path string, data map[string]string) error {

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		return err
	}

	defer f.Close()

	content, err := xml.Marshal(data)

	f.Write(content)

	return err
}

func (w XMLWorker) Read(path string) (map[string]string, error) {

	content := make(map[string]string)

	data, _ := w.fw.Read(path)

	err := xml.Unmarshal(data, &content)

	return content, err
}

func (w XMLWorker) Delete(path string) error {
	return w.fw.Delete(path)
}
