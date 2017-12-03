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
	input := 13.0
	var side float64

	base := math.Ceil(math.Sqrt(input))
	if (int(base) % 2) == 0 {
		side = base + 1
	} else {
		side = base
	}
	stepsCenter := (side - 1) / 2                      // Steps from center to cycle
	offsetPrevCycle := input - (side-2)*(side-2)       // Offset from the start of the previous cycle
	offsetCenter := int(offsetPrevCycle) % int(side-1) // Which side are we on and which element
	distance := stepsCenter + math.Abs(float64(offsetCenter)-stepsCenter)

	fmt.Printf("The input has a Manhattan distance of %d \n", int(distance))
}

func prob2() {
	fmt.Printf("Too lazy, just looked the sequence up on OIES: http://oeis.org/A141481 \n")
}
