package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello world")
	var body string = read_file()
	var lines []string = strings.Split(body, "\n")
	// fmt.Println(part1(lines)) // uncomment for part1
	fmt.Println(part2(lines))

}

func part1(lines []string) int {
	var result []int
	for _, line := range lines {
		var _elf string = strings.Split(line, ",")[0]
		var __elf string = strings.TrimSuffix(strings.Split(line, ",")[1], "\r")
		if isOverlaping(_elf, __elf) {
			result = append(result, 1)
			fmt.Println(_elf, __elf)
		}
	}
	return get_sum(result)

}

func part2(lines []string) int {
	var result []int
	for _, line := range lines {
		var _elf string = strings.Split(line, ",")[0]
		var __elf string = strings.TrimSuffix(strings.Split(line, ",")[1], "\r")
		if allOverlaping(_elf, __elf) {
			result = append(result, 1)
			fmt.Println(_elf, __elf)
		}
	}
	return get_sum(result)
}

func allOverlaping(_elf string, __elf string) bool {
	var _range []string = strings.Split(_elf, "-")
	var __range []string = strings.Split(__elf, "-")

	// if _range[0] > __range[0] {
	// 	_range, __range = __range, _range
	// }
	// return _range[1] >= __range[0]

	if parseStr(_range[0]) == parseStr(__range[0]) {
		return true
	}
	if parseStr(_range[1]) == parseStr(__range[1]) {
		return true
	}
	if parseStr(_range[0]) == parseStr(__range[1]) {
		return true
	}
	if parseStr(_range[1]) == parseStr(__range[0]) {
		return true
	}
	if parseStr(_range[0]) < parseStr(__range[0]) {
		if parseStr(_range[1]) > parseStr(__range[0]) {
			return true
		}
	}

	if parseStr(_range[0]) > parseStr(__range[0]) {
		if parseStr(_range[1]) < parseStr(__range[1]) {
			return true
		}
		if parseStr(_range[1]) > parseStr(__range[1]) {
			if parseStr(_range[0]) < parseStr(__range[1]) {
				return true
			}
		}
	}
	return false
}

func parseStr(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func isOverlaping(_elf string, __elf string) bool {
	var _range []string = strings.Split(_elf, "-")
	var __range []string = strings.Split(__elf, "-")
	if parseStr(_range[0]) <= parseStr(__range[0]) {
		if parseStr(_range[1]) >= parseStr(__range[1]) {
			return true
		}
	}
	if parseStr(_range[0]) >= parseStr(__range[0]) {
		if parseStr(_range[1]) <= parseStr(__range[1]) {
			return true
		}
	}
	if parseStr(_range[0]) == parseStr(_range[1]) {
		if parseStr(__range[0]) == parseStr(__range[1]) {
			if parseStr(_range[0]) == parseStr(__range[1]) {
				return true
			}
		}

	}
	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_file() string {
	body, err := os.ReadFile("file.txt")
	check(err)
	return string(body)
}

func get_sum(scores []int) int {
	var sum int
	for i := range scores {
		sum += scores[i]
	}
	return sum
}
