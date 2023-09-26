package console

import (
	"bufio"
	"fmt"
	"os"
	"osl1/pkg/workers"
)

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

	wd, _ := os.Getwd()
	FileWorker := workers.NewFileWorker(wd)
	JSONWorker := workers.NewJSONWorker(&FileWorker)
	XMLWorker := workers.NewXMLWorker(&FileWorker)
	ZipWorker := workers.NewZipWorker(&FileWorker)

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
			FileWorker.Create(data)
			Menu()
		case '2':
			println("Write file name")
			name_reader := bufio.NewReader(os.Stdin)
			name, _ := name_reader.ReadString('\n')
			println("Write text")
			data_reader := bufio.NewReader(os.Stdin)
			data, _ := data_reader.ReadBytes('\n')
			FileWorker.Write(name, data)
			Menu()
		case '3':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			rdata, _ := FileWorker.Read(data)
			ShowData(rdata)
			Menu()
		case '4':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			FileWorker.Delete(data)
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
			JSONWorker.CreateFile(data)
			Menu()
		case '2':
			println("Write file name")
			name_reader := bufio.NewReader(os.Stdin)
			name, _ := name_reader.ReadString('\n')
			println("Write text")
			data_reader := bufio.NewReader(os.Stdin)
			data, _ := data_reader.ReadBytes('\n')
			JSONWorker.Write(name, data)
			Menu()

		case '3':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			rdata, _ := JSONWorker.Read(data)
			ShowData(rdata)
			Menu()
		case '4':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			JSONWorker.Delete(data)
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
			XMLWorker.CreateFile(data)
			Menu()
		case '2':
			println("Write file name")
			name_reader := bufio.NewReader(os.Stdin)
			name, _ := name_reader.ReadString('\n')
			println("Write text")
			data_reader := bufio.NewReader(os.Stdin)
			data, _ := data_reader.ReadBytes('\n')
			XMLWorker.Write(name, data)
			Menu()

		case '3':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			rdata, _ := XMLWorker.Read(data)
			ShowData(rdata)
			Menu()
		case '4':
			println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			XMLWorker.Delete(data)
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
				fileBody, _ := FileWorker.Read(fName)
				NFile = append(NFile, workers.File{Name: name, Body: fileBody})
			}
			ZipWorker.Compress(name, NFile)
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
			fileBody, _ := FileWorker.Read(fName)
			NFile = workers.File{Name: name, Body: fileBody}

			ZipWorker.AddFile(name, NFile)
		case '3':
			println("Write archive name")
			reader := bufio.NewReader(os.Stdin)
			name, _ := reader.ReadString('\n')
			Data, _ := ZipWorker.Decompress(name)
			for _, v := range Data {
				FileWorker.Create(v.Name)
				FileWorker.Write(v.Name, v.Body)
			}
		case '4':
			println("Write archive name")
			reader := bufio.NewReader(os.Stdin)
			name, _ := reader.ReadString('\n')
			ZipWorker.Delete(name)
		}
		Menu()
	}

	return nil
}
