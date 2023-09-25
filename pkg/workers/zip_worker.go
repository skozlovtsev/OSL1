package workers

import (
	"archive/zip"
	"os"
)

type ZipWorker struct {
	fw *FileWorker
}

func NewZipWorker(fw *FileWorker) ZipWorker {
	return ZipWorker{
		fw: fw,
	}
}

func (w ZipWorker) Compress(archiveName string, files []file) (int, error) {
	archive, err := os.Create(archiveName)

	if err != nil {
		return -1, err
	}

	// Create a new zip archive.
	zw := zip.NewWriter(archive)

	var i int = 0
	// Add some files to the archive.
	for _, file := range files {
		f, err := zw.Create(file.Name)
		if err != nil {
			return i, err
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			return i, err
		}
		i++
	}

	// Make sure to check the error on Close.
	err = zw.Close()
	if err != nil {
		return -1, err
	}

	return i, nil
}

func (w ZipWorker) AddFile(archiveName string, file string) error {}

//func (w *ZipWorker) Info() {}

func (w ZipWorker) Decompress(archiveName string) ([]file, error) {
	r, err := zip.OpenReader(archiveName)

	if err != nil {
		return nil, err
	}

	defer r.Close()

	var files []file

	for _, f := range r.File {
		rc, err := f.Open()

		if err != nil {
			return nil, err
		}

		var buf []byte

		_, err = rc.Read(buf)

		if err != nil {
			return nil, err
		}

		files = append(files, file{Name: f.Name, Body: buf})

		rc.Close()
	}

	return files, nil
}

func (w ZipWorker) Delete(archiveName string) error {
	return w.fw.Delete(archiveName)
}

type file struct {
	Name string
	Body []byte
}
