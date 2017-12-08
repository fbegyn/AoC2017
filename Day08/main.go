package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	bank := make(map[string]int)
	var highest int
	var phrase []string
	var reg, a string

	for scanner.Scan() {
		phrase = strings.Split(scanner.Text(), " ")

		reg = phrase[0]
		a = phrase[4]
		b, err := strconv.Atoi(phrase[6])
		if err != nil {
			panic(err)
		}

		var w bool

		switch phrase[5] {
		case "<":
			w = bank[a] < b
		case ">":
			w = bank[a] > b
		case "<=":
			w = bank[a] <= b
		case ">=":
			w = bank[a] >= b
		case "==":
			w = bank[a] == b
		case "!=":
			w = bank[a] != b
		}

		if w {
			amount, err := strconv.Atoi(phrase[2])
			if err != nil {
				panic(err)
			}
			if phrase[1] == "dec" {
				amount *= -1
			}

			bank[reg] += amount

			if highest < bank[reg] {
				highest = bank[reg]
			}
		}

	}

	var largest int
	for _, v := range bank {
		if largest < v {
			largest = v
		}
	}

	fmt.Printf("The largest register is %d.\n", largest)
	fmt.Printf("The largest register during process is %d.\n", highest)
}
