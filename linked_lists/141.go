package linked_lists

import "fmt"

func hasCycle(head *ListNode) bool {
	a := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := a[head]; ok {
			return true
		}
		a[head] = struct{}{}
		head = head.Next

	}
	return false
}

func Test141() {
	a := hasCycle(&ListNode{
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
