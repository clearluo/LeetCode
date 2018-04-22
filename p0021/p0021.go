package p0021

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	for i, p := 0, l.Next; p != nil; i, p = i+1, p.Next {
		if i > 0 {
			fmt.Printf("->")
		}
		fmt.Printf("%v", p.Val)
	}
	fmt.Printf("\n")
}

func (l *ListNode) Add(val int) {
	q := l
	p := l.Next
	for ; p != nil; p, q = p.Next, q.Next {
	}
	(*q).Next = &ListNode{
		Val: val,
	}

}
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &ListNode{}
	np := head
	for ; l1 != nil && l2 != nil; np = np.Next {
		if l1.Val < l2.Val {
			np.Next = l1
			l1 = l1.Next
		} else {
			np.Next = l2
			l2 = l2.Next
		}
	}
	if l1 != nil {
		np.Next = l1
	}
	if l2 != nil {
		np.Next = l2
	}
	return head.Next
}
