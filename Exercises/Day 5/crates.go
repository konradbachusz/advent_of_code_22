//https://adventofcode.com/2022/day/5

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	s := string(content)

	//List of pairs
	rows := strings.Split(strings.Replace(s, "\r\n", "\n", -1), "\n")
	crates := rows[:8]
	fmt.Println(len(crates[0]))
	fmt.Println(len(crates[2]))
	fmt.Println(len(crates[7]))

	fmt.Println(crates[0][12:15])

	actions := rows[10:len(rows)]
	fmt.Println(actions[0])

	/*Steps
	  1. Split crates, stacks and actions
	  2. Assign crates to rows
	  3. Parse actions to indexes
	  4. "Move the boxes in stacks"
	  5. Get 1st box name in each stack
	  6. Complete the message*/

}
