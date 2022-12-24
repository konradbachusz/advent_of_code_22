//https://adventofcode.com/2022/day/2

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
	_ = s
	// fmt.Println(s)

	folders_list := strings.Split(strings.Replace(s, "\r\n", "\n", -1), "\n")
	var dictionaries = map[string]map[string]map[string]string{}
	dictionaries["/"] = map[string]map[string]string{}

	//TODO uncomment
	folders_list = folders_list[:10]
	previous_command := ""
	line_type := ""
	for i, outputs := range folders_list {
		if i == 0 {
			previous_command = "not applicable"
		} else {
			previous_command = folders_list[i-1]
		}

		if strings.HasPrefix(outputs, "$") {
			line_type = "command"
		} else if strings.HasPrefix(outputs, "dir") {
			line_type = "folder"
			outputs = strings.Split(outputs, " ")[1]
			dictionaries["/"][outputs] = map[string]string{}
		} else {
			line_type = "file"

		}
		fmt.Print("\n", outputs, " ", previous_command, " ", line_type)
	}

	//var dictionaries = map[string]map[string]map[string]string{}

	fmt.Println("\n\ndictionaries: ", dictionaries)

	dictionaries["/"] = map[string]map[string]string{}
	// dictionaries["folder_1"] = make(map[string]map[string]string, 0)
	dictionaries["folder_2"] = map[string]map[string]string{}

	dictionaries["/"]["folder_A"] = map[string]string{}
	// dictionaries["folder_1"]["folder_A"] = make(map[string]string, 0)
	// dictionaries["folder_1"]["folder_A"] = make(map[string]string, 0)

	dictionaries["/"]["folder_A"]["folder_1"] = "white"
	// dictionaries["folder_1"]["folder_A"]["folder_2"] = "blue"
	// dictionaries["folder_1"]["folder_A"]["folder_3"] = "red"

	// dictionaries["folder_2"] = make(map[string]map[string]string)
	// dictionaries["folder_2"]["folder_A"] = make(map[string]string)
	// dictionaries["folder_2"]["folder_A"]["folder_5"] = "violet"

	fmt.Println("dictionaries: ", dictionaries)

	fmt.Println(dictionaries["folder_1"]["folder_A"]["folder_1"])

	my_dict := make(map[string]interface{})
	// my_dict["/"]["folder_A"] = map[string]string{}
	my_dict["/"] = make(map[string]interface{})
	my_dict["file1"] = 123212
	my_dict["file2"] = "hhyh"
	my_dict["/"]["file"] = 123

	fmt.Println(my_dict)
}

/*Steps
1. Identify all directories
2. Identify files corresponding to each dir
3. Get sizes of each dir
4. Get dirs <100000
5. Get total sum of their sizes*/
