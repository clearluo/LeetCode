package main

import (
	"LeetCode/p0026"
	"fmt"
)

func main() {
	data := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	l := p0026.RemoveDuplicates(data)
	fmt.Println("l:", l)
	for i := 0; i < l; i++ {
		fmt.Println(data[i])
	}
}
