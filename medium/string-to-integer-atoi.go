package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i, testcase := range []struct {
		input  string
		output int
	}{
		{"42", 42},
		{"+42", 42},
		{"+-2", 0},
		{"    -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
	} {
		result := myAtoi(testcase.input)

		if want, got := testcase.output, result; want != got {
			panic(fmt.Sprintf("[%d] doesn't match; got %+v, want %+v ", i, got, want))
		}
	}
}

/*
Char code for char: [ ] is 32
Char code for char: [+] is 43
Char code for char: [-] is 45
Char code for char: [0] is 48
Char code for char: [1] is 49
Char code for char: [2] is 50
Char code for char: [3] is 51
Char code for char: [4] is 52
Char code for char: [5] is 53
Char code for char: [6] is 54
Char code for char: [7] is 55
Char code for char: [8] is 56
Char code for char: [9] is 57
*/
//https://leetcode.com/problems/string-to-integer-atoi/
func myAtoi(str string) int {

	intStarted := false
	negative := false
	digits := make([]int32, 0)
	for _, ch := range str {
		if ch == 32 {
			if !intStarted {
				continue
			} else {
				break
			}
		}

		if ch >= 48 && ch <= 57 {
			intStarted = true
			if ch >= 48 && ch <= 57 {
				digits = append(digits, ch)
			}
		} else if ch == 45 || ch == 43 {
			if intStarted {
				break
			} else if ch == 45 {
				negative = true
			}
			intStarted = true
		} else {
			break
		}
	}

	trimmed := string(digits)
	i, _ := strconv.Atoi(trimmed)
	if negative {
		i = 0 - i
	}
	if i < -2147483648 {
		return -2147483648
	}
	if i > 2147483647 {
		return 2147483647
	}

	return i
}
