package console

import (
	"bufio"
	"fmt"
	"os"
	"osl1/pkg/workers"

	"github.com/skozlovtsev/OSL1/pkg/workers"
)

type ConsoleView struct {
	fsWorker   iFSWorker
	fileWorker iFileWorker
	jsonWorker iJSONWorker
	xmlWorker  iXMLWorker
	zipWorker  iZipWorker
}

func ShowData(data []byte, err error) error {
	if err != nil {
		return err
	}
	fmt.Print(data)
	return nil
}

func Menu() error {
	fmt.Printf("1. Show drive Info\n2. Operations with files\n3. Operations with JSON\n4. Operations with XML\n5. Operations with ZIP\n")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	fmt.Println("Please choose number of operation: ")

	if err != nil {
		return err
	}

	wd, _ := os.Getwd()
	FileWorker := workers.NewFileWorker(wd)

	switch char {
	case '1':
		//workers.NewFSWorker(????).DriveInfo()
	case '2':
		fmt.Printf("1. Create new file\n2. Write to file\n3. Read from file\n4. Delete file\n")
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		fmt.Print("Please choose number of operation: ")

		if err != nil {
			return err
		}

		switch char {
		case '1':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			FileWorker.Create(data)

		case '2':
			println("Write file name")
			name_reader := bufio.NewReader(os.Stdin)
			name, _ := name_reader.ReadString('\n')
			println("Write text")
			data_reader := bufio.NewReader(os.Stdin)
			data, _ := data_reader.ReadBytes('\n')
			FileWorker.Write(name, data)

		case '3':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			ShowData(FileWorker.Read(data))

		case '4':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			FileWorker.Delete(data)
		}

	case '3':
		fmt.Printf("1. Create new JSON\n2. Write to JSON\n3. Read from JSON\n4. Delete JSON\n")
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		fmt.Print("Please choose number of operation: ")

		if err != nil {
			return err
		}

		switch char {
		case '1':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			workers.JsonWorker.CreateFile(reader.ReadLine())
		case '2':
			println("Write file name")
			name_reader := bufio.NewReader(os.Stdin)
			println("Write text")
			data_reader := bufio.NewReader(os.Stdin)
			workers.JsonWorker.Write(name_reader.ReadLine(), data_reader.ReadLine())

		case '3':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			ShowData(workers.JsonWorker.Read(reader.ReadLine()), nil)
		case '4':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			workers.JsonWorker.Delete(reader.ReadLine())
		}

	case '4':
		fmt.Printf("1. Create new XML\n2. Write to XML\n3. Read from XML\n4. Delete XML\n")
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		fmt.Print("Please choose number of operation: ")

		if err != nil {
			return err
		}

		switch char {
		case '1':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			workers.XMLWorker.CreateFile(reader.ReadLine())
		case '2':
			println("Write file name")
			name_reader := bufio.NewReader(os.Stdin)
			println("Write text")
			data_reader := bufio.NewReader(os.Stdin)
			workers.XMLWorker.Write(name_reader.ReadLine(), data_reader.ReadLine())

		case '3':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			ShowData(workers.XMLWorker.Read(reader.ReadLine()), nil)
		case '4':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			workers.XMLWorker.Delete(reader.ReadLine())
		}

	case '5':
	}
}

type file struct {
	Name string
	Body []byte
}

type iFileWorker interface {
	FileInfo(string)
	Create(string) error
	Write(string, []byte) (int, error)
	Read(string) ([]byte, error)
	Delete(string) error
}

type iFSWorker interface {
	DriveInfo(string)
}

type iJSONWorker interface {
	CreateFile() error
	CreateObject() error
	Read(string) (map[string]any, error)
	Delete(string) error
}

type iXMLWorker interface {
	CreateFile(string) error
	AddData(string, []byte) (int, error)
	Read(string) ([]byte, error)
	Delete(string) error
}

type iZipWorker interface {
	Compress(files []file) int
	AddFile()
	Info()
	Decompress()
	Delete()
}
