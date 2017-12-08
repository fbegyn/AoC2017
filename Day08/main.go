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

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, ",", "", -1)
		var phrase []string
		linescan := bufio.NewScanner(strings.NewReader(line))
		linescan.Split(bufio.ScanWords)

		for linescan.Scan() {
			word := linescan.Text()
			phrase = append(phrase, word)
		}

		reg := phrase[0]
		op := phrase[1]
		amount, err := strconv.Atoi(phrase[2])
		if err != nil {
			panic(err)
		}
		a := phrase[4]
		comp := phrase[5]
		b, err := strconv.Atoi(phrase[6])
		if err != nil {
			panic(err)
		}

		switch comp {
		case "<":
			if bank[a] < b {
				if op == "inc" {
					bank[reg] += amount
				} else if op == "dec" {
					bank[reg] -= amount
				}
			}
		case ">":
			if bank[a] > b {
				if op == "inc" {
					bank[reg] += amount
				} else if op == "dec" {
					bank[reg] -= amount
				}
			}
		case "<=":
			if bank[a] <= b {
				if op == "inc" {
					bank[reg] += amount
				} else if op == "dec" {
					bank[reg] -= amount
				}
			}
		case ">=":
			if bank[a] >= b {
				if op == "inc" {
					bank[reg] += amount
				} else if op == "dec" {
					bank[reg] -= amount
				}
			}
		case "==":
			if bank[a] == b {
				if op == "inc" {
					bank[reg] += amount
				} else if op == "dec" {
					bank[reg] -= amount
				}
			}
		case "!=":
			if bank[a] != b {
				if op == "inc" {
					bank[reg] += amount
				} else if op == "dec" {
					bank[reg] -= amount
				}
			}
		}
		for _, v := range bank {
			if highest < v {
				highest = v
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
