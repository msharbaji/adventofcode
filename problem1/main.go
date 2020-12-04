package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	data, err := ioutil.ReadFile("inputs/problem1.txt")
	checker := make([]bool, 2021)
	if err != nil {
		fmt.Println(err)
	}
	input := strings.Split(string(data), "\n")
	for i := 0; i < len(input); i++ {
		number1, _ := strconv.Atoi(input[i])

		checker[number1] = true

		complement := 2020 - number1

		if checker[complement] {
			fmt.Println(number1 * complement)
			break
		}
	}
}
