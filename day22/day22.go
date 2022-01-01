package main

import (
	"bufio"
	"fmt"
	"os"
)

type Range struct {
	lo, hi int
}

type Cuboid struct {
	x, y, z Range
}

type Flip struct {
	cuboid Cuboid
	diff   int
}

func (a *Cuboid) intersect(b Cuboid) Cuboid {
	return Cuboid{intersectRng(a.x, b.x), intersectRng(a.y, b.y), intersectRng(a.z, b.z)}
}

func (a *Cuboid) volume() int {
	return (1 + a.x.hi - a.x.lo) * (1 + a.y.hi - a.y.lo) * (1 + a.z.hi - a.z.lo)
}

func (a *Cuboid) real() bool {
	return (a.x.lo <= a.x.hi) && (a.y.lo <= a.y.hi) && (a.z.lo <= a.z.hi)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func intersectRng(a, b Range) Range {
	return Range{max(a.lo, b.lo), min(a.hi, b.hi)}
}

func readInput() (flips []Flip) {
	file, _ := os.Open("data.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var x1, x2, y1, y2, z1, z2 int
		var on string
		diff := 1
		fmt.Sscanf(scanner.Text(), "%s x=%d..%d,y=%d..%d,z=%d..%d", &on, &x1, &x2, &y1, &y2, &z1, &z2)
		if on != "on" {
			diff = -1
		}
		flips = append(flips, Flip{Cuboid{Range{x1, x2}, Range{y1, y2}, Range{z1, z2}}, diff})
	}
	return
}

func main() {
	state := make(map[Cuboid]int)
	for _, flip := range readInput() {
		frontier := make(map[Cuboid]int)
		frontier[flip.cuboid] = max(flip.diff, 0)
		for cub, cnt := range state {
			intersection := flip.cuboid.intersect(cub)
			if intersection.real() {
				frontier[intersection] -= cnt
			}
		}
		for cub, cnt := range frontier {
			state[cub] += cnt
		}
	}

	base := Cuboid{Range{-50, 50}, Range{-50, 50}, Range{-50, 50}}
	part1, part2 := 0, 0
	for cub, cnt := range state {
		subcub := base.intersect(cub)
		if subcub.real() {
			part1 += subcub.volume() * cnt
		}
		part2 += cub.volume() * cnt
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
