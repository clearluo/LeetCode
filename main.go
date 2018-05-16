package main

import (
	"LeetCode/easy"
)

func main() {
	var data *easy.ListNode
	data = data.Add(1)
	data = data.Add(1)
	data = data.Add(2)
	ret := easy.DeleteDuplicates(data)
	ret.Print()
}
