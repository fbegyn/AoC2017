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

	for scanner.Scan() {
		phrase = strings.Split(scanner.Text(), " ")

		reg := phrase[0]
		amount, err := strconv.Atoi(phrase[2])
		if err != nil {
			panic(err)
		}
		if phrase[1] == "dec" {
			amount *= -1
		}
		a := phrase[4]
		comp := phrase[5]
		b, err := strconv.Atoi(phrase[6])
		if err != nil {
			panic(err)
		}

		var w bool

		switch comp {
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
			bank[reg] += amount
		}

		if highest < bank[reg] {
			highest = bank[reg]
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
