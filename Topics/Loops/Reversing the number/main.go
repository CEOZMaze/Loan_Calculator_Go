package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := 0

	for input <= 0 {
		fmt.Scan(&input)
	}
	inputStr := strconv.Itoa(input)

	revertStr := ""
	for i := len(inputStr) - 1; i >= 0; i-- {
		revertStr += string(inputStr[i])
	}
	reversedNum, err := strconv.Atoi(revertStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(reversedNum)
}
