package main

import "fmt"

func main() {
	ch1, ch2 := make(chan uint, 64), make(chan uint, 64)
	const facA uint = 16807
	const facB uint = 48271
	const div uint = 2147483647
	var judge, judge2 uint
	//var size int
	var genA uint = 512
	var genB uint = 191
	var resA, resB, passA, passB []uint
	for i := 0; len(passA) < 5e6 || len(passB) < 5e6; i++ {
		resA = append(resA, (genA*facA)%div)
		resB = append(resB, (genB*facB)%div)
		genA = resA[i]
		genB = resB[i]
		if resA[i]%4 == 0 {
			passA = append(passA, resA[i])
		}
		if resB[i]%8 == 0 {
			passB = append(passB, resB[i])
		}
	}
	go func() {
		for i := 0; i < 40e6; i++ {
			a := resA[i] & 0xffff
			b := resB[i] & 0xffff
			if a == b {
				judge++
			}
		}
		ch1 <- judge
		close(ch1)
	}()
	go func() {
		for i := 0; i < 5e6; i++ {
			a := passA[i] & 0xffff
			b := passB[i] & 0xffff
			if a == b {
				judge2++
			}
		}
		ch2 <- judge2
		close(ch2)
	}()
	fmt.Printf("There were %d matches.\nThere were %d strict matches\n", <-ch1, <-ch2)
}
