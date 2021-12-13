package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	x, y int
}

type Fold struct {
	loc int
	dir string
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func pprint(board [][]int) {
	n, m := len(board), len(board[0])
	for c := 0; c < m; c++ {
		for r := 0; r < n; r++ {
			if board[r][c] > 0 {
				fmt.Print("X ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
	}
}

func add(a, b []int) []int {
	var c []int
	for i := 0; i < len(a); i++ {
		c = append(c, a[i]+b[i])
	}
	return c
}

func setRow(board [][]int, r int, row []int) {
	board[r] = row[:]
}

func setCol(board [][]int, c int, col []int) {
	r := len(board)
	for i := 0; i < r; i++ {
		board[i][c] = col[i]
	}
}

func getCol(board [][]int, c int) []int {
	r := len(board)
	var col []int
	for i := 0; i < r; i++ {
		col = append(col, board[i][c])
	}
	return col
}

func snipRow(board [][]int, r int) [][]int {
	return board[:r]
}

func snipCol(board [][]int, c int) [][]int {
	r := len(board)
	var snipped [][]int
	for i := 0; i < r; i++ {
		var row []int
		for j := 0; j < c; j++ {
			row = append(row, board[i][j])
		}
		snipped = append(snipped, row)
	}
	return snipped
}

func count(board [][]int) int {
	n, m, cnt := len(board), len(board[0]), 0
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if board[r][c] > 0 {
				cnt++
			}
		}
	}
	return cnt
}

func fold(fold Fold, board [][]int) [][]int {
	n, m := len(board), len(board[0])
	f := fold.loc
	folded := board
	if fold.dir == "x" {
		for o := 1; f+o < n; o++ {
			setRow(folded, f-o, add(board[f-o], board[f+o]))
		}
		return snipRow(folded, f)
	} else {
		for o := 1; f+o < m; o++ {
			setCol(folded, f-o, add(getCol(board, f-o), getCol(board, f+o)))
		}
		return snipCol(folded, f)
	}
}

func parsePair(line string) Pair {
	coords := strings.Split(line, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	return Pair{x, y}
}

func parseFold(line string) Fold {
	words := strings.Split(line, " ")
	fold := strings.Split(words[len(words)-1], "=")
	a, _ := strconv.Atoi(fold[1])
	return Fold{a, fold[0]}
}

func prepareBoard(spots []Pair) [][]int {
	n, m := 0, 0
	for _, spot := range spots {
		n = max(n, spot.x+1)
		m = max(m, spot.y+1)
	}
	var board [][]int
	for r := 0; r < n; r++ {
		var row []int
		for c := 0; c < m; c++ {
			row = append(row, 0)
		}
		board = append(board, row)
	}
	for _, spot := range spots {
		board[spot.x][spot.y] = 1
	}
	return board
}

func readInput() ([][]int, []Fold) {
	file, _ := os.Open("data.txt")
	defer file.Close()
	var spots []Pair
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		spots = append(spots, parsePair(line))
	}
	var folds []Fold
	for scanner.Scan() {
		folds = append(folds, parseFold(scanner.Text()))
	}
	return prepareBoard(spots), folds
}

func main() {
	board, folds := readInput()
	for i, f := range folds {
		board = fold(f, board)
		if i == 0 {
			fmt.Println("Part 1", count(board))
		}
	}
	pprint(board)
}
