package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
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

func writeTemplate(contents string) {
	page := Page{
		Heading: "New Page",
		Content: contents,
	}
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	newfile, err := os.Create("new.html")
	if err != nil {
		panic(err)
	}
	t.Execute(newfile, page)
}

func listTextFiles(directory string) {
	files, err := os.ReadDir(directory)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			fmt.Printf("%v \n", file.Name())
		}
	}
}

func main() {
	// file := flag.String("file", "none.txt", "Enter file name")
	dir := flag.String("dir", ".", "Enter the directory path")

	flag.Parse()

	listTextFiles(*dir)

	// contents := readFile(*file)
	// writeTemplate(contents)
}
