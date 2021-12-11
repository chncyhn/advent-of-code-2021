package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mapToInt(arr []string) []int {
	var numbers []int
	for _, v := range arr {
		intVal, _ := strconv.Atoi(v)
		numbers = append(numbers, intVal)
	}
	return numbers
}

func simulate(initialState [9]int, days int) int {
	state := initialState
	for i := 0; i < days; i++ {
		var frontier [9]int
		for s := 1; s < 9; s++ {
			frontier[s-1] += state[s]
		}
		frontier[6] += state[0]
		frontier[8] += state[0]
		state = frontier
	}
	tot := 0
	for s := 0; s < 9; s++ {
		tot += state[s]
	}
	return tot
}

func readInitialState() [9]int {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := mapToInt(strings.Split(scanner.Text(), ","))
	var states [9]int
	for _, s := range input {
		states[s] += 1
	}
	return states
}

func main() {
	initialState := readInitialState()
	fmt.Println("Part 1:", simulate(initialState, 80))
	fmt.Println("Part 2:", simulate(initialState, 256))
}
