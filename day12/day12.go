package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func solve(s, x string, graph map[string][]string, cnt map[string]int, paths map[string]bool, path string) {
	if s == "end" {
		paths[path] = true
		return
	}
	for _, nei := range graph[s] {
		if isLower(nei) && ((x == nei && cnt[nei] >= 2) || (x != nei && cnt[nei] >= 1)) {
			continue
		}
		cnt[nei]++
		solve(nei, x, graph, cnt, paths, path+","+nei)
		cnt[nei]--
	}
}

func addEdge(graph map[string][]string, u, v string) {
	if edges, ok := graph[u]; ok {
		graph[u] = append(edges, v)
	} else {
		graph[u] = []string{v}
	}
}

func readGraph() map[string][]string {
	file, _ := os.Open("data.txt")
	defer file.Close()
	graph := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		uv := strings.Split(scanner.Text(), "-")
		u, v := uv[0], uv[1]
		addEdge(graph, u, v)
		addEdge(graph, v, u)
	}
	return graph
}

func main() {
	graph := readGraph()

	// part 1
	cnts1 := make(map[string]int)
	cnts1["start"] = 1
	paths1 := make(map[string]bool)
	solve("start", "", graph, cnts1, paths1, "start")

	// part 2
	paths2 := make(map[string]bool)
	for k := range graph {
		if isLower(k) {
			cnts := make(map[string]int)
			cnts["start"] = 2
			solve("start", k, graph, cnts, paths2, "start")
		}
	}

	fmt.Println("Part 1:", len(paths1))
	fmt.Println("Part 2:", len(paths2))
}
