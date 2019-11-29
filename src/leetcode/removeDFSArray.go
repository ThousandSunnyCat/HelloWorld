package leetcode

func RemoveDuplicates(nums []int) int {

	if len(nums) < 1 {
		return 1
	}
	// 设置
	Prev := 0

	for i:=1;i<len(nums);i++ {
		if nums[Prev] != nums[i] {
			Prev++
			nums[Prev] = nums[i]
		}
	}

	return Prev + 1
}