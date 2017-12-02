package main

import (
	"io/ioutil"
	"strconv"
)

func prob1() {
	dat, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	var inp string = string(dat)
	var sum int64
	var el2 rune
	var el rune

	size := len(inp) - 1
	offset := 1

	for i, _ := range inp {
		el = rune(inp[i%size])
		el2 = rune(inp[(i+offset)%size])

		if el == el2 {
			number, e := strconv.Atoi(string(el))
			if e != nil {
				panic(e)
			}
			sum += int64(number)
			println(sum)
			println("****************")
		}
	}
}
