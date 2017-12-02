package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var checksum int

	src, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	scanner := bufio.NewScanner(src)

	for scanner.Scan() {
		line := scanner.Text()
		var numbers []int

		lineScan := bufio.NewScanner(strings.NewReader(line))
		lineScan.Split(bufio.ScanWords)

		for lineScan.Scan() {
			number, e := strconv.Atoi(lineScan.Text())
			if e != nil {
				panic(e)
			}
			numbers = append(numbers, number)
		}

		res, even := dividesEvenly(numbers)
		if even {
			checksum += res
		}
	}
	fmt.Printf("The checksum is %d \n", checksum)
}

func dividesEvenly(array []int) (int, bool) {
	var op1 int
	var op2 int
	for index, el := range array {
		op1 = el
		for _, el2 := range array[index+1:] {
			op2 = el2
			if (op1 % op2) == 0 {
				return (op1 / op2), true
			}
			if (op2 % op1) == 0 {
				return (op2 / op1), true
			}
		}
	}
	return 0, false
}
