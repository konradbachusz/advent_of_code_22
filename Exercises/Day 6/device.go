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
	fmt.Println(char_count)
	// _ = s
	//test_str := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
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

		// fmt.Println("Character", string(character))
		// fmt.Println("window_string:", string(window_string))

	}
	char_count = char_count + 1
	fmt.Println("Processed characters", char_count)

}

/*Steps
1. Read the data
2. Initalize char_count
3. Create a loop to iterate through characters
4. Create a window of 4
5. Check if characters in window are unique:
	if no update char_count
	If yes stop*/
