package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}

func max(p, q int) int {
	if p > q {
		return p
	}
	return q
}