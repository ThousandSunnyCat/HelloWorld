package main

import (
	"fmt"
	"leetcode"
)

func main() {
	// a := &leetcode.ListNode{1,&leetcode.ListNode{2,nil}}

	a := []int{9}
	b:=leetcode.PlusOne(a)
	fmt.Println(b)
}