package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func prob1() {
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

		min, max := MinMax(numbers)
		checksum += max - min
	}
	fmt.Printf("The checksum is %d \n", checksum)
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
