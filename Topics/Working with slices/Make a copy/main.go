package main

import (
	"fmt"
	"log"
	"reflect"
)

func solve(s []string) []string {
	// Write the code to make a copy of the 's' slice below:
	sCopy := make([]string, len(s))
	copy(sCopy, s)

	return sCopy // Return the copy of the slice here!
}

// DO NOT delete or modify the contents of the main function!
func main() {
	mySlice := []string{"", "", ""}
	fmt.Scan(&mySlice[0], &mySlice[1], &mySlice[2])

	newSlice := solve(mySlice)
	if !reflect.DeepEqual(mySlice, newSlice) {
		log.Fatal("new slice is not an exact copy")
	}

	mySlice[0], mySlice[1] = mySlice[1], mySlice[0]
	if reflect.DeepEqual(mySlice, newSlice) {
		log.Fatal("new slice is a reference to an old slice, but should be a copy")
	}

	fmt.Println(newSlice)
}
