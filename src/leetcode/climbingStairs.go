package leetcode

func ClimbingStairs(n int) int {
	// 这TM是一个类斐波那契数列，用递归法时间效率为 O(2^n)

	if n <= 2 {
		return n
	}
	n -= 2
	p, q := 1, 2
	for ;n>=1;n-- {
		p, q = q, p+q
	}

	return q
}