package leetcode

func MaxArea(height []int) int {
	//目的：key 的差值乘以 min(value1, value2)

	l, r, area, narea := 0, len(height) - 1, 0, 0

	for l < r {
		if height[l] < height[r] {
			narea = (r - l) * height[l]
			l++
		} else {
			narea = (r - l) * height[r]
			r--
		}

		if narea > area {
			area = narea
		}
	}

	return area
}