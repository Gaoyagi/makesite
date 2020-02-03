package main

import(
	//"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

type info struct {
	Content string
} 

func main() {
 	contents := openFile("first-post.txt")
	// fmt.Print(contents)
	renderFile(contents)

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

//parses through the file contents and renders them so they written to template (step 2 & 3)
//parses through the file contents and stores it into a byte array buffer for info to be written to a file(Step 4)
func renderFile(fileContents string) {
	// Files are provided as a slice of strings.
	//already existing template file
	paths := []string{
		"template.tmpl",
	}
	//creates new template file and bases it off the existing one
	t := template.Must(template.New("template.tmpl").ParseFiles(paths...))
	err := t.Execute(os.Stdout, info{ Content:fileContents})
	if err != nil {
		panic(err)
	}
}

//takes in the file name you want to write to and and the buffer contating the information you want to write to i


