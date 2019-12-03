package leetcode

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head, prev, head.Next = head.Next, head, prev
	}
	return prev
}

func ReverseList2(head *ListNode) *ListNode {
    if (head == nil || head.Next == nil) {
        return head;
    }
    p := ReverseList2(head.Next);
    head.Next.Next = head;
    head.Next = nil;
    return p;
}