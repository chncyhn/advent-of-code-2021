package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve1(inp []string, bitLen int) {
	gamma, eps := 0, 0
	for i := 0; i < bitLen; i++ {
		mostCommon := mostCommonBitVal(inp, i)
		if mostCommon == 1 {
			gamma += 1 << (bitLen - 1 - i)
		} else {
			eps += 1 << (bitLen - 1 - i)
		}
	}
	fmt.Println(gamma * eps)
}

func solve2(inp []string, bitLen int) {
	oxy, _ := strconv.ParseInt(find(inp, bitLen, false), 2, 64)
	co2, _ := strconv.ParseInt(find(inp, bitLen, true), 2, 64)
	fmt.Println(oxy * co2)
}

func cntOnes(inp []string, bit int) int {
	cntOnes := 0
	for _, x := range inp {
		if x[bit] == '1' {
			cntOnes++
		}
	}
	return cntOnes
}

func mostCommonBitVal(inp []string, bit int) int {
	if float64(cntOnes(inp, bit)) >= (float64(len(inp)) / 2.0) {
		return 1
	} else {
		return 0
	}
}

func keys(m map[string]int) []string {
	v := make([]string, 0, len(m))
	for k, _ := range m {
		v = append(v, k)
	}
	return v
}

func find(inp []string, bitLen int, flip bool) string {
	vals := make(map[string]int)
	for _, x := range inp {
		vals[x] = 1
	}
	for i := 0; i < bitLen; i++ {
		remaining := keys(vals)
		mostCommon := mostCommonBitVal(remaining, i)
		for _, x := range remaining {
			remove := string(x[i]) != strconv.Itoa(mostCommon)
			if remove != flip {
				delete(vals, x)
			}
		}
		if len(vals) == 1 {
			for k, _ := range vals {
				return k
			}
		}
	}
	return ""
}

func main() {
	file, _ := os.Open("data.txt")
	var inp []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num := scanner.Text()
		inp = append(inp, num)
	}
	bitLen := len(inp[0])
	solve1(inp, bitLen)
	solve2(inp, bitLen)
}
