package workers

import (
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

func (w JSONWorker) Write(path string, object []byte) error {

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		return err
	}

	defer f.Close()

	stat, _ := f.Stat()

	_, err = f.WriteAt(object, stat.Size()-1)

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

func (w JSONWorker) Read(path string) ([]byte, error) {
	return w.fw.Read(path)
}

func (w JSONWorker) Delete(path string) error {
	return w.fw.Delete(path)
}
