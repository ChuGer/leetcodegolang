package main

import (
	"fmt"
)

func main() {
	for i, testcase := range []struct {
		input   string
		pattern string
		matches bool
	}{
		{"abcdede", "ab.*de", true},
		{"ab", ".*", true},
		{"ab", ".*..", true},
		{"ab", ".*..c*", true},
		{"ab", ".*..c", false},
		{"", "bab", false},
		{"afs", ".*af.*s", true},
		{"", "", true},
		{"", "c*", true},
		{"", ".*", true},
		{"", ".", false},
		{"a", "", false},
		{"a", ".*..a*", false},
		{"aaa", "aaaa", false},
		{"a", "ab*", true},
		{"aaa", "ab*a*c*a", true},
		{"aaa", "a*a", true},
		{"aaa", "a*a*", true},
		{"ab", ".*c", false},
		{"aa", "a", false},
		{"aa", "a*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
	} {
		result := isMatch(testcase.input, testcase.pattern)

		if want, got := testcase.matches, result; want != got {
			println(fmt.Sprintf("[%d] strings doesn't match; got %+v, want %+v ", i, got, want))
		} else {
			println(fmt.Sprintf("[%d] validated", i))
		}
	}
}

/*
., 46
*, 42
a, 97
..
z, 122
*/
//https://leetcode.com/problems/regular-expression-matching/
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(p) > 1 && p[1] == 42 {
		return matchStar(p[0], s, p[2:])
	}
	if len(s) != 0 && (p[0] == 46 || p[0] == s[0]) {
		return isMatch(s[1:], p[1:])
	}
	return false
}

func matchStar(ch uint8, s string, p string) bool {
	if isMatch(s, p) {
		return true
	}
	if len(s) == 0 || (ch != 46 && s[0] != ch) {
		return false
	}
	return matchStar(ch, s[1:], p)
}
