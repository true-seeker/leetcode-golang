package linked_lists

import "fmt"

func detectCycle(head *ListNode) *ListNode {
	var addresses []*ListNode

	for head != nil {
		for _, elem := range addresses {
			if elem == head {
				return elem
			}
		}
		addresses = append(addresses, head)
		head = head.Next
	}
	return nil
}

func Test142() {
	a := detectCycle(&ListNode{
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
