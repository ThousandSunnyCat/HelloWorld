package leetcode

func TwoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	// 设置
	maps := make(map[int]int)

	// 查找
	for k, v := range nums {
		if v < target {
			// find
			if tk, t := maps[target - v]; t {
				return []int{tk, k}
			}

			// set
			maps[v] = k
		}
	}

	return nil
}