package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
	k := &ListNode{}
	p := k
	count := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			count += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			count += l2.Val
			l2 = l2.Next
		}

		// p.Val = count % 10
		p.Next = &ListNode{count % 10,nil}
		p = p.Next
		count /= 10
	}
	if count != 0 {
		p.Next = &ListNode{count,nil}
	}

	return k.Next
}