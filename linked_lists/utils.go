package linked_lists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) print() {
	head := l
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func create(elems ...int) *ListNode {
	head := &ListNode{
		Val:  elems[0],
		Next: nil,
	}
	iter := head

	for i := 1; i < len(elems); i++ {
		next := ListNode{
			Val:  elems[i],
			Next: nil,
		}
		iter.Next = &next
		iter = &next
	}
	return head
}
