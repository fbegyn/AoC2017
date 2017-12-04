package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type runes []rune

func (r runes) Len() int           { return len(r) }
func (r runes) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r runes) Less(i, j int) bool { return r[i] < r[j] }

func main() {
	prob1()
	prob2()
}

func prob1() {
	var validPass int

	src, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	scanner := bufio.NewScanner(src)

	for scanner.Scan() {
		line := scanner.Text()
		var phrase []string

		lineScan := bufio.NewScanner(strings.NewReader(line))
		lineScan.Split(bufio.ScanWords)

		for lineScan.Scan() {
			word := lineScan.Text()
			phrase = append(phrase, word)
		}

		double, _ := checkPassphrase(&phrase)
		if !double {
			validPass++
		}

	}

	fmt.Printf("There are %d valid passwords.\n", validPass)
}

func checkPassphrase(phrase *[]string) (bool, bool) {
	containsDouble := false
	containsAnagrams := false
	for ind, el := range *phrase {
		for _, el2 := range (*phrase)[ind+1:] {
			if strings.Compare(el2, el) == 0 {
				containsDouble = true
			}
			if checkAnagram(el2, el) {
				containsAnagrams = true
			}
		}
	}
	return containsDouble, containsAnagrams
}

func string2Rune(s string) []rune {
	var runes []rune
	for _, r := range s {
		runes = append(runes, r)
	}
	return runes
}

func checkAnagram(s1, s2 string) bool {
	var r1 runes = string2Rune(s1)
	var r2 runes = string2Rune(s2)

	sort.Sort(r1)
	sort.Sort(r2)

	return string(r1) == string(r2)
}

func prob2() {
	var validPass int

	src, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	scanner := bufio.NewScanner(src)

	for scanner.Scan() {
		line := scanner.Text()
		var phrase []string

		lineScan := bufio.NewScanner(strings.NewReader(line))
		lineScan.Split(bufio.ScanWords)

		for lineScan.Scan() {
			word := lineScan.Text()
			phrase = append(phrase, word)
		}

		double, anagram := checkPassphrase(&phrase)
		if !double && !anagram {
			validPass++
		}

	}

	fmt.Printf("There are %d valid passwords with expanded security policy.\n", validPass)
}
