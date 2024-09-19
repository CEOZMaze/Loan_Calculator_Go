package main

import (
	"fmt"
	"math"
)

func main() {
	var radians float64
	fmt.Scanln(&radians)

	// Please use any of the above identities to calculate the 'cotangent' below:
	cotangent := math.Cos(radians) / math.Sin(radians)

	fmt.Printf("%.2f", cotangent)
}
