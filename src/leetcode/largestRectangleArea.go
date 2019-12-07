package leetcode

// 84
// 重点，需要回刷
func LargestRectangleArea(heights []int) int {

	if len(heights) < 1 {
		return 0
	}

	LRA, maxArea, k := lraCreate(), heights[0], 0
	
	for i:=0;i<len(heights);i++ {
		for k = LRA.top(); k != -1 && heights[k] > heights[i]; k = LRA.top() {
			LRA.pop()
			maxArea = max(maxArea, heights[k] * (i - LRA.top() - 1))
		}

		LRA.push(i)
	}


	for k = LRA.top(); k >= 0; k = LRA.top() {
		LRA.pop()
		maxArea = max(maxArea, heights[k] * (len(heights) - LRA.top() - 1))
	}
	return maxArea
}

type lraStack struct {
	value []int
}

func lraCreate() lraStack {
	return lraStack {
		value: make([]int, 0),
	}
}

func (this *lraStack) push(k int) {
	this.value = append(this.value, k)
}

func (this *lraStack) pop() {
	if this.isEmpty() {
		return
	}
	this.value = this.value[:len(this.value)-1]
}

func (this *lraStack) top() int {
	if this.isEmpty() {
		return -1
	}
	return this.value[len(this.value)-1]
}

func (this *lraStack) isEmpty() bool {
	return len(this.value) < 1
}