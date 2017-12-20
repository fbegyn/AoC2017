package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*Particle describes the properties of a particle*/
type Particle struct {
	position     [3]int
	velocity     [3]int
	accellartion [3]int
	exists       bool
}

func (p *Particle) calcDist() int {
	return abs(p.position[0]) + abs(p.position[1]) + abs(p.position[2])
}

func (p *Particle) collide(p2 *Particle) {
	x := p.position[0] == p2.position[0]
	y := p.position[1] == p2.position[1]
	z := p.position[2] == p2.position[2]
	if x && y && z {
		p.exists = false
		p2.exists = false
	}
}

func (p *Particle) checkColl(parts []*Particle) {
	for _, el := range parts {
		p.collide(el)
	}
}

func (p *Particle) update() {
	p.velocity[0] += p.accellartion[0]
	p.velocity[1] += p.accellartion[1]
	p.velocity[2] += p.accellartion[2]
	p.position[0] += p.velocity[0]
	p.position[1] += p.velocity[1]
	p.position[2] += p.velocity[2]
}

func tick(parts *[]*Particle) (index, dist int) {
	dist = 1e6
	var d int
	for i, el := range *parts {
		el.update()
		d = el.calcDist()
		if dist > d {
			dist = d
			index = i
		}
	}
	for i, el := range *parts {
		el.checkColl((*parts)[i+1:])
	}
	return
}

func checkState(parts *[]*Particle) (count int) {
	for _, el := range *parts {
		if el.exists {
			count++
		}
	}
	return
}

func simulate(parts *[]*Particle, ch chan int) {
	var ind, count int
	for {
		count = checkState(parts)
		ind, _ = tick(parts)
		ch <- ind
		ch <- count
	}
}

func main() {
	// Input section
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Data construction section
	var particles []*Particle
	var p, v, a [3]int
	for _, line := range lines {
		fmt.Sscanf(line, "p=<%d,%d,%d>,  v=<%d,%d,%d>, a=<%d,%d,%d> \n", &p[0], &p[1], &p[2], &v[0], &v[1], &v[2], &a[0], &a[1], &a[2])

		part := &Particle{p, v, a, true}
		particles = append(particles, part)

	}

	// Challenge section
	ch1 := make(chan int, 64)
	go simulate(&particles, ch1)
	for {
		fmt.Printf("\r The particle closes to (0,0,0) is %d. There are %d particles at this moment.", <-ch1, <-ch1)
	}
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
