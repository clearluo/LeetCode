package main

import "LeetCode/p0021"

func main() {
	//var l1 p0021.ListNode
	l1 := &p0021.ListNode{}
	l1.Add(1)
	l1.Add(2)
	l1.Add(4)
	l1.Print()
	l2 := &p0021.ListNode{}
	l2.Add(1)
	l2.Add(3)
	l2.Add(4)
	l2.Print()
	nl := p0021.MergeTwoLists(l1, l2)
	nl.Print()
}
