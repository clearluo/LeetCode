package main

import (
	"LeetCode/easy"
	"fmt"
)

func main() {
	//fmt.Println(easy.MoveZeroes([]int{0, 1, 0, 3, 12}))
	arr := []int{0, 1, 0, 3, 12}
	easy.MoveZeroes(arr[:])
	fmt.Println(arr)
}
