package main

import (
	"fmt"
)

type Node struct {
	x, y, xSc, ySc int
}

type Roll struct {
	val, cnt int
}

func part1(x, y int) int {
	xSc, ySc, die := 0, 0, 1
	for xSc < 1000 && ySc < 1000 {
		xRoll := 3*die + 3
		x = (x + xRoll) % 10
		xSc += (x + 1)
		die += 3
		if xSc >= 1000 {
			break
		}
		yRoll := 3*die + 3
		y = (y + yRoll) % 10
		ySc += (y + 1)
		die += 3
	}
	if xSc > ySc {
		return ySc * (die - 1)
	} else {
		return xSc * (die - 1)
	}
}

func expandX(frontier, queue map[Node]int, rolls []Roll, xWon *int) {
	for node, cnt := range queue {
		for _, roll := range rolls {
			x := (node.x + roll.val) % 10
			nw := Node{x, node.y, node.xSc + x + 1, node.ySc}
			if nw.xSc < 21 {
				frontier[nw] += cnt * roll.cnt
			} else {
				*xWon += cnt * roll.cnt
			}
		}
	}
}

func expandY(frontier, queue map[Node]int, rolls []Roll, yWon *int) {
	for node, cnt := range queue {
		for _, roll := range rolls {
			y := (node.y + roll.val) % 10
			nw := Node{node.x, y, node.xSc, node.ySc + y + 1}
			if nw.ySc < 21 {
				frontier[nw] += cnt * roll.cnt
			} else {
				*yWon += cnt * roll.cnt
			}
		}
	}
}

func part2(x, y int) int {
	rolls := []Roll{{3, 1}, {4, 3}, {5, 6}, {6, 7}, {7, 6}, {8, 3}, {9, 1}}
	xWon, yWon, queue := 0, 0, make(map[Node]int)
	queue[Node{x, y, 0, 0}] = 1
	for iter, xturn := 0, true; len(queue) > 0; iter, xturn = iter+1, !xturn {
		frontier := make(map[Node]int)
		if xturn {
			expandX(frontier, queue, rolls, &xWon)
		} else {
			expandY(frontier, queue, rolls, &yWon)
		}
		queue = frontier
	}
	if xWon > yWon {
		return xWon
	} else {
		return yWon
	}
}

func main() {
	x, y := 2, 3
	fmt.Println("Part 1", part1(x, y))
	fmt.Println("Part 2", part2(x, y))
}
