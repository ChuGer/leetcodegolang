package main

import "fmt"

func main() {
	nums := []int{-3, 4, 3, 90}
	target := 0
	ints := twoSum(nums, target)
	if len(ints) != 2 {
		panic(fmt.Sprintf("wrong size of indicies: %d", len(ints)))
	}
	if ints[0] != 0 || ints[1] != 2 {
		panic("wrong result")
	}
}

//https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i, v := range nums {
		if fi, ok := m[v]; ok {
			return []int{fi, i}
		} else {
			m[target-v] = i
		}
	}
	return nil
}
