package main

import "fmt"

// Could be done with container/ring, is a good one to remember

func main() {
	prob1()
	prob2()
}

func prob1() {
	spinlock := []uint64{0}
	var x uint64 = 1
	var step uint64 = 380
	for i := 1; i < 2018; i++ {
		spinlock = append(spinlock[:x], append([]uint64{uint64(i)}, spinlock[x:]...)...)
		if i == 2017 {
			fmt.Printf("The region around 2017 looks like this %v\n", spinlock[x-3:x+4])
		}
		x = (x + step + 1) % uint64(len(spinlock))
	}
}

func prob2() {
	steps := 380
	var found, pos int
	for i := 1; i < 50e6; i++ {
		pos = (pos + steps) % i
		if pos == 0 {
			found = i
		}
		pos++
	}
	fmt.Printf("The value after 0 is %v\n", found)
}
