package leetcode

// 66
// 根据 leetcode 国际站的 diaa 大神的代码修改
func PlusOne(digits []int) []int {
	for k:=len(digits)-1;k>=0;k-- {
		if digits[k] < 9 {
			digits[k]++
			return digits
		}
		digits[k]=0
	}

	digits = append([]int{1}, digits[:]...)
	return digits
}