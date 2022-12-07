//https://adventofcode.com/2022/day/5

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"
)

func str_to_int(str_number string) int {
	number, err := strconv.Atoi(str_number)

	_ = err

	return number
}
func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func SplitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}

func IsLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

func main() {

	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	s := string(content)

	//List of pairs
	rows := strings.Split(strings.Replace(s, "\r\n", "\n", -1), "\n")
	crates := rows[:8]

	actions := rows[10:len(rows)]

	stacks := make([]string, 9)
	for n, row := range stacks {
		_ = row
		for i, row := range crates {
			_ = row
			value := SplitSubN(crates[i], 4)[n]
			value = strings.Replace(value, "[", "", 1)
			value = strings.Replace(value, "]", "", 1)
			value = strings.Replace(value, " ", "", 1)
			if IsLetter(value) {
				stacks[n] = stacks[n] + value
			}

		}
	}

	split_Action := strings.Split(actions[0], " ")

	for i, action := range actions {
		_ = i
		split_Action := strings.Split(action, " ")
		quantity := str_to_int(split_Action[1])
		departure_stack := str_to_int(split_Action[3]) - 1
		destination_stack := str_to_int(split_Action[5]) - 1

		for i := 1; i <= quantity; i++ {

			stacks[destination_stack] = stacks[destination_stack][:0] + stacks[departure_stack][0:1] + stacks[destination_stack][0:]
			stacks[departure_stack] = trimFirstRune(stacks[departure_stack])
		}

	}
	fmt.Println("Part 1 Answer:")
	for i, row := range stacks {
		_ = row
		fmt.Println(stacks[i][:1])

	}

	////////Part 2/////////////
	//Re-create original stacks
	stacks = make([]string, 9)
	for n, row := range stacks {
		_ = row
		for i, row := range crates {
			_ = row
			value := SplitSubN(crates[i], 4)[n]
			value = strings.Replace(value, "[", "", 1)
			value = strings.Replace(value, "]", "", 1)
			value = strings.Replace(value, " ", "", 1)
			if IsLetter(value) {
				stacks[n] = stacks[n] + value
			}

		}
	}

	for i, action := range actions {
		_ = i
		split_Action = strings.Split(action, " ")
		quantity := str_to_int(split_Action[1])
		departure_stack := str_to_int(split_Action[3]) - 1
		destination_stack := str_to_int(split_Action[5]) - 1

		stacks[destination_stack] = stacks[departure_stack][:quantity] + stacks[destination_stack]
		stacks[departure_stack] = stacks[departure_stack][quantity:]

	}
	fmt.Println("Part 2 Answer:")
	for i, row := range stacks {
		_ = row
		fmt.Println(stacks[i][:1])

	}

}
