package main

import (
	"fmt"
	"math"
)

func main() {
	for i, testcase := range []struct {
		input  int
		output int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
		{2147483648, 0},
		{1534236469, 0},
		{-1563847412, 0},
	} {
		result := reverse(testcase.input)

		if want, got := testcase.output, result; want != got {
			panic(fmt.Sprintf("[%d] doesn't match; got %+v, want %+v ", i, got, want))
		}
	}
}

// https://leetcode.com/problems/reverse-integer/
func reverse(x int) int {
	if x < -2147483648 || x > 2147483647 || x == 0 {
		return 0
	}
	reverseDigits := make([]int, 0)
	for {
		if x == 0 {
			break
		}
		mod := x % 10
		reverseDigits = append(reverseDigits, mod)
		x = (x - mod) / 10
	}

	result := 0
	i := len(reverseDigits) - 1
	for _, v := range reverseDigits {
		result += v * int(math.Pow10(i))
		i--
	}
	if result < -2147483648 || result > 2147483647 {
		return 0
	}
	return result
}
