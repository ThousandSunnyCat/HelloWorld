package leetcode

func maxSlidingWindow(nums []int, k int) []int {

	n := len(nums)

	if n * k <= 0 {
		return []int{}
	}

	if k == 1 {
		return nums
	}

	re := make([]int, n - k + 1)
	left, right, j := make([]int, n), make([]int, n), 0
	
	for i:=1;i<n;i++ {
		j = n - 1 - i
		if i % k == 0 {
			left[i] = nums[i]
		} else {
			left[i] = max(left[i-1], nums[i])
		}

		if (n - i) % k == 0 {
			right[j] = nums[j]
		} else {
			right[j] = max(right[j+1], nums[j])
		}
	}

	for i:=k-1;i<n;i++ {
		re[i+1-k] = max(left[i], right[i-k+1])
	}

	return re
}