package console

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"osl1/pkg/workers"
	"time"

	ansi "github.com/bit101/go-ansi"
	tty "github.com/mattn/go-tty"
)

const (
	txtPostfix = ".txt"

	jsonPostfix = ".json"

	xmlPostfix = ".xml"

	zipPostfix = ".zip"
)

var (
	fileWorker *workers.FileWorker

	jsonWorker workers.JSONWorker

	xmlWorker workers.XMLWorker

	zipWorker workers.ZipWorker
)

func init() {
	wd, _ := os.Getwd()

	fileWorker = workers.NewFileWorker(wd)

	jsonWorker = workers.NewJSONWorker(fileWorker)

	xmlWorker = workers.NewXMLWorker(fileWorker)

	zipWorker = workers.NewZipWorker(fileWorker)
}

func showData(data []byte) {
	fmt.Println(string(data))
}

func MainMenu() error {

	const ops = 5

	var mainMenuOp int

	ansi.ClearScreen()

	for {
		ansi.Println(ansi.Cyan, fileWorker.WD)
		ansi.Print(ansi.Yellow, "|  0. <--..\n")
		fmt.Print("|  1. Show drive Info\n|  2. Operations with files\n|  3. Operations with JSON\n|  4. Operations with XML\n|  5. Operations with ZIP\n")
		fmt.Println(mainMenuOp)

		r, _ := readKeyRuneTTY()

		switch r {
		default:
			continue
		case 'w': // Up Arrow
			if mainMenuOp > 0 {
				mainMenuOp -= 1
			}
		case 's': // Down Arrow
			if mainMenuOp < ops {
				mainMenuOp += 1
			}
		case 13: // Enter
			switch mainMenuOp {
			case 0:
				return nil
			case 1:
				workers.DriveInfo()
				continue
			case 2:
				fileMenu()
				ansi.ClearScreen()
			case 3:
				jsonMenu()
				ansi.ClearScreen()
			case 4:
				xmlMenu()
				ansi.ClearScreen()
			case 5:
				zipMenu()
				ansi.ClearScreen()
			default:
				ansi.ClearLine()
				ansi.Print(ansi.Red, "Wrong operation. ")
			}
		}
	}
}

func fileMenu() error {
	ansi.ClearScreen()
	ansi.Println(ansi.Cyan, fileWorker.WD)
	ansi.Print(ansi.Yellow, "|  0. <--..\n")
	fmt.Print("|  1. Create new file\n|  2. Write to file\n|  3. Read from file\n")
	ansi.Print(ansi.Red, "|  4. Delete file\n")

	for {
		char, _ := readByte()

		switch char {
		case '0':
			return nil
		case '1':
			name, _ := readFileNameTTY(txtPostfix)
			return fileWorker.Create(name)
		case '2':
			path, _ := readFileNameTTY(txtPostfix)
			data, _ := readString()
			return fileWorker.Write(path, []byte(data))
		case '3':
			path, _ := readFileNameTTY(txtPostfix)
			data, err := fileWorker.Read(path)
			fmt.Println(string(data))
			time.Sleep(time.Second * 10)
			return err
		case '4':
			name, _ := readFileName(txtPostfix)
			return fileWorker.Delete(name)
		default:
			continue
		}
	}
}

func jsonMenu() error {
	ansi.ClearScreen()
	ansi.Println(ansi.Cyan, fileWorker.WD)
	ansi.Print(ansi.Yellow, "|  0. <--..\n")
	fmt.Print("|  1. Create new JSON\n|  2. Write to JSON\n|  3. Read from JSON\n")
	ansi.Print(ansi.Red, "|  4. Delete JSON\n")

	for {
		char, _ := readByte()

		switch char {
		case '0':
			return nil
		case '1':
			name, _ := readFileName(jsonPostfix)
			return jsonWorker.Create(name)
		case '2':
			path, _ := readFileNameTTY(jsonPostfix)
			var data map[string]string = make(map[string]string)

			for {
				k, err := readStringTTY()
				if err != nil {
					jsonWorker.Write(path, data)
				}

				fmt.Print(" : ")

				var v string

				v, err = readStringTTY()

				if err != nil {
					jsonWorker.Write(path, data)
				}

				data[k] = v
			}
		case '3':
			path, _ := readFileNameTTY(jsonPostfix)
			data, err := jsonWorker.Read(path)

			if err != nil {
				fmt.Println(err.Error())
				return err
			}

			for k, v := range data {
				fmt.Printf("\n%s: %s", k, v)
			}
		case '4':
			name, _ := readFileName(jsonPostfix)
			return jsonWorker.Delete(name)
		default:
			continue
		}
	}
}

func xmlMenu() error {
	ansi.ClearScreen()
	ansi.Println(ansi.Cyan, fileWorker.WD)
	ansi.Print(ansi.Yellow, "|  0. <--..\n")
	fmt.Print("|  1. Create new XML\n|  2. Write to XML\n|  3. Read from XML\n")
	ansi.Print(ansi.Red, "|  4. Delete XML\n")

	for {
		char, _ := readByte()

		switch char {
		case '0':
			return nil
		case '1':
			name, _ := readFileName(xmlPostfix)
			return xmlWorker.Create(name)
		case '2':
			path, _ := readFileNameTTY(xmlPostfix)
			var data workers.XMLFile = workers.XMLFile{}

			for {
				k, err := readStringTTY()
				if err != nil {
					xmlWorker.Write(path, data)
				}

				fmt.Print(" : ")

				var v string

				v, err = readStringTTY()

				if err != nil {
					xmlWorker.Write(path, data)
				}

				data.Object = append(data.Object, struct {
					Id   string
					Text string
				}{k, v})
			}
		case '3':
			path, _ := readFileNameTTY(xmlPostfix)
			data, err := xmlWorker.Read(path)

			if err != nil {
				fmt.Println(err.Error())
				return err
			}

			for _, obj := range data.Object {
				fmt.Printf("\n%s: %s", obj.Id, obj.Text)
			}
		case '4':
			name, _ := readFileName(xmlPostfix)
			return xmlWorker.Delete(name)
		default:
			continue
		}
	}
}

func zipMenu() error {
	ansi.ClearScreen()
	ansi.Println(ansi.Cyan, fileWorker.WD)
	ansi.Print(ansi.Yellow, "|  0. <--..\n")
	fmt.Print("|  1. Create new archive\n|  2. Write to archive\n|  3. Decompress archive\n")
	ansi.Print(ansi.Red, "|  4. Delete archive\n")

	for {
		char, _ := readByte()

		switch char {
		case '0':
			return nil
		case '1':
			name, _ := readFileName(zipPostfix)

			var files []workers.File

			for {
				fName, _ := readFileName("")

				if fName == "0" {
					break
				}

				fileBody, _ := fileWorker.Read(fName)

				files = append(files, workers.File{Name: name, Body: fileBody})
			}
			return zipWorker.Compress(name, files)
		case '2':

			name, _ := readFileNameTTY(zipPostfix)
			var NFile workers.File

			fName, _ := readFileNameTTY("")
			if fName == "" {
				break
			}
			fileBody, _ := fileWorker.Read(fName)
			NFile = workers.File{Name: name, Body: fileBody}

			zipWorker.AddFile(name, NFile)
		case '3':

			name, _ := readFileNameTTY(zipPostfix)
			Data, _ := zipWorker.Decompress(name)
			for _, v := range Data {
				fileWorker.Create(v.Name)
				fileWorker.Write(v.Name, v.Body)
			}
		case '4':
			name, _ := readFileName(zipPostfix)
			return zipWorker.Delete(name)
		default:
			continue
		}
	}
}

func readByte() (byte, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please choose number of operation: ")

	return reader.ReadByte()
}

func readString() (string, error) {
	fmt.Print("Write file name: ")

	reader := bufio.NewReader(os.Stdin)

	return reader.ReadString('\n')
}

func readFileName(formatPostfix string) (string, error) {
	fmt.Printf("Write file name: %s", formatPostfix)

	ansi.MoveLeft(len(formatPostfix))

	ansi.SetReversed(true)

	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')

	ansi.SetReversed(false)

	return name + formatPostfix, err
}

func readKeyRuneTTY() (rune, error) {
	tty, err := tty.Open()
	if err != nil {
		return '0', err
	}
	defer tty.Close()

	return tty.ReadRune()
}

func readStringTTY() (string, error) {
	tty, err := tty.Open()
	if err != nil {
		return "", err
	}
	defer tty.Close()

	var s string

	for {
		r, err := tty.ReadRune()
		if err != nil {
			return "", err
		}

		// handle key event
		switch {
		case r == 27: //ESC
			return "", errors.New("ESC character: end of input")
		case r == 13: //Enter
			return s, nil
		case r == 127 || r == 8: //Backspace
			if len(s) >= 1 {
				s = s[:len(s)-1]
			}

			ansi.ClearLine()

			ansi.SetReversed(true)
			fmt.Print(s)

			ansi.SetReversed(false)

			ansi.MoveLeft(1)

		case 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || '0' <= r && r <= '9':
			s += string(r)

			ansi.ClearLine()

			ansi.SetReversed(true)
			fmt.Print(s)

			ansi.SetReversed(false)

			ansi.MoveLeft(1) // !
		}
	}
}

func readFileNameTTY(formatPostfix string) (string, error) {
	tty, err := tty.Open()
	if err != nil {
		return "", err
	}
	defer tty.Close()

	fmt.Printf("Write file name:  %s", formatPostfix)

	ansi.MoveLeft(len(formatPostfix) + 1)

	var s string

	for {
		r, err := tty.ReadRune()
		if err != nil {
			return "", err
		}

		// handle key event
		switch {
		case r == 13: //Enter
			return s + formatPostfix, nil
		case r == 127 || r == 8: //Backspace
			if len(s) >= 1 {
				s = s[:len(s)-1]
			}

			ansi.ClearLine()

			if len(s) > 1 {
				fmt.Print("Write file name: ")
			} else {
				fmt.Print("Write file name:  ")
			}

			ansi.SetReversed(true)
			fmt.Print(s)

			ansi.SetReversed(false)
			fmt.Print(formatPostfix)

			ansi.MoveLeft(len(formatPostfix) + 1)

		case 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z':
			s += string(r)

			ansi.ClearLine()
			fmt.Print("Write file name: ")

			ansi.SetReversed(true)
			fmt.Print(s)

			ansi.SetReversed(false)
			fmt.Print(formatPostfix)

			ansi.MoveLeft(len(formatPostfix) + 1) // !
		}
	}
}

/*func Menu() error {
	switch char {
	case '1':
		workers.DriveInfo()
		Menu()
	case '2':
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
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		fmt.Println("Please choose number of operation: ")

		if err != nil {
			return err
		}

		switch char {
		case '1':
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
}*/
