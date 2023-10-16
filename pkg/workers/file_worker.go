package workers

import (
	"os"
)

type FileWorker struct {
	WD string
}

func NewFileWorker(wd string) *FileWorker {
	return &FileWorker{
		WD: wd + "/",
	}
}

/* func (w *FileWorker) FileInfo(path string) (fileInfo, error) {
	f, err := os.Open(w.WD + path)

	if err != nil {
		return fileInfo{}, err
	}

	var finfo fs.FileInfo
	finfo, err = f.Stat()

	if err != nil {
		return fileInfo{}, err
	}

	return fileInfo{}, nil
} */

// Create new file
func (w *FileWorker) Create(path string) error {
	f, err := os.Create(w.WD + path)

	if err != nil {
		return err
	}

	f.Close()

	return nil
}

// Write data in file
func (w *FileWorker) Write(path string, data []byte) error {

	f, err := os.OpenFile(w.WD+path, os.O_APPEND|os.O_WRONLY, 0)

	if err != nil {
		return err
	}

	f.Write(data)

	f.Close()

	return err
}

// Read data from file
func (w *FileWorker) Read(path string) ([]byte, error) {

	return os.ReadFile(w.WD + path)
}

// Delete file
func (w *FileWorker) Delete(path string) error {
	return os.Remove(w.WD + path)
}
