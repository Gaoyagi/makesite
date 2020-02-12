package main

import(
	//"fmt"
	"strings"
	"html/template"
	"io/ioutil"
	//"os"
	"bytes"
	//"flag"
	"path/filepath"
	//"context"

	//"cloud.google.com/go/translate"
	//"google.golang.org/api/option"
	"github.com/gomarkdown/markdown"
)

type info struct {
	Content string
} 

func main() {
	//file := flag.String("file", "first-post.txt", "flag for taking in file name")
	//dir := flag.String("dir", "hello", "flag for taking in directory ")
	//after all the flag values  have been taken in you have to parse themg
	//flag.Parse()

	//v1.1 main stuff
	// names := openDir(*dir)
	// for _, file := range names{
	// 	fmt.Println(file+".txt")
	// 	contents := openFile(file+".txt")
	// 	renderFile(contents)
	// 	writeFile((renderFile(contents)), file+".html")
	// }
	
	//v1.0 main stuff
	//contents := openFile(*file)
	//fmt.Print(contents)
	//renderFile(contents)
	//writeFile((renderFile(contents)), "first-post.html")
	//writeFile((renderFile(contents)), "last-post.html")

	//v1.2 stuff
	contents := openFile("README.md")
	convertHtml(contents, "README")
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

//opens a directory and goes through all the files and finds .txt files and returns the names of them
func openDir(dirName string) []string{
	//newstring sub slice array
	textFiles := make([]string, 0)
	//opens directory
	files, err := ioutil.ReadDir(dirName)
    if err != nil {
        panic(err)
	}
	//parses through all the files of the diretory
    for _, file := range files {
		//if the file is a .txt file, add the nameof thefile to the string sub slice  without the .txt extension
		if filepath.Ext(file.Name()) == ".txt"{
			textFiles = append(textFiles, strings.TrimSuffix(file.Name(), ".txt"))
		}
	}
	return textFiles
}

//converts .md files into html files
func convertHtml(content string, fileName string) {
	md := (renderFile(content))
	html := markdown.ToHTML(md, nil, nil)
	writeFile(html, fileName+".html")
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
	err := t.Execute(buffer, info{ Content:fileContents })

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


