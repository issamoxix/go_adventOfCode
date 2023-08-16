package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func get_file() string {
	body, err := os.ReadFile("file.txt")
	check(err)
	return string(body)
}

func main() {
	var textFile string = get_file()
	var lines []string = strings.Split(textFile, "\n")
	var result []int

	for _, line := range lines {
		var halflen int = (len(line)) / 2
		var _half string = line[:halflen+1]
		var __half string = line[halflen:]
		for _, i := range _half {
			if strings.Contains(__half, string(i)) {
				// fmt.Println(string(i), getScore(string(i)))
				result = append(result, getScore(string(i)))
				break
			}
		}
	}
	fmt.Println(get_sum(result))
}

func getScore(ch string) int {
	var r []rune = []rune(ch)
	var lowerString string = strings.ToLower(ch)

	var abcs []string = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for i := range abcs {
		if abcs[i] == lowerString {
			if unicode.IsUpper(r[0]) {
				return (i + 1) + 26
			}
			return i + 1
		}
	}
	return 0
}

func get_sum(scores []int) int {
	var sum int
	for i := range scores {
		sum += scores[i]
	}
	return sum
}
