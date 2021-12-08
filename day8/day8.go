package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func decode(signals []string) map[string]int {
	mapping := make(map[int]string)
	fivers := withLen(signals, 5)
	sixers := withLen(signals, 6)

	mapping[1] = withLen(signals, 2)[0]
	mapping[7] = withLen(signals, 3)[0]
	mapping[4] = withLen(signals, 4)[0]
	mapping[8] = withLen(signals, 7)[0]
	mapping[3] = withSubstr(fivers, mapping[1])
	mapping[9] = withSubstr(sixers, mapping[4])
	mapping[5] = whichSubstr(except(fivers, mapping[3]), mapping[9])
	mapping[2] = other(fivers, mapping[3], mapping[5])
	mapping[6] = withSubstr(except(sixers, mapping[9]), mapping[5])
	mapping[0] = other(sixers, mapping[6], mapping[9])

	reverseMapping := make(map[string]int)
	for i := 0; i < 10; i++ {
		reverseMapping[sortStr(mapping[i])] = i
	}
	return reverseMapping
}

func withLen(signals []string, size int) []string {
	var filtered []string
	for _, v := range signals {
		if len(v) != size {
			continue
		}
		filtered = append(filtered, v)
	}
	return filtered
}

func withSubstr(signals []string, substr string) string {
	for _, v := range signals {
		if contains(v, substr) {
			return v
		}
	}
	return ""
}

func except(signals []string, str string) []string {
	var filtered []string
	for _, v := range signals {
		if v == str {
			continue
		}
		filtered = append(filtered, v)
	}
	return filtered
}

func whichSubstr(signals []string, str string) string {
	for _, v := range signals {
		if contains(str, v) {
			return v
		}
	}
	return ""
}

func other(signals []string, str1, str2 string) string {
	for _, v := range signals {
		if v != str1 && v != str2 {
			return v
		}
	}
	return ""
}

func contains(s1, s2 string) bool {
	cnts1, cnts2 := make(map[rune]int), make(map[rune]int)
	for _, ch := range s1 {
		cnts1[ch]++
	}
	for _, ch := range s2 {
		cnts2[ch]++
	}
	for k, v := range cnts2 {
		if cnts1[k] < v {
			return false
		}
	}
	return true
}

func sortStr(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	part1, part2 := 0, 0
	for scanner.Scan() {
		inp := strings.Split(scanner.Text(), " | ")
		signals, output := strings.Split(inp[0], " "), strings.Split(inp[1], " ")
		mapping := decode(signals)
		value := 0
		for _, v := range output {
			if len(v) == 2 || len(v) == 3 || len(v) == 4 || len(v) == 7 {
				part1++
			}
			value = 10*value + mapping[sortStr(v)]
		}
		part2 += value
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
