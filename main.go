package main

import (
	"fmt"
	"os"
	"strconv"
)

func printCode(list []string) {
	var tmp int
	fmt.Println(list[0])
	fmt.Println(list[len(list)-1])
	for i := 0; i != len(list)-1; i++ {
		if list[i] == "before" {
			tmp, _ = strconv.Atoi(list[i+1])
		}
	}
	for i := 0; i != len(list)-1; i++ {
		if i == tmp-1 {
			fmt.Println(list[i])
		}
	}
	for i := 0; i != len(list)-1; i++ {
		if list[i] == "now" {
			tmp = i - 1
		}
	}
	for i, letter := range list[tmp] {
		if i == 1 {
			fmt.Println(list[int(letter)/len(list)-1])
		}
	}
}

func main() {
	var data []byte
	var list []string
	var tmp string
	data, _ = os.ReadFile("File.txt")
	for _, letter := range data {
		if letter == '\n' {
			list = append(list, tmp)
			tmp = ""
		} else if letter != '\r' {
			tmp += string(letter)
		}
	}
	list = append(list, tmp)
	printCode(list)
}
