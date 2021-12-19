package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Triplet struct {
	x, y, z int
}

type Pair struct {
	x, y int
}

type Orientation struct {
	facing Triplet // ±1s
	order  Triplet // 012 permutation
}

type Scanner struct {
	loc Triplet
	or  Orientation
}

func mapToInt(arr []string) []int {
	var numbers []int
	for _, v := range arr {
		intVal, _ := strconv.Atoi(v)
		numbers = append(numbers, intVal)
	}
	return numbers
}

func mapBeacons(sc Scanner, readings []Triplet) (beacons []Triplet) {
	for _, r := range readings {
		rord := reorder(sc.or, r)
		beacons = append(beacons, Triplet{sc.loc.x + rord.x, sc.loc.y + rord.y, sc.loc.z + rord.z})
	}
	return beacons
}

func reorder(or Orientation, t Triplet) Triplet {
	var rord [3]int
	rord[or.order.x] = t.x * or.facing.x
	rord[or.order.y] = t.y * or.facing.y
	rord[or.order.z] = t.z * or.facing.z
	return Triplet{rord[0], rord[1], rord[2]}
}

func countOverlap(b1, b2 []Triplet) (cnt int) {
	b1Has := make(map[Triplet]bool)
	for _, t1 := range b1 {
		b1Has[t1] = true
	}
	for _, t2 := range b2 {
		if b1Has[t2] {
			cnt++
		}
	}
	return cnt
}

func dir(shift, mask int) int {
	if mask&(1<<shift) > 0 {
		return -1
	}
	return 1
}

func orientations() (pos []Orientation) {
	for faceMask := 0; faceMask < 8; faceMask++ {
		xDir, yDir, zDir := dir(0, faceMask), dir(1, faceMask), dir(2, faceMask)
		pos = append(pos, Orientation{Triplet{xDir, yDir, zDir}, Triplet{0, 1, 2}})
		pos = append(pos, Orientation{Triplet{xDir, yDir, zDir}, Triplet{0, 2, 1}})
		pos = append(pos, Orientation{Triplet{xDir, yDir, zDir}, Triplet{1, 0, 2}})
		pos = append(pos, Orientation{Triplet{xDir, yDir, zDir}, Triplet{1, 2, 0}})
		pos = append(pos, Orientation{Triplet{xDir, yDir, zDir}, Triplet{2, 0, 1}})
		pos = append(pos, Orientation{Triplet{xDir, yDir, zDir}, Triplet{2, 1, 0}})
	}
	return pos
}

func mapOther(sc Scanner, orOther Orientation, readOwn, readOther Triplet) (locOther Triplet) {
	t1 := reorder(sc.or, readOwn)
	t2 := reorder(orOther, readOther)
	return Triplet{
		sc.loc.x + t1.x - t2.x,
		sc.loc.y + t1.y - t2.y,
		sc.loc.z + t1.z - t2.z,
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func manhattan(loc1, loc2 Triplet) int {
	tot := 0
	tot += abs(loc1.x - loc2.x)
	tot += abs(loc1.y - loc2.y)
	tot += abs(loc1.z - loc2.z)
	return tot
}

func locateScanner(
	scanners map[int]Scanner,
	readings [][]Triplet,
	ors []Orientation,
	visited map[Pair]bool,
) {
	for knownI, knownSc := range scanners {
		for otherI := 1; otherI < len(readings); otherI++ {
			if _, ok := scanners[otherI]; ok {
				continue
			}
			if _, ok := visited[Pair{knownI, otherI}]; ok {
				continue
			}
			visited[Pair{knownI, otherI}] = true
			for _, or := range ors {
				for _, knownR := range readings[knownI] {
					for _, otherR := range readings[otherI] {
						locOther := mapOther(knownSc, or, knownR, otherR)
						candSc := Scanner{locOther, or}
						cnt := countOverlap(
							mapBeacons(knownSc, readings[knownI]),
							mapBeacons(candSc, readings[otherI]),
						)
						if cnt >= 12 {
							fmt.Println("Matched", knownI, "to", otherI)
							scanners[otherI] = candSc
							return
						}
					}
				}
			}
		}
	}
	panic("Failed to locate!")
}

func readInput() (readings [][]Triplet) {
	file, _ := os.Open("data.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.Contains(txt, "scanner") {
			readings = append(readings, []Triplet{})
			continue
		} else if len(txt) != 0 {
			nums := mapToInt(strings.Split(txt, ","))
			tri := Triplet{nums[0], nums[1], nums[2]}
			readings[len(readings)-1] = append(readings[len(readings)-1], tri)
		}
	}
	return readings
}

func main() {
	readings := readInput()
	ors := orientations()
	// Locate all scanners
	visited := make(map[Pair]bool)
	scanners := make(map[int]Scanner)
	scanners[0] = Scanner{Triplet{0, 0, 0}, Orientation{Triplet{1, 1, 1}, Triplet{0, 1, 2}}}
	for len(scanners) < len(readings) {
		locateScanner(scanners, readings, ors, visited)
	}
	// Part 1
	beacons := make(map[Triplet]bool)
	for scI, sc := range scanners {
		for _, bc := range mapBeacons(sc, readings[scI]) {
			beacons[bc] = true
		}
	}
	// Part 2
	maxDist := 0
	for _, sc1 := range scanners {
		for _, sc2 := range scanners {
			maxDist = max(maxDist, manhattan(sc1.loc, sc2.loc))
		}
	}
	fmt.Println("Part 1:", len(beacons))
	fmt.Println("Part 2:", maxDist)
}
