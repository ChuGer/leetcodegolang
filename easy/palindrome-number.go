package main

import (
	"fmt"
	"math"
)

func main() {
	for i, testcase := range []struct {
		input  int
		output bool
	}{
		{123, false},
		{121, true},
		{1221, true},
		{-121, false},
		{0, true},
		{123454321, true},
	} {
		result := isPalindrome(testcase.input)

		if want, got := testcase.output, result; want != got {
			panic(fmt.Sprintf("[%d] doesn't match; got %+v, want %+v ", i, got, want))
		}
	}
}

//https://leetcode.com/problems/palindrome-number/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	iLength := digits(x)
	if iLength == 1 {
		return true
	}
	for i := iLength - 2; i >= 0; i -= 2 {
		last := x % 10
		first := x / int(math.Pow10(i+1))
		if last != first {
			return false
		}
		x = (x / 10) % int(math.Pow10(i))
	}
	return true
}

func digits(x int) int {
	if x == 0 {
		return 0
	}
	return 1 + digits(x/10)
}

func isPalindromeSimple(x int) bool {
	if x < 0 {
		return false
	}

	digits := make([]uint8, 0)
	for x != 0 {
		digits = append(digits, uint8(x%10))
		x /= 10
	}
	for i, j := 0, len(digits)-1; i < len(digits)/2; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}
