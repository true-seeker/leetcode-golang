package linked_lists

import "fmt"

func isPalindrome(head *ListNode) bool {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	arrLen := len(arr)
	for i := 0; i < arrLen/2; i++ {
		if arr[i] != arr[arrLen-i-1] {
			return false
		}
	}
	return true
}

func Test234() {
	a := isPalindrome(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	})

	fmt.Println(a)

	a = isPalindrome(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	})

	fmt.Println(a)

	a = isPalindrome(&ListNode{
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
