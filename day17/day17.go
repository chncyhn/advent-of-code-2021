package main

import (
	"fmt"
)

type Pair struct {
	a, b int
}

func simX(xd int, xrange Pair) []int {
	var times []int
	for t, loc := 0, 0; loc <= xrange.b && xd >= 0; t++ {
		if xrange.a <= loc && loc <= xrange.b {
			times = append(times, t)
			if xd == 0 {
				for nt := t + 1; nt < 1000; nt++ {
					times = append(times, nt)
				}
			}
		}
		loc, xd = loc+xd, xd-1
	}
	return times
}

func simY(yd int, yrange Pair) []int {
	var times []int
	for t, loc := 0, 0; !((loc < yrange.a) && (loc < 0)); t++ {
		if yrange.a <= loc && loc <= yrange.b {
			times = append(times, t)
		}
		loc, yd = loc+yd, yd-1
	}
	return times
}

func main() {
	xrange, yrange := Pair{20, 30}, Pair{-10, -5}
	xPossibleAt := make(map[int][]int)
	for xd := 1; xd < 250; xd++ {
		for _, t := range simX(xd, xrange) {
			if xds, ok := xPossibleAt[t]; ok {
				xPossibleAt[t] = append(xds, xd)
			} else {
				xPossibleAt[t] = []int{xd}
			}
		}
	}

	maxYd := 0
	possibleVels := make(map[Pair]bool)
	for yd := -250; yd < 250; yd++ {
		for _, t := range simY(yd, yrange) {
			if _, ok := xPossibleAt[t]; !ok {
				continue
			}
			maxYd = yd
			for _, xd := range xPossibleAt[t] {
				possibleVels[Pair{xd, yd}] = true
			}
		}
	}
	fmt.Println("Part 1:", maxYd*(maxYd+1)/2)
	fmt.Println("Part 2:", len(possibleVels))
}
