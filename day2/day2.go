package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	cmd string
	d   int
}

func solve1(mvs []Move) {
	x, y := 0, 0
	for _, mv := range mvs {
		switch mv.cmd {
		case "forward":
			x += mv.d
		case "down":
			y += mv.d
		case "up":
			y -= mv.d
		}
	}
	fmt.Println("Part 1: ", x*y)
}

func solve2(mvs []Move) {
	x, y, aim := 0, 0, 0
	for _, mv := range mvs {
		switch mv.cmd {
		case "forward":
			x += mv.d
			y += aim * mv.d
		case "down":
			aim += mv.d
		case "up":
			aim -= mv.d
		}
	}
	fmt.Println("Part 2: ", x*y)
}

func main() {
	file, err := os.Open("data2.txt")
	if err != nil {
	}
	scanner := bufio.NewScanner(file)
	var mvs []Move
	for scanner.Scan() {
		cmd := strings.Split(scanner.Text(), " ")
		d, err := strconv.Atoi(cmd[1])
		if err != nil {
			return
		}
		mvs = append(mvs, Move{cmd[0], d})
	}
	solve1(mvs)
	solve2(mvs)
}
