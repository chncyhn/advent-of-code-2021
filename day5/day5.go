package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const GRID_SIZE = 1_000

type Coord struct {
	x int
	y int
}

type Line struct {
	a Coord
	b Coord
}

func (l *Line) Points() []Coord {
	dx, dy := signum(l.b.x-l.a.x), signum(l.b.y-l.a.y)
	var coords []Coord
	for cx, cy := l.a.x, l.a.y; cx != l.b.x || cy != l.b.y; cx, cy = cx+dx, cy+dy {
		coords = append(coords, Coord{cx, cy})
	}
	coords = append(coords, Coord{l.b.x, l.b.y})
	return coords
}

func signum(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	} else {
		return 0
	}
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func parseCoord(txt string) Coord {
	a := strings.Split(txt, ",")
	ax, _ := strconv.Atoi(a[0])
	ay, _ := strconv.Atoi(a[1])
	return Coord{ax, ay}
}

func applyLines(grid [][]int, lines []Line, onlyHorizontal bool) {
	for _, line := range lines {
		horizontal := line.a.x == line.b.x || line.a.y == line.b.y
		if horizontal != onlyHorizontal {
			continue
		}
		for _, coord := range line.Points() {
			grid[coord.x][coord.y] += 1
		}
	}
}

func countSpots(grid [][]int) int {
	ans := 0
	for i := 0; i < GRID_SIZE; i++ {
		for j := 0; j < GRID_SIZE; j++ {
			if (grid)[i][j] > 1 {
				ans++
			}
		}
	}
	return ans
}

func prepareInput() ([]Line, [][]int) {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	var lines []Line
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), " -> ")
		lines = append(lines, Line{parseCoord(coords[0]), parseCoord(coords[1])})
	}
	grid := make([][]int, GRID_SIZE)
	for i := 0; i < GRID_SIZE; i++ {
		grid[i] = make([]int, GRID_SIZE)
	}
	return lines, grid
}

func main() {
	lines, grid := prepareInput()
	applyLines(grid, lines, true)
	fmt.Println("Part 1:", countSpots(grid))
	applyLines(grid, lines, false)
	fmt.Println("Part 2:", countSpots(grid))
}
