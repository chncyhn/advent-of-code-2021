package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Pair struct {
	x, y int
}

func isValid(x, y, n, m int) bool {
	return x >= 0 && x < n && y >= 0 && y < m
}

func neighbors(coord Pair, n, m int) []Pair {
	var neis []Pair
	offsets := []Pair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, o := range offsets {
		if isValid(coord.x+o.x, coord.y+o.y, n, m) {
			neis = append(neis, Pair{coord.x + o.x, coord.y + o.y})
		}
	}
	return neis
}

func readInput() [][]int {
	file, _ := os.Open("data.txt")
	var rows [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row []int
		for _, ch := range scanner.Text() {
			row = append(row, int(ch)-int('0'))
		}
		rows = append(rows, row)
	}
	return rows
}

func solveBasin(basin Pair, n, m int, rows [][]int, visited map[Pair]bool) int {
	basinSize := 1
	queue := []Pair{basin}
	for len(queue) > 0 {
		var frontier []Pair
		for _, coord := range queue {
			for _, neigh := range neighbors(coord, n, m) {
				if !visited[neigh] && rows[neigh.x][neigh.y] != 9 {
					visited[neigh] = true
					frontier = append(frontier, neigh)
					basinSize += 1
				}
			}
		}
		queue = frontier
	}
	return basinSize
}

func findBasins(n, m int, rows [][]int) []Pair {
	var basins []Pair
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			coord := Pair{i, j}
			basin := true
			for _, nei := range neighbors(coord, n, m) {
				if rows[i][j] >= rows[nei.x][nei.y] {
					basin = false
				}
			}
			if basin {
				basins = append(basins, coord)
			}
		}
	}
	return basins
}

func main() {
	rows := readInput()
	n, m := len(rows), len(rows[0])

	basins := findBasins(n, m, rows)
	visited := make(map[Pair]bool)
	part1 := 0
	for _, basin := range basins {
		part1 += 1 + rows[basin.x][basin.y]
		visited[basin] = true
	}

	var basinSizes []int
	for _, basin := range basins {
		basinSizes = append(basinSizes, solveBasin(basin, n, m, rows, visited))
	}
	sort.Slice(basinSizes, func(i, j int) bool {
		return basinSizes[i] > basinSizes[j]
	})
	part2 := basinSizes[0] * basinSizes[1] * basinSizes[2]

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
