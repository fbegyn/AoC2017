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
	pipes := make(map[int][]int)
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var line string
	var tmp []string
	for scanner.Scan() {
		line = scanner.Text()
		tmp = strings.Split(line, " <-> ")
		key, err := strconv.Atoi(tmp[0])
		if err != nil {
			log.Fatal(err)
		}
		tmp[1] = strings.Replace(tmp[1], " ", "", -1)
		tmp2 := strings.Split(tmp[1], ",")
		var values []int
		for _, el := range tmp2 {
			value, err := strconv.Atoi(el)
			if err != nil {
				log.Fatal(err)
			}
			values = append(values, value)
		}
		pipes[key] = values
	}
	prob1(pipes)
}

func prob1(pipes map[int][]int) {
	zero, ok := getMap(&pipes, 0)
	fmt.Printf("Er zijn %d programmas met pipes naar 0.\n", len(*zero))
	var store []*map[int]bool
	store = append(store, zero)
	x := 1
	for len(pipes) > 0 {
		zero, ok = getMap(&pipes, x)
		if ok {
			store = append(store, zero)
		}
		x++
	}
	fmt.Printf("There are %d groups.\n", len(store))
}

func getMap(pipes *map[int][]int, key int) (*map[int]bool, bool) {
	set := make(map[int]bool)
	ok := setLinks(&set, pipes, key)
	for k, _ := range set {
		delete(*pipes, k)
	}
	return &set, ok
}

func setLinks(set *map[int]bool, pipes *map[int][]int, key int) bool {
	if _, ok := (*pipes)[key]; ok {
		(*set)[key] = true
		for _, el := range (*pipes)[key] {
			if !(*set)[el] {
				setLinks(set, pipes, el)
			}
		}
		return true
	} else {
		return false
	}
}
