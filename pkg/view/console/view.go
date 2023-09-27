package console

import (
	"bufio"
	"fmt"
	"os"
	"osl1/pkg/workers"
)

var (
	fileWorker *workers.FileWorker
	jsonWorker workers.JSONWorker
	xmlWorker  workers.XMLWorker
	zipWorker  workers.ZipWorker
)

func init() {
	wd, _ := os.Getwd()

	fileWorker = workers.NewFileWorker(wd)
	jsonWorker = workers.NewJSONWorker(fileWorker)
	xmlWorker = workers.NewXMLWorker(fileWorker)
	zipWorker = workers.NewZipWorker(fileWorker)
}

func ShowData(data []byte) {
	fmt.Println(string(data))
}

func Menu() error {
	fmt.Printf("1. Show drive Info\n2. Operations with files\n3. Operations with JSON\n4. Operations with XML\n5. Operations with ZIP\n")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	fmt.Println("Please choose number of operation: ")

	if err != nil {
		return err
	}

	switch char {
	case '1':
		workers.DriveInfo()
		Menu()
	case '2':
		fmt.Printf("1. Create new file\n2. Write to file\n3. Read from file\n4. Delete file\n")
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		fmt.Println("Please choose number of operation: ")

		if err != nil {
			return err
		}

		switch char {
		case '1':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			fileWorker.Create(data)
			Menu()
		case '2':
			println("Write file name")
			name_reader := bufio.NewReader(os.Stdin)
			name, _ := name_reader.ReadString('\n')
			println("Write text")
			data_reader := bufio.NewReader(os.Stdin)
			data, _ := data_reader.ReadBytes('\n')
			fileWorker.Write(name, data)
			Menu()
		case '3':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			rdata, _ := fileWorker.Read(data)
			ShowData(rdata)
			Menu()
		case '4':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			fileWorker.Delete(data)
			Menu()
		}

	case '3':
		fmt.Printf("1. Create new JSON\n2. Write to JSON\n3. Read from JSON\n4. Delete JSON\n")
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		fmt.Println("Please choose number of operation: ")

		if err != nil {
			return err
		}

		switch char {
		case '1':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			jsonWorker.CreateFile(data)
			Menu()
		case '2':
			println("Write file name")
			name_reader := bufio.NewReader(os.Stdin)
			name, _ := name_reader.ReadString('\n')
			println("Write text")
			data_reader := bufio.NewReader(os.Stdin)
			data, _ := data_reader.ReadBytes('\n')
			jsonWorker.Write(name, data)
			Menu()

		case '3':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			rdata, _ := jsonWorker.Read(data)
			ShowData(rdata)
			Menu()
		case '4':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			jsonWorker.Delete(data)
			Menu()
		}

	case '4':
		fmt.Printf("1. Create new XML\n2. Write to XML\n3. Read from XML\n4. Delete XML\n")
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		fmt.Println("Please choose number of operation: ")

		if err != nil {
			return err
		}

		switch char {
		case '1':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			xmlWorker.CreateFile(data)
			Menu()
		case '2':
			println("Write file name")
			name_reader := bufio.NewReader(os.Stdin)
			name, _ := name_reader.ReadString('\n')
			println("Write text")
			data_reader := bufio.NewReader(os.Stdin)
			data, _ := data_reader.ReadBytes('\n')
			xmlWorker.Write(name, data)
			Menu()

		case '3':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			rdata, _ := xmlWorker.Read(data)
			ShowData(rdata)
			Menu()
		case '4':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			xmlWorker.Delete(data)
			Menu()
		}

	case '5':
		fmt.Printf("1. Create new archive\n2. Write to archive\n3. Decompress archive\n4. Delete archive\n")
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		fmt.Println("Please choose number of operation: ")

		if err != nil {
			return err
		}
		switch char {
		case '1':
			println("Write archive name")
			reader := bufio.NewReader(os.Stdin)
			name, _ := reader.ReadString('\n')
			var NFile []workers.File
			for {
				println("Write file name")
				reader := bufio.NewReader(os.Stdin)
				fName, _ := reader.ReadString('\n')
				if fName == "0" {
					break
				}
				fileBody, _ := fileWorker.Read(fName)
				NFile = append(NFile, workers.File{Name: name, Body: fileBody})
			}
			zipWorker.Compress(name, NFile)
		case '2':
			println("Write archive name")
			reader := bufio.NewReader(os.Stdin)
			name, _ := reader.ReadString('\n')
			var NFile workers.File
			println("Write file name")
			freader := bufio.NewReader(os.Stdin)
			fName, _ := freader.ReadString('\n')
			if fName == "0" {
				break
			}
			fileBody, _ := fileWorker.Read(fName)
			NFile = workers.File{Name: name, Body: fileBody}

			zipWorker.AddFile(name, NFile)
		case '3':
			println("Write archive name")
			reader := bufio.NewReader(os.Stdin)
			name, _ := reader.ReadString('\n')
			Data, _ := zipWorker.Decompress(name)
			for _, v := range Data {
				fileWorker.Create(v.Name)
				fileWorker.Write(v.Name, v.Body)
			}
		case '4':
			println("Write archive name")
			reader := bufio.NewReader(os.Stdin)
			name, _ := reader.ReadString('\n')
			zipWorker.Delete(name)
		}
		Menu()
	}

	return nil
}
