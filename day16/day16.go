package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func hexVal(rn rune) int {
	if rn >= '0' && rn <= '9' {
		return int(rn) - int('0')
	} else {
		return int(rn) - int('A') + 10
	}
}

func toBinary(hex string) []int {
	var bin []int
	for _, h := range hex {
		hx := hexVal(h)
		for i := 3; i >= 0; i-- {
			if hx&(1<<i) > 0 {
				bin = append(bin, 1)
			} else {
				bin = append(bin, 0)
			}
		}
	}
	return bin
}

func read(stream []int, loc, len int) (int, int) {
	read := 0
	for i := 0; i < len; i++ {
		read *= 2
		read += stream[loc+i]
	}
	return loc + len, read
}

func reduce(op int, vals []int) int {
	totVal := 0
	if op == 0 {
		for _, v := range vals {
			totVal += v
		}
	} else if op == 1 {
		totVal = 1
		for _, v := range vals {
			totVal *= v
		}
	} else if op == 2 {
		totVal = math.MaxInt
		for _, v := range vals {
			if totVal > v {
				totVal = v
			}
		}
	} else if op == 3 {
		for _, v := range vals {
			if totVal < v {
				totVal = v
			}
		}
	} else if op == 5 && vals[0] > vals[1] {
		totVal = 1
	} else if op == 6 && vals[0] < vals[1] {
		totVal = 1
	} else if op == 7 && vals[0] == vals[1] {
		totVal = 1
	}
	return totVal
}

func readLiteral(stream []int, loc int) (int, int) {
	var val, totVal int
	cur, hasMore := loc, 1
	for hasMore != 0 {
		cur, hasMore = read(stream, cur, 1)
		cur, val = read(stream, cur, 4)
		totVal <<= 4
		totVal += val
	}
	return cur, totVal
}

func readOp(stream []int, loc, typ int) (int, int, int) {
	var lenType, subLen, cnt, subVer, subVal, totVer int
	var subVals []int
	cur := loc
	cur, lenType = read(stream, cur, 1)
	if lenType == 0 {
		cur, subLen = read(stream, cur, 15)
		start := cur
		for cur-start < subLen {
			cur, subVer, subVal = readPacket(stream, cur)
			totVer += subVer
			subVals = append(subVals, subVal)
		}
	} else {
		cur, cnt = read(stream, cur, 11)
		for i := 0; i < cnt; i++ {
			cur, subVer, subVal = readPacket(stream, cur)
			totVer += subVer
			subVals = append(subVals, subVal)
		}
	}
	return cur, totVer, reduce(typ, subVals)
}

func readPacket(stream []int, loc int) (int, int, int) {
	var typ, ver, val int
	cur := loc
	cur, ver = read(stream, cur, 3)
	cur, typ = read(stream, cur, 3)
	if typ == 4 { // literal
		cur, val = readLiteral(stream, cur)
		return cur, ver, val
	} else { // op
		var opVer int
		cur, opVer, val = readOp(stream, cur, typ)
		return cur, ver + opVer, val
	}
}

func main() {
	file, _ := os.Open("data.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	_, totVer, totVal := readPacket(toBinary(scanner.Text()), 0)
	fmt.Println("Part 1", totVer)
	fmt.Println("Part 2", totVal)
}
