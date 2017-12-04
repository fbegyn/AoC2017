package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	prob1(368078.0)
	prob2(368078)
	fmt.Printf("Time elapsed for program %s\n", time.Since(start))
}

func prob1(input float64) {
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

func prob2(input uint64) {
	fmt.Printf("Too lazy, just looked the sequence up on OIES: http://oeis.org/A141481 \n")

	spiral := [11][11]uint64{}
	mask := make([][3]uint64, 3)
	i := 11 / 2
	j := 11 / 2

	spiral[i][j] = 1

	lLimit := 1
	uLimit := 1
	rLimit := 1
	dLimit := 1
	direction := 0
	larger := false

	for lLimit < 11 || uLimit < 11 || rLimit < 11 || dLimit < 11 {
		if direction == 0 {
			for l := 0; l < lLimit; l++ {
				j++
				mask = *applyMask(&spiral, i, j)
				sum := calcSum(&mask)
				larger = checkSum(input, sum, &spiral)
				spiral[i][j] = sum
			}
			lLimit++
			rLimit++
			direction = 1
		} else if direction == 1 {
			for u := 0; u < uLimit; u++ {
				i--
				mask = *applyMask(&spiral, i, j)
				sum := calcSum(&mask)
				larger = checkSum(input, sum, &spiral)
				spiral[i][j] = sum
			}
			uLimit++
			dLimit++
			direction = 2
		} else if direction == 2 {
			for r := 0; r < rLimit; r++ {
				j--
				mask = *applyMask(&spiral, i, j)
				sum := calcSum(&mask)
				larger = checkSum(input, sum, &spiral)
				spiral[i][j] = sum
			}
			rLimit++
			lLimit++
			direction = 3
		} else if direction == 3 {
			for d := 0; d < dLimit; d++ {
				i++
				mask = *applyMask(&spiral, i, j)
				sum := calcSum(&mask)
				larger = checkSum(input, sum, &spiral)
				spiral[i][j] = sum
			}
			dLimit++
			uLimit++
			direction = 0
		}
		if larger {
			break
		}
	}
}

func applyMask(spiral *[11][11]uint64, i int, j int) *[][3]uint64 {
	mask := make([][3]uint64, 3)
	for true {
		if i-1 < 0 || i+1 > 10 || j-1 < 0 || j+1 > 10 {
			break
		}
		x := 0
		for ind, el := range *spiral {
			if ind < i-1 || ind > i+1 {
				continue
			}
			copy(mask[x][:3], el[j-1:j+2])
			x++
		}
		break
	}
	return &mask
}

func calcSum(mask *[][3]uint64) uint64 {
	var sum uint64
	for _, el := range *mask {
		for _, el2 := range el {
			sum += el2
		}
	}
	return sum
}

func checkSum(input uint64, sum uint64, spiral *[11][11]uint64) bool {
	if sum > input {
		fmt.Printf("Got bored and did it racefully. The first number that is larger then the input is %d\n", sum)
		return true
	}
	return false
}
