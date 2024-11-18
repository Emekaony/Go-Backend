package main

import (
	"learning/textbasics"

	"github.com/kr/pretty"
)

func main() {
	filename := "/Users/shadon./Developer/golang_stuff/learning/textbasics/info_text.txt"
	info, err := textbasics.ParseFileWithRegex(filename)
	if err != nil {
		panic(err.Error())
	}
	pretty.Print(info)
}
