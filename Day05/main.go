package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	prob1()
	prob2()
}

func prob1() {
	var instructions []int

	src, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	scanner := bufio.NewScanner(src)

	for scanner.Scan() {
		instr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, instr)
	}

	fmt.Printf("The operation took %d steps.\n", execInstr(&instructions, false))
}

func execInstr(instructions *[]int, prob2 bool) uint {
	var index int
	var op int
	var ops uint
	for index < len(*instructions) {
		op = (*instructions)[index]
		if op >= 3 && prob2 {
			(*instructions)[index]--
		} else {
			(*instructions)[index]++
		}
		index = index + op
		ops++
	}
	return ops
}

func prob2() {
	var instructions []int

	src, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	scanner := bufio.NewScanner(src)

	for scanner.Scan() {
		instr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, instr)
	}

	fmt.Printf("The operation took %d steps.\n", execInstr(&instructions, true))
}
