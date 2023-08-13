package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func get_sum(calories []int) int {
	var sum int
	for i := range calories {
		sum += calories[i]
	}
	return sum
}

func main() {
	// Day 1 Adventofcode

	var body []byte
	body, err := os.ReadFile("file.txt")
	check(err)
	var fileBody string = string(body)
	aStrings := strings.Split(fileBody, "\n")

	// First task & sencond
	var temp int

	var m_calories []int

	for i := range aStrings {
		calory := strings.TrimSuffix(aStrings[i], "\r")
		if len(calory) == 0 {
			m_calories = append(m_calories, temp)
			temp = 0
		} else {
			v, err := strconv.Atoi(calory)
			check(err)
			temp += v
		}
	}
	sort.Ints(m_calories)
	fmt.Println(m_calories[len(m_calories)-1])
	fmt.Println(get_sum(m_calories[len(m_calories)-3 : len(m_calories)]))

}
