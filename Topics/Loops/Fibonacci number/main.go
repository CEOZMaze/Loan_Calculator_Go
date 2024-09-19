package main

import (
	"fmt"
	"math"
)

func main() {
	// DO NOT delete this code block. It takes as an input the `n` value to be used in the for loop
	var n float64
	fmt.Scanln(&n)

	// DO NOT delete the variable declarations that will be used within the for loop!

	var lastDigit int

	// Write the additional required code to finish the for loop declaration below:

	lastDigit = int(math.Round((math.Pow(math.Phi, n)-math.Pow(1-math.Phi, n))/math.Sqrt(5))) % 10

	fmt.Println(lastDigit) // This line outputs the calculated `lastDigit`
}
