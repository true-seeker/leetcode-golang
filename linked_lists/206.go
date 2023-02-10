package linked_lists

import (
	"fmt"
)

func reverseList(head *ListNode) *ListNode {
	curr := head
	next := &ListNode{}
	prev := (*ListNode)(nil)
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func Test206() {
	a := reverseList(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	})

	fmt.Println(a)
}
