package main

import "fmt"

func main() {
	for i, testcase := range []struct {
		nums1  []int
		nums2  []int
		median float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{3}, []int{-2, -1}, -1.0},
		{[]int{2}, []int{}, 2.0},
		{[]int{3, 4}, []int{}, 3.5},
	} {
		m := findMedianSortedArrays(testcase.nums1, testcase.nums2)

		if want, got := testcase.median, m; want != got {
			panic(fmt.Sprintf("[%d] doesn't match; got %+v, want %+v ", i, got, want))
		}
	}
}

//https://leetcode.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	var idx1, idx2 int
	lensum := len1 + len2
	if lensum%2 == 0 {
		idx2 = lensum / 2
		idx1 = (lensum / 2) - 1
	} else {
		idx1 = (lensum - 1) / 2
		idx2 = idx1
	}

	var val1 int
	for i, i1, i2 := 0, 0, 0; i < lensum; i++ {
		if i == idx1 {
			if i1 != len1 && (i2 == len2 || nums1[i1] < nums2[i2]) {
				val1 = nums1[i1]
			} else {
				val1 = nums2[i2]
			}
			if idx1 == idx2 {
				return float64(val1)
			}
		}

		if i == idx2 {
			if i1 != len1 && (i2 == len2 || nums1[i1] < nums2[i2]) {
				return median(val1, nums1[i1])
			} else {
				return median(val1, nums2[i2])
			}
		}

		if i1 != len1 && (i2 == len2 || nums1[i1] < nums2[i2]) {
			i1++
		} else {
			i2++
		}
	}
	return 0
}

func median(val1 int, val2 int) float64 {
	return float64(val1+val2) / 2
}
