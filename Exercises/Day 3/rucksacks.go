//https://adventofcode.com/2022/day/2

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

	//Items string
	var items = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println(len(items))

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

	fmt.Println(sum_of_items)

	fmt.Println(rucksack_list[0])

	string_len := len(rucksack_list[0])
	half_len := string_len / 2
	fmt.Println(string_len)

	first_compartment := rucksack_list[0][:half_len]
	fmt.Println(first_compartment)

	second_compartment := rucksack_list[0][half_len:string_len]
	fmt.Println(second_compartment)

	//Intersection
	first_compartment_list := strings.Split(first_compartment, "")

	fmt.Println(first_compartment_list)

	second_compartment_list := strings.Split(second_compartment, "")

	fmt.Println(second_compartment_list)

	result := intersection(first_compartment_list, second_compartment_list) // or intersection(second, first)
	fmt.Println(result)

	var value = strings.Index(items, result[0]) + 1
	fmt.Println("The value of the item is: ", value)

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
	fmt.Println(sum_of_items)

}