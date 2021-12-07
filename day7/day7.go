package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const UPPER_LIMIT = 2000

func mapToInt(arr []string) []int {
	var numbers []int
	for _, v := range arr {
		intVal, _ := strconv.Atoi(v)
		numbers = append(numbers, intVal)
	}
	return numbers
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func evaluate(arr []int, x int, evalFunc func(int) int) int {
	val := 0
	for _, v := range arr {
		val += evalFunc(abs(x - v))
	}
	return val
}

func bisect(arr []int, evalFunc func(int) int) int {
	lo, hi := 0, UPPER_LIMIT
	for lo < hi {
		mid := (lo + hi) / 2
		lv, hv := evaluate(arr, mid, evalFunc), evaluate(arr, mid+1, evalFunc)
		if lv < hv {
			hi = mid
		} else if lv > hv {
			lo = mid + 1
		}
	}
	return evaluate(arr, lo, evalFunc)
}

func identityDist(x int) int {
	return x
}

func increasingDist(x int) int {
	return x * (x + 1) / 2
}

func readInput() []int {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := mapToInt(strings.Split(scanner.Text(), ","))
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})
	return input
}

func main() {
	data := readInput()
	fmt.Println("Part 1:", bisect(data, identityDist))
	fmt.Println("Part 2:", bisect(data, increasingDist))
}
