package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BOARD_SIZE = 5

type Coord struct {
	x int
	y int
}

type Board struct {
	grid     [][]int
	location map[int]Coord
}

func mapToInt(arr []string) []int {
	var numbers []int
	for _, v := range arr {
		intVal, _ := strconv.Atoi(v)
		numbers = append(numbers, intVal)
	}
	return numbers
}

func readBoards(scanner *bufio.Scanner) []Board {
	var boards []Board
	for scanner.Scan() {
		scanner.Text()
		var rows [][]int
		for i := 0; i < BOARD_SIZE; i++ {
			scanner.Scan()
			numbers := mapToInt(strings.Fields(scanner.Text()))
			rows = append(rows, numbers)
		}
		boards = append(boards, createBoard(rows))
	}
	return boards
}

func createBoard(rows [][]int) Board {
	locations := make(map[int]Coord)
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			val := rows[i][j]
			locations[val] = Coord{i, j}
		}
	}
	return Board{rows, locations}
}

func apply(board Board, num int) bool {
	if coord, ok := board.location[num]; ok {
		board.grid[coord.x][coord.y] -= 10_000
		return true
	} else {
		return false
	}
}

func done(board Board, num int) bool {
	coord := board.location[num]
	rows := true
	cols := true
	for i := 0; i < BOARD_SIZE; i++ {
		rows = rows && board.grid[i][coord.y] < 0
		cols = cols && board.grid[coord.x][i] < 0
	}
	return rows || cols
}

func solve(boards []Board, numbers []int) {
	solved := make(map[int]bool)
	for _, num := range numbers {
		for b, board := range boards {
			if solved[b] {
				continue
			}
			hit := apply(board, num)
			if hit && done(board, num) {
				solved[b] = true
				if len(solved) == 1 {
					fmt.Println("Part 1: ", scorify(board, num))
				} else if len(solved) == len(boards) {
					fmt.Println("Part 2: ", scorify(board, num))
				}
			}
		}
	}
}

func scorify(board Board, last int) int {
	tot := 0
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if board.grid[i][j] > 0 {
				tot += board.grid[i][j]
			}
		}
	}
	return tot * last
}

func main() {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	numbers := mapToInt(strings.Split(scanner.Text(), ","))
	boards := readBoards(scanner)
	solve(boards, numbers)
}
