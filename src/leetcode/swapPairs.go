package leetcode

// 24
func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head.Next
    var prev *ListNode
    for head != nil && head.Next != nil {
		head, head.Next.Next, prev = head.Next.Next, head, head
		if head != nil {
			prev.Next = head.Next
		}
	}
	prev.Next = head
    return cur
}


// 25
// 代码流程不清晰，需要优化
func ReverseKGroup(head *ListNode, k int) (re *ListNode) {
	if head == nil || head.Next == nil {
		return head
	}
	
	var prev, khead, ktail *ListNode
	for head != nil {
		ktail, prev = head, nil

		for i:=0;i<k;i++ {
			if head == nil {
				for ;i>0;i-- {
					prev.Next, prev, head = head, prev.Next, prev
				}
				if re == nil {
					return head
				}
				khead.Next = head
				return
			}
			head.Next, head, prev = prev, head.Next, head
		}

		if re == nil {
			khead, re = ktail, prev
			continue
		}
		khead, khead.Next = ktail, prev

	}

	return
}