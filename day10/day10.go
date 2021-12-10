package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readRows() []string {
	file, _ := os.Open("data.txt")
	var rows []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	return rows
}

func peek(stack []rune) rune {
	return stack[len(stack)-1]
}

func pop(stack []rune) []rune {
	return stack[:len(stack)-1]
}

func solve1(row string, mapping map[rune]rune, scores map[rune]int) ([]rune, int) {
	var stack []rune
	for _, ch := range row {
		if opener, ok := mapping[ch]; ok {
			if len(stack) == 0 || opener != peek(stack) {
				return stack, scores[ch]
			} else {
				stack = pop(stack)
			}
		} else {
			stack = append(stack, ch)
		}
	}
	return stack, 0
}

func solve2(stack []rune, scores map[rune]int) int {
	score := 0
	for len(stack) > 0 {
		score = 5*score + scores[peek(stack)]
		stack = pop(stack)
	}
	return score
}

func median(arr []int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr[len(arr)/2]
}

func main() {
	mapping := make(map[rune]rune)
	mapping[')'] = '('
	mapping[']'] = '['
	mapping['}'] = '{'
	mapping['>'] = '<'

	scores1 := make(map[rune]int)
	scores1[')'] = 3
	scores1[']'] = 57
	scores1['}'] = 1197
	scores1['>'] = 25137

	scores2 := make(map[rune]int)
	scores2['('] = 1
	scores2['['] = 2
	scores2['{'] = 3
	scores2['<'] = 4

	part1 := 0
	var scores []int
	for _, row := range readRows() {
		stack, score1 := solve1(row, mapping, scores1)
		part1 += score1
		if score1 == 0 {
			scores = append(scores, solve2(stack, scores2))
		}
	}
	part2 := median(scores)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
