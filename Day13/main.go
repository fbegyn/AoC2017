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
	store := make(map[int]int)
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		var tmpInt [2]int
		for i, el := range strings.Split(line, ": ") {
			num, err := strconv.Atoi(el)
			if err != nil {
				log.Fatal(err)
			}
			tmpInt[i] = num
		}
		store[tmpInt[0]] = tmpInt[1]
	}
	fmt.Println(store)
	sev, _ := severity(store)
	del := delay(store)
	fmt.Printf("Severity: %v \nDelay: %d\n", sev, del)
}

func delay(store map[int]int) (delay int) {
scan:
	for {
		for depth, size := range store {
			if found(size, depth+delay) {
				delay++
				continue scan
			}
		}
		return
	}
}

func severity(store map[int]int) (sev int, caught bool) {
	for depth, size := range store {
		if found(size, depth) {
			caught = true
			sev += depth * size
		}
	}
	return
}

func found(size, delay int) bool {
	if size == 1 {
		return true
	}
	return delay%(2*(size-1)) == 0
}
