package main

import (
	"bufio"
	"fmt"
	"os"
)

const FLASH_THRESHOLD = 10

type Pair struct {
	x, y int
}

func readBoard() [][]int {
	file, _ := os.Open("data.txt")
	var rows [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, mapToInt(scanner.Text()))
	}
	return rows
}

func mapToInt(str string) []int {
	var numbers []int
	for _, v := range str {
		numbers = append(numbers, int(v)-int('0'))
	}
	return numbers
}

func offsets() []Pair {
	var offs []Pair
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				offs = append(offs, Pair{i, j})
			}
		}
	}
	return offs
}

func valid(i, j, n, m int) bool {
	return i >= 0 && i < n && j >= 0 && j < m
}

func neis(i, j, n, m int) []Pair {
	var neis []Pair
	for _, o := range offsets() {
		x, y := i+o.x, j+o.y
		if valid(x, y, n, m) {
			neis = append(neis, Pair{x, y})
		}
	}
	return neis
}

func flash(board [][]int, coord Pair) {
	n, m := len(board), len(board[0])
	for _, nei := range neis(coord.x, coord.y, n, m) {
		if board[nei.x][nei.y] < FLASH_THRESHOLD {
			board[nei.x][nei.y]++
			if board[nei.x][nei.y] == FLASH_THRESHOLD {
				flash(board, nei)
			}
		}
	}
}

func tick(board [][]int) int {
	n, m := len(board), len(board[0])
	// increment
	var flashers []Pair
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			board[i][j]++
			if board[i][j] == FLASH_THRESHOLD {
				flashers = append(flashers, Pair{i, j})
			}
		}
	}
	// flash
	for _, coord := range flashers {
		flash(board, coord)
	}
	// reset
	flashCnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == FLASH_THRESHOLD {
				flashCnt++
				board[i][j] = 0
			}
		}
	}
	return flashCnt
}

func main() {
	board := readBoard()
	part1, part2 := 0, 0
	for i := 1; ; i++ {
		flashed := tick(board)
		if i <= 100 {
			part1 += flashed
		}
		if flashed == 100 {
			part2 = i
			break
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
