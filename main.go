package main

import (
	"fmt"
	"learning/textbasics"
)

func main() {
	words := []string{"emeka", "kamsi", "🗓️"}
	fmt.Println(textbasics.LenOfWords(words))
	fmt.Println(len("📌"))
}
