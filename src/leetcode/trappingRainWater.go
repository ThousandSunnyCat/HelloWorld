package leetcode

// 42
// 有优化空间
func TrapRain(height []int) int {

	if len(height) < 2 {
		return 0
	}
	
	allArea, n, allHeight := 0, len(height), height[0]
	leftLast, rightLast := 0, n-1
	for i,j := 1, 0; i < n; i++ {
		j = n - 1 - i
		allHeight += height[i]
		if height[i] > height[leftLast] {
			allArea += height[leftLast] * (i - leftLast)
			leftLast = i
		}
		if height[j] > height[rightLast] {
			allArea += height[rightLast] * (rightLast - j)
			rightLast = j
		}
	}

	rainArea := allArea - allHeight
	if leftLast == rightLast {
		rainArea += height[leftLast]
	} else {
		rainArea += height[leftLast] * (rightLast - leftLast + 1)
	}

	return rainArea
}