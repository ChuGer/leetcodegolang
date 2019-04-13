package main

import (
	"fmt"
)

func main() {
	for i, testcase := range []struct {
		search string
		count  int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{" ", 1},
	} {
		assets := lengthOfLongestSubstring(testcase.search)

		if want, got := testcase.count, assets; want != got {
			panic(fmt.Sprintf("[%d] longest subscrting doesn't match; got %+v, want %+v ", i, got, want))
		}
	}
}

//https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	charIndex := make(map[int32]int)
	ans, i := 0, 0
	for j, ch := range s {
		if pos, ok := charIndex[ch]; ok {
			i = Max(pos, i)
		}
		ans = Max(ans, j-i+1)
		charIndex[ch] = j + 1
	}
	return ans
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
