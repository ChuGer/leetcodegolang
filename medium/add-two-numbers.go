package main

func main() {
	l1 := ListNode{5, nil}
	l2 := ListNode{5, nil}

	sum := addTwoNumbers(&l1, &l2)
	if sum.Val != 0 || sum.Next.Val != 1 {
		panic("wrong result")
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersInt(l1, l2, 0)
}

func addTwoNumbersInt(l1 *ListNode, l2 *ListNode, extra int) *ListNode {
	if l1 == nil && l2 == nil {
		if extra == 1 {
			return &ListNode{1, nil}
		} else {
			return nil
		}
	}
	var v1, v2 int
	if l1 != nil {
		v1 = l1.Val
	}
	if l2 != nil {
		v2 = l2.Val
	}
	sum := v1 + v2 + extra
	extra = sum / 10

	if l1 == nil {
		return &ListNode{Val: sum % 10, Next: addTwoNumbersInt(nil, l2.Next, extra)}
	} else if l2 == nil {
		return &ListNode{Val: sum % 10, Next: addTwoNumbersInt(l1.Next, nil, extra)}
	} else {
		return &ListNode{Val: sum % 10, Next: addTwoNumbersInt(l1.Next, l2.Next, extra)}
	}
}
