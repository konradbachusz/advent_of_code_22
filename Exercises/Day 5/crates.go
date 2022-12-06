//https://adventofcode.com/2022/day/5

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

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
			// fmt.Println("Value", value, i)
			stacks[n] = stacks[n] + value
		}
	}
	// stacks[0] = crates[0]

	// for n, row := range crates {
	// 	_ = n
	// 	split_row = SplitSubN(row, 4)
	// 	stacks[i] = split_row[0] + split_row[1] + split_row[2] + split_row[3] + split_row[4] + split_row[5] + split_row[6] + split_row[7] + split_row[8]
	// }

	fmt.Println("stacks0", stacks[0])
	fmt.Println("stacks1", stacks[1])
	fmt.Println("stacks2", stacks[2])
	fmt.Println("stacks5", stacks[5])
	fmt.Println("stacks8", stacks[8])

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
