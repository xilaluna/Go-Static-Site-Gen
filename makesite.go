package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Page struct {
	Heading string
	Content string
}

func readFile(file string) (string){
	fileContents, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return (string(fileContents))
}

func writeFile(contents string) {
	data := []byte(contents)
	err := os.WriteFile("newfile.txt", data, 0644)
	if err != nil {
		panic(err)
	}
}

func createHtmlFile(contents string, filename string) {
	page := Page{
		Heading: "New Page",
		Content: contents,
	}
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	htmlFileName := filename + ".html"
	newfile, err := os.Create(htmlFileName)
	if err != nil {
		panic(err)
	}
	t.Execute(newfile, page)
}

func writeFromDir(directory string) {
	files, err := os.ReadDir(directory)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			contents := readFile(file.Name())
			formatFileName := strings.Split(file.Name(), ".")[0]
			createHtmlFile(contents, formatFileName)
		}
	}
}

func main() {
	// file := flag.String("file", "none.txt", "Enter file name")
	dir := flag.String("dir", ".", "Enter the directory path")

	flag.Parse()

	writeFromDir(*dir)

	// contents := readFile(*file)
	// writeTemplate(contents)
}
