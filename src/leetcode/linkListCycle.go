package leetcode

// 141
func HasCycle(head *ListNode) bool {
	// 快慢指针法
	if head == nil {
		return false
	}
	p, q := head, head
	for q.Next != nil && q.Next.Next != nil {
		p, q = p.Next, q.Next.Next
		if p == q {
			return true
		}
	}
	return false
}

func HasCycle2(head *ListNode) (b bool) {
	// 快慢指针法
	defer func() {
        if r := recover(); r != nil {
            b = false
        }
    }()
    p, q := head, head
    for {
        p, q = p.Next, q.Next.Next
        if p == q {
            b = true
            break
        }
    }
    return b
}

//142
func DetectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p, q, re := head, head, false
    for q.Next != nil && q.Next.Next != nil {
		p, q = p.Next, q.Next.Next
		if p == q {
			re = true
			break
		}
	}

	if !re {
		return nil
	}

	for p = head;p != q;p,q = p.Next,q.Next {}
	return p
}

func DetectCycle2(head *ListNode) (re *ListNode) {
	defer func() {
        if r := recover(); r != nil {
            re = nil
        }
    }()
    p, q := head, head
    for {
		p, q = p.Next, q.Next.Next
		if p == q {
			break
		}
    }
	
	for p = head;p != q;p,q = p.Next,q.Next {}
	return p
}