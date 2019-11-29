package main

import (
	"fmt"
	"leetcode"
)

func main() {
	fmt.Println("Hello World");

	arr := [] int {0,0,1,1,1,2,2,3,3,4}
	i := leetcode.RemoveDuplicates(arr)

	fmt.Println(i)
	fmt.Println(arr)
}