package main

import (
	"container/heap"
	"fmt"
	"sort"
)

const AMPH_COUNT = 16

var COSTS = [4]int{1, 10, 100, 1000}
var PERMS = precomputePerms()

type State [AMPH_COUNT]int
type Node struct {
	locs State
	dist int
}
type Coord struct {
	x, y int
}

// based on https://pkg.go.dev/container/heap#example-package-IntHeap
type Heap []Node

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(Node)) }
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func cost(a int) int {
	return COSTS[a/4]
}

func coordify(a int) Coord {
	if a <= 10 {
		return Coord{0, a}
	} else {
		lt := (a - 11) / 4
		y := 2 + 2*lt
		x := 1 + a - (11 + 4*lt)
		return Coord{x, y}
	}
}

func dist(a, b int) int {
	ac, bc := coordify(a), coordify(b)
	if ac.y != bc.y {
		return abs(ac.y-bc.y) + ac.x + bc.x
	} else {
		return abs(ac.x - bc.x)
	}
}

func normalize(st State) {
	for s := 0; s < 16; s += 4 {
		sort.Slice(st[s:(s+4)], func(i, j int) bool { return st[s+i] < st[s+j] })
	}
}

func isFull(loc int, occuppied map[int]int) bool {
	_, ok := occuppied[loc]
	return ok
}

func addNeiToRoom(neis *[]Node, st State, occuppied map[int]int, a int) {
	loc, at, targetY := st[a], a/4, 2*(a/4+1)
	if loc > targetY {
		for y := loc - 1; y >= targetY; y-- {
			if isFull(y, occuppied) {
				return
			}
		}
	} else {
		for y := loc + 1; y <= targetY; y++ {
			if isFull(y, occuppied) {
				return
			}
		}
	}
	for y := 11 + 4*at; y < 15+4*at; y++ {
		if bt, ok := occuppied[y]; ok && bt != at {
			return
		}
	}
	entryLoc, best := 11+4*at, -1
	for y := entryLoc; y < entryLoc+4; y++ {
		if isFull(y, occuppied) {
			break
		}
		best = y
	}
	if best != -1 {
		*neis = append(*neis, Node{apply(st, a, best), cost(a) * dist(loc, best)})
	}
}

func isEntrypoint(l int) bool {
	return l >= 2 && l <= 8 && l%2 == 0
}

func addNeisToHall(neis *[]Node, st State, occuppied map[int]int, a int) {
	loc := st[a]
	curCol := (loc - 11) / 4
	curY := (curCol + 1) * 2
	for r := loc - 1; r >= 11+curCol*4; r-- {
		if isFull(r, occuppied) {
			return
		}
	}
	if isFull(curY, occuppied) {
		return
	}
	for y := curY; y <= 10; y++ {
		if isFull(y, occuppied) {
			break
		}
		if isEntrypoint(y) {
			continue
		}
		*neis = append(*neis, Node{apply(st, a, y), cost(a) * dist(loc, y)})
	}
	for y := curY - 1; y >= 0; y-- {
		if isFull(y, occuppied) {
			break
		}
		if isEntrypoint(y) {
			continue
		}
		*neis = append(*neis, Node{apply(st, a, y), cost(a) * dist(loc, y)})
	}
}

func neis(st State) (nxt []Node) {
	occuppied := make(map[int]int)
	for a := 0; a < AMPH_COUNT; a++ {
		occuppied[st[a]] = a / 4
	}
	for at := 0; at < 4; at++ {
		for a := 4 * at; a <= 4*at+3; a++ {
			if st[a] <= 10 {
				addNeiToRoom(&nxt, st, occuppied, a)
			} else {
				addNeisToHall(&nxt, st, occuppied, a)
			}
		}
	}
	return
}

func apply(st State, amph, newLoc int) State {
	nxtState := st
	nxtState[amph] = newLoc
	normalize(nxtState)
	return nxtState
}

// source: https://stackoverflow.com/a/30226442
func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}
	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func heur(st State) int {
	ans := 0
	for at := 0; at < 4; at++ {
		distRes := 100000000
		for _, amphs := range PERMS[at] {
			curDist := 0
			for i := 0; i < 4; i++ {
				curDist += dist(4*at+11+i, st[amphs[i]]) * COSTS[at]
			}
			distRes = min(distRes, curDist)
		}
		ans += distRes
	}
	return ans
}

func astar(start State) int {
	g := make(map[State]int)
	f := make(map[State]int)
	processed := make(map[State]bool)

	pq := &Heap{Node{start, heur(start)}}
	heap.Init(pq)
	g[start] = 0
	f[start] = heur(start)
	for {
		node := heap.Pop(pq).(Node)
		if heur(node.locs) == 0 {
			return node.dist
		}
		if _, ok := processed[node.locs]; ok {
			continue
		}
		processed[node.locs] = true
		for _, nei := range neis(node.locs) {
			if _, ok := processed[nei.locs]; ok {
				continue
			}
			cand := g[node.locs] + nei.dist
			neiDist := 100000000
			if val, ok := g[nei.locs]; ok {
				neiDist = val
			}
			if cand < neiDist {
				g[nei.locs] = cand
				f[nei.locs] = cand + heur(nei.locs)
				heap.Push(pq, Node{nei.locs, f[nei.locs]})
			}
		}
	}
}

func precomputePerms() map[int][][]int {
	perms := make(map[int][][]int)
	perms[0] = permutations([]int{0, 1, 2, 3})
	perms[1] = permutations([]int{4, 5, 6, 7})
	perms[2] = permutations([]int{8, 9, 10, 11})
	perms[3] = permutations([]int{12, 13, 14, 15})
	return perms
}

func main() {
	/*
	 #############
	 #0123456789X#
	 ###1#5#9#3###
	   #2#6#X#4#
	   #3#7#1#5#
	   #4#8#2#6#
	   #########
	*/
	start := State{
		19, 21, 23, 24,
		17, 20, 22, 26,
		14, 16, 18, 25,
		11, 12, 13, 15,
	}
	fmt.Println(astar(start))
}
