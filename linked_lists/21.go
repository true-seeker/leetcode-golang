package linked_lists

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	var ar []int
	for list1 != nil {
		ar = append(ar, list1.Val)
		list1 = list1.Next
	}
	for list2 != nil {
		ar = append(ar, list2.Val)
		list2 = list2.Next
	}

	sort.Ints(ar)

	var result = &ListNode{}
	var resultHead = result
	if len(ar) > 0 {
		result.Val = ar[0]
	} else {
		return resultHead
	}

	for i := 1; i < len(ar); i++ {
		result.Next = &ListNode{
			Val:  ar[i],
			Next: nil,
		}
		result = result.Next
	}

	return resultHead
}

func Test21() {
	a := mergeTwoLists(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}, &ListNode{
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
