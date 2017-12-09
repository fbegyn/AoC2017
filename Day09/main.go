package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	var value, score, countGarbage, garbage uint
	for i := 0; i < len(f); i++ {
		ch := f[i]
		if ch == '!' {
			i++
		} else {
			switch ch {
			case '{':
				if garbage == 0 {
					value++
				} else if garbage == 1 {
					countGarbage++
				}
			case '}':
				if garbage == 0 {
					score += value
					value--
				} else if garbage == 1 {
					countGarbage++
				}
			case '<':
				if garbage == 1 {
					countGarbage++
				}
				garbage = 1
			case '>':
				garbage = 0
			default:
				if garbage == 1 {
					countGarbage++
				}
			}
		}
	}
	fmt.Printf("The total score is %d and the amount of garbage char is %d\n", score, countGarbage)
}
