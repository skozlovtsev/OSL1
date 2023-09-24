package workers

type ZipWorker struct{}

func (w *ZipWorker) Compress() {}

func (w *ZipWorker) AddFile() error {}

func (w *ZipWorker) Info() {}

func (w *ZipWorker) Decompress() {}

func (w *ZipWorker) Delete() error {}
