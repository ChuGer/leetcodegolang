package main

import (
	"fmt"
)

func main() {
	for i, testcase := range []struct {
		height []int
		square int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	} {
		square := maxArea(testcase.height)

		if want, got := testcase.square, square; want != got {
			println(fmt.Sprintf("[%d] doesn't match; got %+v, want %+v ", i, got, want))
		} else {
			println(fmt.Sprintf("[%d] validated", i))
		}
	}
}

//https://leetcode.com/problems/container-with-most-water/
func maxArea(height []int) int {
	square := 0
	for i, j := 0, len(height)-1; i < j; {
		length := j - i
		var heigth int
		if height[i] < height[j] {
			heigth = height[i]
			i++
		} else {
			heigth = height[j]
			j--
		}
		tsquare := length * heigth
		if tsquare > square {
			square = tsquare
		}
	}
	return square
}
