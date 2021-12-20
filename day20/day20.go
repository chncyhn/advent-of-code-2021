package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	x, y int
}

const OFFSET = 60

func neis(p Pair) []Pair {
	return []Pair{
		{p.x - 1, p.y - 1},
		{p.x - 1, p.y},
		{p.x - 1, p.y + 1},

		{p.x, p.y - 1},
		{p.x, p.y},
		{p.x, p.y + 1},

		{p.x + 1, p.y - 1},
		{p.x + 1, p.y},
		{p.x + 1, p.y + 1},
	}
}

func step(algo []string, board map[Pair]bool, lenGrid, step int) map[Pair]bool {
	frontier := make(map[Pair]bool)
	for i := -OFFSET; i < lenGrid+OFFSET; i++ {
		for j := -OFFSET; j < lenGrid+OFFSET; j++ {
			cur := Pair{i, j}
			msk := 0
			for ix, nei := range neis(cur) {
				if flipped(Pair{nei.x, nei.y}, step, board) {
					msk = msk | (1 << (8 - ix))
				}
			}
			frontier[cur] = algo[msk] == "#"
		}
	}
	return frontier
}

func count(board map[Pair]bool, lenGrid, step int) (count int) {
	for i := -OFFSET; i < lenGrid+OFFSET; i++ {
		for j := -OFFSET; j < lenGrid+OFFSET; j++ {
			if flipped(Pair{i, j}, step, board) {
				count++
			}
		}
	}
	return
}

func flipped(loc Pair, step int, board map[Pair]bool) bool {
	if flipped, ok := board[loc]; flipped && ok {
		return true
	} else if !ok && step%2 == 0 {
		return true
	} else {
		return false
	}
}

func readInput() (algo []string, grid [][]string) {
	file, _ := os.Open("data.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	algo = strings.Split(scanner.Text(), "")
	scanner.Scan()
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}
	return
}

func main() {
	algo, grid := readInput()

	// Prepare board
	board := make(map[Pair]bool)
	for i := -OFFSET; i < len(grid)+OFFSET; i++ {
		for j := -OFFSET; j < len(grid[0])+OFFSET; j++ {
			if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == "." {
				board[Pair{i, j}] = false
			} else {
				board[Pair{i, j}] = true
			}
		}
	}

	// Solve
	n := len((grid))
	for stp := 1; stp <= 50; stp++ {
		board = step(algo, board, n, stp)
		if stp == 2 {
			fmt.Println("Part 1", count(board, n, stp))
		} else if stp == 50 {
			fmt.Println("Part 2", count(board, n, stp))
		}
	}
}
