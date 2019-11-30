package leetcode

func RemoveDuplicates(nums []int) int {

	if len(nums) < 1 {
		return 1
	}
	// 设置
	p := 0

	for i:=1;i<len(nums);i++ {
		if nums[p] != nums[i] {
			p++
			nums[p] = nums[i]
		}
	}

	return p + 1
}