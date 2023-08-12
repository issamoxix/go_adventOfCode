package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("hello world")
	var body []byte
	body, err := os.ReadFile("file.txt")
	check(err)
	var fileBody string = string(body)
	var aStrings []string = strings.Split(fileBody, "\n")

	for i := range aStrings {
		fmt.Println(aStrings[i])
	}
}
