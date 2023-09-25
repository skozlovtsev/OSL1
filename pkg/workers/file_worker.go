package workers

import (
	"os"
)

type FileWorker struct {
	WD string
}

type fileInfo struct{}

func NewFileWorker(wd string) FileWorker {
	return FileWorker{
		WD: wd,
	}
}

func (w FileWorker) FileInfo(path string) (fileInfo, error) {
	f, err := os.Open(w.WD + path)

	if err != nil {
		return fileInfo{}, err
	}

	finfo, err := f.Stat()

	if err != nil {
		return fileInfo{}, err
	}
}

func (w FileWorker) Create(path string) error {
	f, err := os.Create(w.WD + path)

	if err != nil {
		return err
	}

	f.Close()

	return nil
}

func (w FileWorker) Write(path string, data []byte) (int, error) {
	f, err := os.Open(w.WD + path)

	if err != nil {
		return 0, err
	}

	defer f.Close()

}

func (w FileWorker) Read(path string) ([]byte, error) {
	f, err := os.Open(w.WD + path)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	var data []byte

	f.Read(data)

	return data, nil
}

func (w FileWorker) Delete(path string) error {
	return os.Remove(w.WD + path)
}
