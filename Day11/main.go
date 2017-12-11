package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var line string
	if scanner.Scan() {
		line = scanner.Text()
	}
	steps := strings.Split(line, ",")
	tracker := make(map[string]int)
	var distance, maxDist int
	for _, step := range steps {
		tracker[step] += 1
		vertical := tracker["n"] - tracker["s"]
		diagonalLeft := tracker["nw"] - tracker["se"]
		diagonalRight := tracker["ne"] - tracker["sw"]
		a, b := vertical+diagonalRight, diagonalLeft+diagonalRight
		if a < 0 {
			a *= -1
		}
		if b < 0 {
			b *= -1
		}
		if a > b {
			distance = a
		} else {
			distance = b
		}
		if maxDist < distance {
			maxDist = distance
		}
	}
	fmt.Printf("He is %d tiles away and was at a max distance of %d\n", distance, maxDist)
}
