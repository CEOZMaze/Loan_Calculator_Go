package main

import "fmt"

func main() {
	words := []string{"saltwater", "backstage", "bedrock", "sandcastle", "snowman"}

	w1 := words[0][4:]
	w2 := words[1][4:]
	w3 := words[2][3:]
	w4 := words[3][4:]
	w5 := words[4][4:]

	fmt.Println(w1, w2, w3, w4, w5) // DO NOT delete this line, it prints the substrings!
}
