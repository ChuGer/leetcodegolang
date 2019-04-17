package main

import "fmt"

func main() {
	for i, testcase := range []struct {
		decimal int
		roman   string
	}{
		{58, "LVIII"},
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{1994, "MCMXCIV"},
	} {
		roman := intToRoman(testcase.decimal)

		if want, got := testcase.roman, roman; want != got {
			println(fmt.Sprintf("[%d] doesn't match; got %+v, want %+v ", i, got, want))
		} else {
			println(fmt.Sprintf("[%d] validated", i))
		}
	}
}

//https://leetcode.com/problems/integer-to-roman/
func intToRoman(num int) string {
	result := ""
	for i := 0; i < 4; i++ {
		last := num % 10
		if last < 4 {
			for j := 0; j < last; j++ {
				result = literals[i*2] + result
			}
		} else if last == 4 {
			result = literals[i*2] + literals[i*2+1] + result
		} else if last >= 5 && last < 9 {
			for j := 0; j < last-5; j++ {
				result = literals[i*2] + result
			}
			result = literals[i*2+1] + result
		} else if last == 9 {
			result = literals[i*2] + literals[i*2+2] + result
		}
		num /= 10
	}

	return result
}

var literals = [...]string{"I", "V", "X", "L", "C", "D", "M"}
