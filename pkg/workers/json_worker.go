package workers

import (
	"encoding/json"
)

type JSONWorker struct {
	fw *FileWorker
}

func NewJSONWorker(fw *FileWorker) *JSONWorker {
	return &JSONWorker{
		fw: fw,
	}
}

func (w *JSONWorker) CreateFile() error {}

func (w *JSONWorker) CreateObject() error {}

func (w *JSONWorker) Read(path string) (map[string]any, error) {
	v := make(map[string]any)

	data, err := w.fw.Read(path)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &v)

	if err != nil {
		return nil, err
	}

	return v, nil
}

func (w *JSONWorker) Delete(path string) error {
	return w.fw.Delete(path)
}
