package linked_lists

import (
	"fmt"
)

func middleNode(head *ListNode) *ListNode {
	elementsCount := 0
	temp := head
	for temp != nil {
		temp = temp.Next
		elementsCount++
	}
	for i := 0; i < elementsCount/2.0; i++ {
		head = head.Next
	}
	return head
}

func Test876() {
	a := middleNode(&ListNode{
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

	a = middleNode(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	})

	fmt.Println(a)

	a = middleNode(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	})

	fmt.Println(a)
}
