package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var value, score, countGarbage uint
	garbage := false
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		//line := "<!!!>>"
		for i := 0; i < len(line); i++ {
			ch := line[i]
			switch ch {
			case '{':
				if !garbage {
					value++
				} else if garbage {
					countGarbage++
				}
			case '}':
				if !garbage {
					score += value
					value--
				} else if garbage {
					countGarbage++
				}
			case '<':
				if garbage {
					countGarbage++
				}
				garbage = true
			case '>':
				garbage = false
			case '!':
				i++
			default:
				if garbage {
					countGarbage++
				}
			}
		}
	}
	fmt.Printf("The total score is %d\n", score)
	fmt.Printf("The total amount of garbage characters is %d\n", countGarbage)
}
