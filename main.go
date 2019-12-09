package main

import (
	"HelloWorld/src/leetcode"
	"fmt"
)

func main() {
	// a := &leetcode.ListNode{1,&leetcode.ListNode{2,nil}}

	// a := []int{9}
	// b:=leetcode.PlusOne(a)

	// a:=leetcode.MinConstructor()
	// a.Push(-2)
	// a.Push(0)
	// a.Push(-3)
	// fmt.Println(a.GetMin())
	// a.Pop()
	// fmt.Println(a.Top())
	// fmt.Println(a.GetMin())

	// a:=[]int{0,1,0,2,1,0,1,3,2,1,2,1}
	// fmt.Println(leetcode.TrapRain(a))

	a := "a________xcddcba"
	fmt.Println(leetcode.IsPalindrome3(a))
}
