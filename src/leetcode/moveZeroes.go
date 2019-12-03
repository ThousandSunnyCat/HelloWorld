package leetcode

func MoveZeroes(nums []int) {
	// 另类排序
	p := -1
	// 遍历
	for k, v := range nums {
		if v != 0 {
			if p >= 0 {
				// nums[p] 可直接改为 0 值
				nums[k], nums[p] = nums[p], nums[k]
				p++
			}
		} else if p < 0 {
			p = k
		}
	}
}