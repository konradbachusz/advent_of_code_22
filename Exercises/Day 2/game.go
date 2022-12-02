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

	rock := [3]interface{}{"A", "X", 1}
	paper := [3]interface{}{"B", "Y", 2}
	scissors := [3]interface{}{"C", "Z", 3}

	fmt.Println(rock[2], paper, scissors)

	outcomes := map[interface{}]interface{}{
		"lost": 0,
		"draw": 3,
		"won":  6,
	}
	fmt.Println(outcomes)
	score := 0

	rock_val := rock[2].(int)
	paper_val := paper[2].(int)
	scissors_val := scissors[2].(int)

	draw_val := outcomes["draw"].(int)
	lost_val := outcomes["lost"].(int)
	won_val := outcomes["won"].(int)

	s := string(content)

	games_list := strings.Split(strings.Replace(s, "\r\n", "\n", -1), "\n")

	fmt.Println("Part One")

	for i, s := range games_list {
		if strings.HasPrefix(s, "A") && strings.HasSuffix(s, "X") {
			score = score + rock_val + draw_val

		} else if strings.HasPrefix(s, "A") && strings.HasSuffix(s, "Y") {
			score = score + paper_val + won_val
		} else if strings.HasPrefix(s, "A") && strings.HasSuffix(s, "Z") {
			score = score + scissors_val + lost_val
		} else if strings.HasPrefix(s, "B") && strings.HasSuffix(s, "X") {
			score = score + rock_val + lost_val
		} else if strings.HasPrefix(s, "B") && strings.HasSuffix(s, "Y") {
			score = score + paper_val + draw_val
		} else if strings.HasPrefix(s, "B") && strings.HasSuffix(s, "Z") {
			score = score + scissors_val + won_val
		} else if strings.HasPrefix(s, "C") && strings.HasSuffix(s, "X") {
			score = score + rock_val + won_val
		} else if strings.HasPrefix(s, "C") && strings.HasSuffix(s, "Y") {
			score = score + paper_val + lost_val
		} else if strings.HasPrefix(s, "C") && strings.HasSuffix(s, "Z") {
			score = score + scissors_val + draw_val
		}

		fmt.Println(i, s, score)
	}

	fmt.Println("Part Two")
	score = 0
	for i, s := range games_list {
		if strings.HasPrefix(s, "A") && strings.HasSuffix(s, "X") {
			score = score + scissors_val + lost_val
		} else if strings.HasPrefix(s, "A") && strings.HasSuffix(s, "Y") {
			score = score + rock_val + draw_val
		} else if strings.HasPrefix(s, "A") && strings.HasSuffix(s, "Z") {
			score = score + paper_val + won_val
		} else if strings.HasPrefix(s, "B") && strings.HasSuffix(s, "X") {
			score = score + rock_val + lost_val
		} else if strings.HasPrefix(s, "B") && strings.HasSuffix(s, "Y") {
			score = score + paper_val + draw_val
		} else if strings.HasPrefix(s, "B") && strings.HasSuffix(s, "Z") {
			score = score + scissors_val + won_val
		} else if strings.HasPrefix(s, "C") && strings.HasSuffix(s, "X") {
			score = score + paper_val + lost_val
		} else if strings.HasPrefix(s, "C") && strings.HasSuffix(s, "Y") {
			score = score + scissors_val + draw_val
		} else if strings.HasPrefix(s, "C") && strings.HasSuffix(s, "Z") {
			score = score + rock_val + won_val
		}

		fmt.Println(i, s, score)
	}
}
