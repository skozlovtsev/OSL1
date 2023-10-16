package workers

import (
	"encoding/json"
	"os"
)

type JSONWorker struct {
	fw *FileWorker
}

func NewJSONWorker(fw *FileWorker) JSONWorker {
	return JSONWorker{
		fw: fw,
	}
}

func (w JSONWorker) Create(path string) error {
	w.fw.Create(path)

	return w.fw.Create(path)
}

func (w JSONWorker) Write(path string, data map[string]string) error {

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		return err
	}

	defer f.Close()

	content, err := json.Marshal(data)

	f.Write(content)

	return err
}

/* func (w JSONWorker) Read(path string) (map[string]any, error) {
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
} */

func (w JSONWorker) Read(path string) (map[string]string, error) {

	content := make(map[string]string)

	data, _ := w.fw.Read(path)

	err := json.Unmarshal(data, &content)

	return content, err
}

func (w JSONWorker) Delete(path string) error {
	return w.fw.Delete(path)
}
