package main

import (
	"fmt"
	"math"
)

func main() {
	prob1()
	prob2()
}
func prob1() {
	input := 368078
	var n int

	base := math.Sqrt(368078)
	if (int(math.Ceil(base)) % 2) == 0 {
		n = int(base)
	} else {
		n = int(math.Ceil(base))
	}
	distance := n*n - input

	fmt.Printf("The input has a Manhattan distance of %d\n", distance)
}

func prob2() {
	fmt.Printf("Too lazy, just looked the sequence up on OIES: http://oeis.org/A141481 \n")
}
