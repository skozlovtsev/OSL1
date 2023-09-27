package workers

import (
	"archive/zip"
	"io"
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

// Cereate Zip archive and add files to it
func (w ZipWorker) Compress(archiveName string, files []File) error {

	archive, err := os.Create(w.fw.WD + archiveName)

	if err != nil {
		return err
	}

	// Create a new zip archive.
	zw := zip.NewWriter(archive)

	// Add some files to the archive.
	for _, file := range files {
		f, err := zw.Create(file.Name)
		if err != nil {
			return err
		}
		_, err = f.Write(file.Body)
		if err != nil {
			return err
		}
	}

	// Make sure to check the error on Close.
	err = zw.Close()

	archive.Close()

	if err != nil {
		return err
	}

	return nil
}

// Add File to the archive
func (w ZipWorker) AddFile(archiveName string, file File) error {

	archive, err := os.Open(w.fw.WD + archiveName)

	if err != nil {
		return err
	}

	zw := zip.NewWriter(archive)

	var f io.Writer

	f, err = zw.Create(file.Name)

	if err != nil {
		return err
	}

	_, err = f.Write(file.Body)

	if err != nil {
		return err
	}

	err = zw.Close()

	archive.Close()

	if err != nil {
		return err
	}

	return nil
}

//func (w *ZipWorker) Info() {}

// Decompress files from the zip archive
func (w ZipWorker) Decompress(archiveName string) ([]File, error) {

	r, err := zip.OpenReader(w.fw.WD + archiveName)

	if err != nil {
		return nil, err
	}

	defer r.Close()

	var files []File

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

		files = append(files, File{Name: f.Name, Body: buf})

		rc.Close()
	}
	r.Close()

	return files, nil
}

// Delete aerchive file
func (w ZipWorker) Delete(archiveName string) error {
	return w.fw.Delete(archiveName)
}

// Structure representing file
type File struct {
	Name string
	Body []byte
}
