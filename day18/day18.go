package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	left, right *Node
	val         int
}

type NodeDepth struct {
	node  *Node
	depth int
}

func numeric(x string) bool {
	_, err := strconv.Atoi(x)
	return err == nil
}

func add(left, right *Node) *Node {
	sum := &Node{left, right, -1}
	for {
		if explode(sum) {
			continue
		}
		if !split(sum, nil) {
			break
		}
	}
	return sum
}

func isRegular(node *Node) bool {
	return node.val != -1
}

func split(node *Node, parent *Node) bool {
	if node == nil {
		return false
	}
	if isRegular(node) && node.val >= 10 {
		left := int(node.val / 2)
		right := int(math.Ceil(float64(node.val) / 2.0))
		newNode := &Node{&Node{nil, nil, left}, &Node{nil, nil, right}, -1}
		replaceChild(parent, node, newNode)
		return true
	}
	return split(node.left, node) || split(node.right, node)
}

func replaceChild(parent, oldChild, newChild *Node) {
	if parent.left == oldChild {
		parent.left = newChild
	} else if parent.right == oldChild {
		parent.right = newChild
	} else {
		panic("Did not find child!")
	}
}

func explode(rootNode *Node) bool {
	var last *Node
	var queue []NodeDepth
	queue = append(queue, NodeDepth{rootNode, 0})
	carry := -1
	parent := make(map[*Node]*Node)
	for len(queue) > 0 {
		q := queue[len(queue)-1]
		queue = queue[:(len(queue) - 1)]
		if carry == -1 && q.depth == 4 && !isRegular(q.node) {
			left := q.node.left.val
			right := q.node.right.val
			if last != nil {
				last.val += left
			}
			carry = right
			replaceChild(parent[q.node], q.node, &Node{nil, nil, 0})
			continue
		}
		if isRegular(q.node) {
			last = q.node
			if carry >= 0 {
				last.val += carry
				carry = -2
			}
		}
		if q.node.right != nil {
			parent[q.node.right] = q.node
			queue = append(queue, NodeDepth{q.node.right, q.depth + 1})
		}
		if q.node.left != nil {
			parent[q.node.left] = q.node
			queue = append(queue, NodeDepth{q.node.left, q.depth + 1})
		}
	}
	return carry != -1
}

func setChild(node, regChild *Node) {
	if node.left == nil {
		node.left = regChild
	} else {
		node.right = regChild
	}
}

func pparse(exp []string) *Node {
	var stack [](*Node)
	root := Node{nil, nil, -1}
	stack = append(stack, &root)
	for i := 1; i < len(exp); i++ {
		if exp[i] == "[" {
			node := Node{nil, nil, -1}
			setChild(stack[len(stack)-1], &node)
			stack = append(stack, &node)
		} else if exp[i] == "]" {
			stack = stack[:(len(stack) - 1)]
		} else if numeric(exp[i]) {
			v, _ := strconv.Atoi(exp[i])
			setChild(stack[len(stack)-1], &Node{nil, nil, v})
		}
	}
	return &root
}

func magnitude(node *Node) int {
	if isRegular(node) {
		return node.val
	}
	return 3*magnitude(node.left) + 2*magnitude(node.right)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	file, _ := os.Open("data.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var nodes [][]string
	for scanner.Scan() {
		nodes = append(nodes, strings.Split(scanner.Text(), ""))
	}

	// part 1
	var root *Node
	for _, node := range nodes {
		if root == nil {
			root = pparse(node)
		} else {
			root = add(root, pparse(node))
		}
	}
	fmt.Println("Part 1:", magnitude(root))

	// part 2
	best := 0
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			best = max(best, magnitude(add(pparse(nodes[i]), pparse(nodes[j]))))
			best = max(best, magnitude(add(pparse(nodes[j]), pparse(nodes[i]))))
		}
	}
	fmt.Println("Part 2:", best)
}
