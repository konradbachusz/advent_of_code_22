//https://adventofcode.com/2022/day/2

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// func intersection(s1, s2 []string) (inter []string) {
// 	hash := make(map[string]bool)
// 	for _, e := range s1 {
// 		hash[e] = true
// 	}
// 	for _, e := range s2 {
// 		// If elements present in the hashmap then append intersection list.
// 		if hash[e] {
// 			inter = append(inter, e)
// 		}
// 	}
// 	return
// }

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func str_to_int(str_number string) int {
	number, err := strconv.Atoi(str_number)

	_ = err

	return number
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {

	///////////Part 1///////////////

	//Loading file
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	//Initialize count
	count := 0

	//Reading string
	s := string(content)

	//List of pairs
	pairs_list := strings.Split(strings.Replace(s, "\r\n", "\n", -1), "\n")

	for i, inventory := range pairs_list {
		_ = inventory

		//Split pair into separate into separate individuals
		split_pair := strings.Split(pairs_list[i], ",")

		elf1 := strings.Split(split_pair[0], "-")

		elf2 := strings.Split(split_pair[1], "-")

		slice1 := makeRange(str_to_int(elf1[0]), str_to_int(elf1[1]))
		slice2 := makeRange(str_to_int(elf2[0]), str_to_int(elf2[1]))

		//Check if slice in another slice
		if (contains(slice1, str_to_int(elf2[0])) && contains(slice1, str_to_int(elf2[1]))) || (contains(slice2, str_to_int(elf1[0])) && contains(slice2, str_to_int(elf1[1]))) {
			count = count + 1

		}

	}

	fmt.Println("Part 1 solution:", count)

}
