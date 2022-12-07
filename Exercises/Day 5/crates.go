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

	fmt.Println(crates[0][12:15])

	split_row := SplitSubN(crates[0], 4)
	fmt.Println("Print char split", split_row)
	fmt.Println("Split row value", split_row[3])

	actions := rows[10:len(rows)]
	fmt.Println(actions[0])

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
			// fmt.Println("Value", value, i)

		}
	}
	// stacks[0] = crates[0]

	// for n, row := range crates {
	// 	_ = n
	// 	split_row = SplitSubN(row, 4)
	// 	stacks[i] = split_row[0] + split_row[1] + split_row[2] + split_row[3] + split_row[4] + split_row[5] + split_row[6] + split_row[7] + split_row[8]
	// }

	fmt.Println("stacks1", stacks[0])
	fmt.Println("stacks2", stacks[1])
	fmt.Println("stacks3", stacks[2])
	fmt.Println("stacks6", stacks[5])
	fmt.Println("stacks9", stacks[8])

	fmt.Println("Actions", actions[0])
	split_Action := strings.Split(actions[0], " ")
	fmt.Println(split_Action[1], split_Action[3], split_Action[5])

	for i, action := range actions {
		_ = i
		split_Action := strings.Split(action, " ")
		quantity := str_to_int(split_Action[1])
		departure_stack := str_to_int(split_Action[3]) - 1
		destination_stack := str_to_int(split_Action[5]) - 1
		// fmt.Println("Departure stack before", stacks[departure_stack])
		// fmt.Println("Destination stack before", stacks[destination_stack])

		// fmt.Println("Actions", quantity, departure_stack, destination_stack)

		for i := 1; i <= quantity; i++ {
			// fmt.Println(i)
			// fmt.Println(stacks[departure_stack])
			// fmt.Println(stacks[destination_stack])

			stacks[destination_stack] = stacks[destination_stack][:0] + stacks[departure_stack][0:1] + stacks[destination_stack][0:]
			stacks[departure_stack] = trimFirstRune(stacks[departure_stack])
		}

		// fmt.Println("Departure stack after", stacks[departure_stack])
		// fmt.Println("Destination stack after", stacks[destination_stack])

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
			// fmt.Println("Value", value, i)

		}
	}
	fmt.Println("Part 2")
	fmt.Println("Stacks", stacks)

	for i, action := range actions {
		_ = i
		split_Action = strings.Split(action, " ")
		quantity := str_to_int(split_Action[1])
		departure_stack := str_to_int(split_Action[3]) - 1
		destination_stack := str_to_int(split_Action[5]) - 1
		fmt.Println("Departure stack before", stacks[departure_stack])
		fmt.Println("Destination stack before", stacks[destination_stack])

		fmt.Println("Actions", quantity, departure_stack, destination_stack)

		stacks[destination_stack] = stacks[departure_stack][:quantity] + stacks[destination_stack]
		stacks[departure_stack] = stacks[departure_stack][quantity:]

		fmt.Println("Departure stack after", stacks[departure_stack])
		fmt.Println("Destination stack after", stacks[destination_stack])

	}
	fmt.Println("Part 2 Answer:")
	for i, row := range stacks {
		_ = row
		fmt.Println(stacks[i][:1])

	}

	// fmt.Println(stacks[0])
	// fmt.Println(stacks[0][:2])
	// fmt.Println(stacks[1])
	// fmt.Println(stacks[0][:2] + stacks[1])
	// fmt.Println(stacks[0][2:])

}

// for i, row := range crates {
// 	fmt.Println("row", i+1, row)
// 	split_row = SplitSubN(row, 4)
// 	stacks[0] = split_row[0] + split_row[1] + split_row[2] + split_row[3] + split_row[4] + split_row[5] + split_row[6] + split_row[7] + split_row[8]
// 	// for n, box := range stacks {
// 	// 	_ = box
// 	// 	stacks[n] = split_row

// 	// }
// }

/*Steps
  1. Split crates, stacks and actions
  2. Assign crates to rows
  3. Parse actions to indexes
  4. "Move the boxes in stacks"
  5. Get 1st box name in each stack
  6. Complete the message*/
