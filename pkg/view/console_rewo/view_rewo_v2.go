package consolerewo

import (
	"bufio"
	"fmt"
	"os"
	"osl1/pkg/view/colors"
	"osl1/pkg/workers"
)

var (
	MainMenu Menu
	FileMenu Menu
	JsonMenu Menu
	XMLMenu  Menu
	ZipMenu  Menu

	NextMenu *Menu

	fileWorker *workers.FileWorker

	jsonWorker workers.JSONWorker

	xmlWorker workers.XMLWorker

	zipWorker workers.ZipWorker
)

type Menu []*Option

type Option struct {
	Text   string
	Action func()
}

func init() {
	wd, _ := os.Getwd()

	fileWorker = workers.NewFileWorker(wd)

	jsonWorker = workers.NewJSONWorker(fileWorker)

	xmlWorker = workers.NewXMLWorker(fileWorker)

	zipWorker = workers.NewZipWorker(fileWorker)

	MainMenu = Menu{
		&Option{"<--..", func() { panic("program successfully stopped") }},
		&Option{"Show drive Info", func() {}},
		&Option{"Operations with files", func() { NextMenu = &FileMenu }},
		&Option{"Operations with JSON", func() { NextMenu = &JsonMenu }},
		&Option{"Operations with XML", func() { NextMenu = &XMLMenu }},
		&Option{"Operations with ZI", func() { NextMenu = &ZipMenu }},
	}

	FileMenu = Menu{
		&Option{"<--..", func() {
			NextMenu = &MainMenu
		}},
		&Option{"Create new file", func() {
			name, _ := readString()
			fileWorker.Create(name)
		}},
		&Option{"Write to file", func() {
			path, _ := readString()
			data, _ := readBytes()
			fileWorker.Write(path, data)
		}},
		&Option{"Read from file", func() {
			path, _ := readString()
			data, _ := fileWorker.Read(path)
			fmt.Println(string(data))
		}},
		&Option{"Delete file", func() {
			name, _ := readString()
			fileWorker.Delete(name)
		}},
	}

	JsonMenu = Menu{
		&Option{"<--..", func() { NextMenu = &MainMenu }},
		&Option{"Create new JSON", func() {
			name, _ := readString()
			jsonWorker.Create(name)
		}},
		&Option{"Write to JSON", func() {
			path, _ := readString()
			var data map[string]string = make(map[string]string)

			for {
				k, err := readString()
				if err != nil {
					jsonWorker.Write(path, data)
				}

				fmt.Print(" : ")

				var v string

				v, err = readString()

				if err != nil {
					jsonWorker.Write(path, data)
				}

				data[k] = v
			}
		}},
		&Option{"Read from JSON", func() {
			path, _ := readString()
			data, err := jsonWorker.Read(path)

			if err != nil {
				fmt.Println(err.Error())
			}

			for k, v := range data {
				fmt.Printf("\n%s: %s", k, v)
			}
		}},
		&Option{"Delete JSON", func() {
			name, _ := readString()
			jsonWorker.Delete(name)
		}},
	}

	XMLMenu = Menu{
		&Option{"<--..", func() { NextMenu = &MainMenu }},
		&Option{"Create new XML", func() {
			name, _ := readString()
			xmlWorker.Create(name)
		}},
		&Option{"Write to XML", func() {
			path, _ := readString()
			var data workers.XMLFile = workers.XMLFile{}

			for {
				k, err := readString()
				if err != nil {
					xmlWorker.Write(path, data)
				}

				fmt.Print(" : ")

				var v string

				v, err = readString()

				if err != nil {
					xmlWorker.Write(path, data)
				}

				data.Object = append(data.Object, struct {
					Id   string
					Text string
				}{k, v})
			}
		}},
		&Option{"Read from XML", func() {
			path, _ := readString()
			data, err := xmlWorker.Read(path)

			if err != nil {
				fmt.Println(err.Error())
			}

			for _, obj := range data.Object {
				fmt.Printf("\n%s: %s", obj.Id, obj.Text)
			}
		}},
		&Option{"Delete XML", func() {
			name, _ := readString()
			xmlWorker.Delete(name)
		}},
	}

	ZipMenu = Menu{
		&Option{"<--..", func() { NextMenu = &MainMenu }},
		&Option{"Create new archive", func() {
			name, _ := readString()

			var files []workers.File

			for {
				fName, _ := readString()

				if fName == "" {
					break
				}

				fileBody, _ := fileWorker.Read(fName)

				files = append(files, workers.File{Name: name, Body: fileBody})
			}
			zipWorker.Compress(name, files)
		}},
		&Option{"Write to archive", func() {
			name, _ := readString()
			var NFile workers.File

			fName, _ := readString()

			fileBody, _ := fileWorker.Read(fName)
			NFile = workers.File{Name: name, Body: fileBody}

			zipWorker.AddFile(name, NFile)
		}},
		&Option{"Decompress archive", func() {
			name, _ := readString()
			Data, _ := zipWorker.Decompress(name)
			for _, v := range Data {
				fileWorker.Create(v.Name)
				fileWorker.Write(v.Name, v.Body)
			}
		}},
		&Option{"Delete archive", func() {
			name, _ := readString()
			zipWorker.Delete(name)
		}},
	}

	NextMenu = &MainMenu
}

func MainLoop() {
	for {
		printMenu(*NextMenu)

		op := readInt(4)

		if op == -1 {
			continue
		}

		[]*Option(*NextMenu)[op].Action()
	}
}

func printMenu(menu Menu) {
	fmt.Print(colors.ClearScreen)
	for i, v := range menu {
		fmt.Printf("| %d) %s\n", i, v.Text)
	}
}

func readInt(end int) int {
	fmt.Print("Write option: ")

	var i int

	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		return -1
	}
	return i
}

func readString() (string, error) {
	fmt.Print("Write file name: ")

	reader := bufio.NewReader(os.Stdin)

	return reader.ReadString('\n')
}

func readBytes() ([]byte, error) {
	fmt.Print("Write file name: ")

	reader := bufio.NewReader(os.Stdin)

	return reader.ReadBytes('\n')
}
