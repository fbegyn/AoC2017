package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	programs := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'}
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var in string
	scanner := bufio.NewScanner(f)
	if scanner.Scan() {
		in = scanner.Text()
	}
	instructions := strings.Split(in, ",")
	solve(programs, instructions)
}

func solve(progs []byte, instructions []string) {
	var store []string
	repeat := true
	for repeat {
		store = append(store, string(progs))
		for _, el := range instructions {
			switch el[0] {
			case 's':
				oper := el[1:]
				step, err := strconv.Atoi(oper)
				if err != nil {
					log.Fatal(err)
				}
				rotate(progs, step)
			case 'x':
				oper := strings.Split(el[1:], "/")
				step1, err := strconv.Atoi(oper[0])
				step2, err := strconv.Atoi(oper[1])
				if err != nil {
					log.Fatal(err)
				}
				swap(progs, step1, step2)
			case 'p':
				partner(progs, el[1], el[3])
			}
		}
		for _, el := range store {
			if string(progs) == el {
				repeat = false
			}
		}
	}
	index := 1e6 % len(store)
	fmt.Printf("The first iteration gives us %s\nThe 1 billionth iteration gives us %s\n", store[1], store[index])
}

func rotate(inp []byte, x int) {
	cop := make([]byte, len(inp))
	copy(cop, inp)
	copy(cop[x:], inp[:len(inp)-x])
	copy(cop[:x], inp[len(inp)-x:])
	copy(inp, cop)

}

func swap(inp []byte, x, y int) {
	inp[x], inp[y] = inp[y], inp[x]
}

func partner(inp []byte, a, b byte) {
	var x, y int
	for i, el := range inp {
		if el == a {
			x = i
		}
		if el == b {
			y = i
		}
	}
	swap(inp, x, y)
}
