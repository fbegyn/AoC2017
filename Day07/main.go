package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Program struct {
	name    string
	weight  uint
	subProg []string
	link    []*Program
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	stack := make([]*Program, 0)

	scanner := bufio.NewScanner(f)

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

		w := strings.Replace(phrase[1], "(", "", -1)
		w = strings.Replace(w, ")", "", -1)
		weight, err := strconv.Atoi(w)
		if err != nil {
			panic(err)
		}
		if len(phrase) > 3 {
			linker := make([]*Program, len(phrase[3:]))
			newP := &Program{phrase[0], uint(weight), phrase[3:], linker}
			stack = append(stack, newP)
		} else {
			newP := &Program{phrase[0], uint(weight), nil, nil}
			stack = append(stack, newP)
		}
	}
	findLinks(stack)
	lowest := findRoot(stack)
	fmt.Printf("The root program is %q with weight %d.\n\n", lowest.name, lowest.weight)
	outl, norm := findInBalance(stack)
	fmt.Printf("The sub-tower should weigh %d.\n", outl.weight-(outl.determineWeight()-norm.determineChildWeight()))
}

func findLinks(stack []*Program) {
	for _, prog := range stack {
		linker := make([]*Program, len(prog.subProg))
		for i, el := range prog.subProg {
			for _, p := range stack {
				if strings.Compare(el, p.name) == 0 {
					linker[i] = p
				}
			}
		}
		copy(prog.link[:], linker)
	}
}

func findRoot(stack []*Program) *Program {
	root := stack[0]
	for _, prog := range stack {
		if !prog.isChild(stack) {
			root = prog
		}
	}
	return root
}

func findInBalance(stack []*Program) (*Program, *Program) {
	root := findRoot(stack)
	var outl *Program
	var normal *Program
	var parentOutl *Program
	var parentNorm *Program
	outl, normal = root.getOutlier()
	for outl != nil {
		parentNorm = normal
		parentOutl = outl
		outl, normal = parentOutl.getOutlier()
	}
	return parentOutl, parentNorm
}

func (p *Program) getOutlier() (*Program, *Program) {
	freq := make(map[uint]uint)
	w := p.weightVector()
	var outlier uint
	var normal uint
	var outl *Program
	var norm *Program
	for _, el := range w {
		freq[el]++
	}
	for k, v := range freq {
		if v == 1 {
			outlier = k
		} else {
			normal = k
		}
	}
	for i, we := range w {
		if we == outlier {
			outl = p.link[i]
		}
		if we == normal {
			norm = p.link[i]
		}
	}
	return outl, norm
}

func (p *Program) isTop() bool {
	if len(p.link) <= 0 || len(p.subProg) <= 0 {
		return true
	}
	return false
}

func (p *Program) isChild(stack []*Program) bool {
	for _, prog := range stack {
		for _, link := range prog.link {
			if p == link {
				return true
			}
		}
	}
	return false
}

func (p *Program) weightVector() []uint {
	var w []uint
	for _, child := range p.link {
		w = append(w, child.determineWeight())
	}
	return w
}

func (p *Program) determineWeight() uint {
	var sum uint
	sum += p.weight
	for _, child := range p.link {
		sum += child.determineWeight()
	}
	return sum
}

func (p *Program) determineChildWeight() uint {
	var sum uint
	sum += p.weight
	for _, child := range p.link {
		sum += child.determineWeight()
	}
	return sum
}

func (p *Program) isParent(p2 *Program) bool {
	for _, child := range p.link {
		if child == p2 {
			return true
		}
	}
	return false
}
