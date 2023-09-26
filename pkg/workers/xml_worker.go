package workers

type XMLWorker struct {
	fw *FileWorker
}

func NewXMLWorker(fw *FileWorker) XMLWorker {
	return XMLWorker{
		fw: fw,
	}
}

func (w XMLWorker) CreateFile(path string) error {
	return w.fw.Create(path)
}

func (w XMLWorker) AddData() {}

func (w XMLWorker) Read(path string) ([]byte, error) {
	return w.fw.Read(path)
}

func (w XMLWorker) Delete(path string) error {
	return w.fw.Delete(path)
}
