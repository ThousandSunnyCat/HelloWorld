package leetcode

func MergeTwoList(l1, l2 *ListNode) *ListNode {
	// 递归
	k := &ListNode{}
	mergeList(k, l1, l2)

	return k.Next
}

func mergeList(k, l1, l2 *ListNode) {

	if l1 == nil || l2 == nil {
		if l1 != nil {
			k.Next = l1
		} else if l2 != nil {
			k.Next = l2
		}

		return
	}

	if l1.Val > l2.Val {
		k.Next = l2
		mergeList(l2, l1, l2.Next)
	} else {
		k.Next = l1
		mergeList(l1, l1.Next, l2)
	}
}