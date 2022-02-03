package main

import (
	"fmt"
	"io/ioutil"
)

type paragraph struct {
	Data string
}

type Page struct {
	Heading string
	Content []paragraph
}

func readFile() {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(fileContents))
}

func main() {
	readFile()
	fmt.Println("Hello, world!")
}
