package leetcode

func RotateArray(nums []int, k int) {
	if k == 0 {
		return
	}
	alen := len(nums)
	count := 0
	for start:=0;count<alen;start++ {
		current := start
		prev := nums[start]

		for b:=true; b; b=(current!=start) {
			next := (current+k)%alen
			prev, nums[next] = nums[next], prev
			current = next
			count++
		}
	}
}