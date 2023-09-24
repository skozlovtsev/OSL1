package workers

type XMLWorker struct {
	fw *FileWorker
}

func NewXMLWorker(fw *FileWorker) *XMLWorker {
	return &XMLWorker{
		fw: fw,
	}
}

func (w *XMLWorker) CreateFile() {}

func (w *XMLWorker) AddData() {}

func (w *XMLWorker) Read() {}

func (w *XMLWorker) Delete(path string) error {
	return w.fw.Delete(path)
}
