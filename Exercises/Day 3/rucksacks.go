//https://adventofcode.com/2022/day/3

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func intersection(s1, s2 []string) (inter []string) {
	hash := make(map[string]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		// If elements present in the hashmap then append intersection list.
		if hash[e] {
			inter = append(inter, e)
		}
	}
	return
}

func main() {

	///////////Part 1///////////////

	//Items string
	var items = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	//Loading file
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	//Reading string
	s := string(content)

	//List of rucksacs
	rucksack_list := strings.Split(strings.Replace(s, "\r\n", "\n", -1), "\n")

	//Initialize initial sum of items
	sum_of_items := 0

	for i, inventory := range rucksack_list {
		_ = i

		string_len := len(inventory)
		half_len := string_len / 2

		first_compartment := inventory[:half_len]
		second_compartment := inventory[half_len:string_len]

		//Intersection
		first_compartment_list := strings.Split(first_compartment, "")
		second_compartment_list := strings.Split(second_compartment, "")

		result := intersection(first_compartment_list, second_compartment_list)
		result_value := strings.Index(items, result[0]) + 1

		sum_of_items = result_value + sum_of_items

	}
	fmt.Println("Part 1 solution:", sum_of_items)

	///////////Part 2///////////////

	sum_of_items = 0

	size := 3
	var j int
	for i := 0; i < len(rucksack_list); i += size {
		j += size
		if j > len(rucksack_list) {
			j = len(rucksack_list)
		}

		first_rucksack_list := strings.Split(rucksack_list[i:j][0], "")
		second_rucksack_list := strings.Split(rucksack_list[i:j][1], "")
		third_compartment_list := strings.Split(rucksack_list[i:j][2], "")

		result := intersection(first_rucksack_list, second_rucksack_list)
		result = intersection(result, third_compartment_list)
		result_value := strings.Index(items, result[0]) + 1

		sum_of_items = result_value + sum_of_items
	}
	fmt.Println("Part 2 solution:", sum_of_items)

}
