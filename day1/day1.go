package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data1.txt")
	if err != nil {
	}

	var nums []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
		}
		nums = append(nums, num)
	}

	ans1, ans2 := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			ans1++
		}
		if i > 2 && nums[i] > nums[i-3] {
			ans2++
		}
	}
	fmt.Println("Part 1: ", ans1)
	fmt.Println("Part 2: ", ans2)
}
