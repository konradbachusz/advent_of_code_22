//https://adventofcode.com/2022/day/4

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func unique(arr string) bool {
	m := make(map[rune]bool)
	for _, i := range arr {
		_, ok := m[i]
		if ok {
			return false
		}

		m[i] = true
	}

	return true
}

func main() {

	///////////Part 1///////////////

	//Loading file
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	//Initialize count
	char_count := 0

	//Reading string
	s := string(content)

	window_string := ""
	for i, character := range s {
		_ = i
		window_string = window_string + string(character)

		if len(window_string) > 4 {
			window_string = window_string[1:]
		}

		if len(window_string) == 4 && unique(window_string) {
			break
		}

		char_count = char_count + 1

	}
	char_count = char_count + 1
	fmt.Println("Part 1 Solution:", char_count)

	/////////////Part 2///////////////
	window_string = ""
	char_count = 0
	for i, character := range s {
		_ = i
		window_string = window_string + string(character)

		if len(window_string) > 14 {
			window_string = window_string[1:]
		}

		if len(window_string) == 14 && unique(window_string) {
			break
		}

		char_count = char_count + 1

	}
	char_count = char_count + 1
	fmt.Println("Part 2 Solution:", char_count)

}
