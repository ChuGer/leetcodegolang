package main

import "fmt"

func main() {
	for i, testcase := range []struct {
		search string
		row    int
		output string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"ABCDEFGHIKLMNOPRS", 5, "AISBHKRCGLPDFMOEN"},
		{"ABC", 1, "ABC"},
		{"ABC", 10, "ABC"},
		{"ABC", 2, "ACB"},
	} {
		result := convert(testcase.search, testcase.row)

		if want, got := testcase.output, result; want != got {
			panic(fmt.Sprintf("[%d] strings doesn't match; got %+v, want %+v ", i, got, want))
		}
	}
}

//https://leetcode.com/problems/zigzag-conversion/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	matrix := make([]string, numRows)

	counter := 0
	reverse := false

	for i := 0; i < len(s); i++ {
		matrix[counter] = matrix[counter] + string(s[i])

		if counter == numRows-1 && !reverse {
			reverse = true
		} else if counter == 0 && reverse {
			reverse = false
		}
		if !reverse {
			counter++
		} else {
			counter--
		}
	}

	for i := 1; i < numRows; i++ {
		matrix[0] = matrix[0] + matrix[i]
	}

	return matrix[0]
}
