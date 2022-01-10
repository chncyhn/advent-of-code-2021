package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput() (grid [][]string) {
	file, _ := os.Open("data.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}
	return
}

func nextCoord(x, y, n, m int, kind string) (int, int) {
	if kind == ">" {
		return x, (y + 1) % m
	} else {
		return (x + 1) % n, y
	}
}

func move(grid [][]string, kind string) (anyMovement bool) {
	n, m := len(grid), len(grid[0])
	moveMark, deleteMark := kind+"|", kind+"~"
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if grid[x][y] == kind {
				nx, ny := nextCoord(x, y, n, m, kind)
				if grid[nx][ny] == "." {
					grid[nx][ny] = moveMark
					grid[x][y] = deleteMark
					anyMovement = true
				}
			}
		}
	}
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if grid[x][y] == deleteMark {
				grid[x][y] = "."
			} else if grid[x][y] == moveMark {
				grid[x][y] = kind
			}
		}
	}
	return
}

func step(grid [][]string) bool {
	right := move(grid, ">")
	down := move(grid, "v")
	return right || down
}

func main() {
	grid := readInput()
	for i := 1; ; i++ {
		if !step(grid) {
			fmt.Println("Answer:", i)
			break
		}
	}
}
