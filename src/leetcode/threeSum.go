package leetcode

func ThreeSum(nums []int) [][]int {

	if len(nums) < 3 {
		return [][]int{}
	}

	// 排序，随机快排
	quickSort(nums)

	// 双指针遍历
	re := make([][]int,0,0)
	for i:=0;i<len(nums);i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		p, q := i+1, len(nums)-1
		for p < q {
			if p > i+1 && nums[p] == nums[p-1] {
				p++
				continue
			}
			a := nums[p] + nums[q] + nums[i]
			if a < 0 {
				p++
			} else if a > 0 {
				q--
			} else {
				re = append(re, []int{nums[i], nums[p], nums[q]})
				p++
			}
		}
	}

	return re
}

func quickSort(r []int) {
	if len(r) < 2 {
		return
	}
	q := 1
	for i:=1;i<len(r);i++ {
		if r[i] <= r[0] {
			r[i], r[q] = r[q], r[i]
			q++
		}
	}
	r[0], r[q-1] = r[q-1], r[0]
	quickSort(r[0:q-1])
	quickSort(r[q:])
}


func ThreeSum2(nums []int) [][]int {

	re := [][]int{}
	
	for a, v := range nums {
		for i,iv:= range nums {
			if i == a {
				continue
			}
			for j,jv := range nums {
				if i == j || a == j {
					continue
				}
				if v+iv+jv == 0 {
					re = append(re, []int{v,iv,jv})
				}
			}
		}
	}

	return re
}