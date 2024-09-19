package main

import "fmt"

func main() {
	// DO NOT delete or modify the code block below, it reads a random input word
	var word string
	fmt.Scan(&word)

	// Use slice expressions with the len() function
	// to remove all the letters from 'word' except for the last letter.

	fmt.Println(string(word[len(word)-1]))
}
