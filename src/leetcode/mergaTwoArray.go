package leetcode

// 推荐解法，从后往前
func MergaTwoArray(nums1, nums2 []int, m, n int) {
	if n < 1 {
		return
	}
	if m < 1 {
		copy(nums1, nums2)
		return
	}
	temp := nums1[0:m+1]
	n1, n2, all := 0, 0, m
	for ;n1<all;n1++ {
		if n2 == len(nums2) {
			break
		}
		// 切片典型问题：当 temp 大小超过 nums1 时，temp重新申请了一个新数组，故而不再修改 nums1
		if nums1[n1] > nums2[n2] {
			temp = append(temp[0:n1+1], temp[n1:]...)
			temp[n1] = nums2[n2]
			n2++
			all++
		}
	}
	
	if n2 < len(nums2) {
		temp = append(temp[0:n1], nums2[n2:]...)
	}
	// 将切片 temp 中的数据拷贝到 nums1 中
	copy(nums1,temp[0:m+n])
}