package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	x, y int
}

type Node struct {
	pair Pair
	dist int
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

func mapToInt(arr []string) []int {
	var numbers []int
	for _, v := range arr {
		intVal, _ := strconv.Atoi(v)
		numbers = append(numbers, intVal)
	}
	return numbers
}

func neis(p Pair, n int) []Pair {
	var neis []Pair
	x, y := p.x, p.y
	for _, d := range []Pair{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		if x+d.x >= 0 && x+d.x < n && y+d.y >= 0 && y+d.y < n {
			neis = append(neis, Pair{x + d.x, y + d.y})
		}
	}
	return neis
}

func risk(p Pair, board [][]int) int {
	n := len(board)
	val := board[p.x%n][p.y%n] + (p.y / n) + (p.x / n)
	if val >= 10 {
		return val - 9
	}
	return val
}

func dijkstra(N int, board [][]int) int {
	distTo := make(map[Pair]int)
	processed := make(map[Pair]bool)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			pair := Pair{i, j}
			distTo[pair] = 10000000
		}
	}
	h := &Heap{Node{Pair{0, 0}, 0}}
	heap.Init(h)
	distTo[Pair{0, 0}] = 0
	for len(processed) < N*N {
		node := heap.Pop(h).(Node)
		if _, ok := processed[node.pair]; ok {
			continue
		}
		processed[node.pair] = true
		for _, nei := range neis(node.pair, N) {
			cand := node.dist + risk(nei, board)
			if cand < distTo[nei] {
				distTo[nei] = cand
				heap.Push(h, Node{nei, cand})
			}
		}
	}
	return distTo[Pair{N - 1, N - 1}]
}

func main() {
	file, _ := os.Open("data.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var board [][]int
	for scanner.Scan() {
		board = append(board, mapToInt(strings.Split(scanner.Text(), "")))
	}
	n := len(board)
	fmt.Println("Part 1:", dijkstra(n, board))
	fmt.Println("Part 2:", dijkstra(5*n, board))

}
