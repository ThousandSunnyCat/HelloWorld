package leetcode

func MergeTwoList(l1, l2 *ListNode) *ListNode {
	// 边界条件不要漏
	if l1 == nil {return l2}
    if l2 == nil {return l1}
	// 递归，会消耗栈空间，故空间复杂度为 O(n+m)
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