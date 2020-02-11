package main

import(
	"fmt"
	"strings"
	"html/template"
	"io/ioutil"
	//"os"
	"bytes"
	"flag"
	"path/filepath"
)

type info struct {
	Content string
} 

func main() {
	//file := flag.String("file", "first-post.txt", "flag for taking in file name")
	dir := flag.String("dir", "hello", "flag for taking in directory ")
	//after all the flag values  have been taken in you have to parse themg
	flag.Parse()

	names := openDir(*dir)
	for _, file := range names{
		fmt.Println(file)
		contents := openFile(file+".txt")
		renderFile(contents)
		writeFile((renderFile(contents)), file+".html")
		
	}
	
	//contents := openFile(*file)
	//fmt.Print(contents)
	//renderFile(contents)
	//writeFile((renderFile(contents)), "first-post.html")
	
	//writeFile((renderFile(contents)), "last-post.html")
}

//opens the file and returns its contents as a string
func openFile(fileName string) string{
	fileContents, err := ioutil.ReadFile(fileName)
	//error checking the file opening
	if err != nil {		
		panic(err)
	}
	return string(fileContents)
}

func openDir(dirName string) []string{
	textFiles := make([]string, 0)
	files, err := ioutil.ReadDir(dirName)
    if err != nil {
        panic(err)
    }
    for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt"{
			textFiles = append(textFiles, strings.TrimSuffix(file.Name(), ".txt"))
		}
	}
	return textFiles
}

//parses through the file contents and renders them so they written to template (step 2 & 3)
//parses through the file contents and stores it into a byte array buffer for info to be written to a file(Step 4)
func renderFile(fileContents string) []byte{
	//list of existing tempaltes
	paths := []string{
		"template.tmpl",
	}

	//buffer object for writing to a new file
	buffer := new(bytes.Buffer)

	//creates new template file? or object? and bases it off the existing one named template.tmpl
	//paths... lets you choose what possible tmpl files you wish to use as a base
	t := template.Must(template.New("template.tmpl").ParseFiles(paths...))

	//executes and writes the contents out to terminal
	// err := t.Execute(os.Stdout, info{ Content:fileContents})

	//executes and writes contents out the buffer 
	err := t.Execute(buffer, info{ Content:fileContents})

	if err != nil {
		panic(err)
	}
	//returns the buffer
	return buffer.Bytes()
}

//takes in the file name you want to write to and and the buffer contating the information you want to write to it
func writeFile(fileContents []byte, fileName string) {
	err := ioutil.WriteFile(fileName, fileContents, 0644)
    if err != nil {
		panic(err)
	}
}


