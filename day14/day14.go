package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func score(cnts map[string]int) int {
	min, max := -1, -1
	for _, v := range cnts {
		if min == -1 || v < min {
			min = v
		}
		if max == -1 || v > max {
			max = v
		}
	}
	return max - min
}

func main() {
	file, _ := os.Open("data.txt")
	defer file.Close()
	// prepare counts & pairs
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	poly := strings.Split(scanner.Text(), "")
	pairs := make(map[string]int)
	cnts := make(map[string]int)
	for i := 0; i < len(poly)-1; i++ {
		pairs[poly[i]+poly[i+1]]++
		cnts[poly[i]]++
	}
	cnts[poly[len(poly)-1]]++
	// prepare templates
	templates := make(map[string]string)
	scanner.Scan()
	for scanner.Scan() {
		inp := strings.Split(scanner.Text(), " ")
		templates[inp[0]] = inp[2]
	}
	// solve
	cur := pairs
	for step := 1; step <= 40; step++ {
		frontier := make(map[string]int)
		for pair, cnt := range cur {
			if res, ok := templates[pair]; ok {
				p := strings.Split(pair, "")
				cnts[res] += cnt
				frontier[p[0]+res] += cnt
				frontier[res+p[1]] += cnt
			} else {
				frontier[pair] += cnt
			}
		}
		cur = frontier
		if step == 10 {
			fmt.Println("Part 1:", score(cnts))
		}
	}
	fmt.Println("Part 2:", score(cnts))
}
