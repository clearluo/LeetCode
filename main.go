package main

import (
	"LeetCode/p0027"
	"fmt"
)

func main() {
	data := []int{0, 1, 2, 2, 3, 0, 4, 2}
	l := p0027.RemoveElement(data, 2)
	fmt.Println("l:", l)
	for i := 0; i < l; i++ {
		fmt.Println(data[i])
	}
}
