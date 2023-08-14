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

func get_file() string {
	body, err := os.ReadFile("file.txt")
	check(err)
	return string(body)
}

type Dictionary map[string]interface {
}

func get_sum(scores []int) int {
	var sum int
	for i := range scores {
		sum += scores[i]
	}
	return sum
}

func main() {
	fmt.Println("Hello World")
	var body string = get_file()
	var games []string = strings.Split(body, "\n")

	scores := []Dictionary{
		{"A": "rock", "B": "papper", "C": "scissor"},
		{"X": "rock", "Y": "papper", "Z": "scissor"},
		{"rock": 1, "papper": 2, "scissor": 3},
	}
	var played_game []int

	for _, game := range games {
		var p []string = strings.Split(game, " ")
		var p1 string = scores[0][p[0]].(string)
		var p2 string = scores[1][strings.TrimSuffix(p[1], "\r")].(string)
		var hand int = scores[2][p2].(int)
		played_game = append(played_game, WL_get_score(p1, p2)+hand)
	}
	fmt.Println(get_sum(played_game))

}

func WL_get_score(p1 string, p2 string) int {

	if p1 == p2 {
		return 3
	}
	if p2 == "rock" {
		if p1 == "scissor" {
			return 6
		}
	}
	if p2 == "papper" {
		if p1 == "rock" {
			return 6
		}
	}

	if p2 == "scissor" {
		if p1 == "papper" {
			return 6
		}
	}
	return 0
}
