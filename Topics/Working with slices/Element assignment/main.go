package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

func solve(s []string, x string, i int) []string {
	// Write the code to make a copy of the `s` slice below:

	sliceCopy := make([]string, len(s))
	copy(sliceCopy, s)

	// Assign `x` to the element at the `i` index of the `sliceCopy` slice:
	sliceCopy[i] = x
	return sliceCopy // Return `sliceCopy` here!
}

// DO NOT delete or modify the contents of the main function!
func main() {
	var i int
	var x string

	mySlice := make([]string, 1)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	mySlice[i], x = scanner.Text(), scanner.Text()

	newSlice := solve(mySlice, x, i)
	if !reflect.DeepEqual(mySlice, newSlice) {
		log.Fatal("You didn't assign `x` to the element at the `i` index of the slice!")
	}
	fmt.Println(newSlice)
}
