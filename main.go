package main

import (
	"fmt"
	"leetcode"
)

func main() {
	for bb:=0;bb<9;bb++ {
		test(bb)
		fmt.Println("______________________________")
	}
	// test(2)
}

func test(k int) {
	arr := []int{1,2,3,4}
	leetcode.RotateArray(arr,k)

	for _, v := range arr {
		fmt.Println(v)
	}
}