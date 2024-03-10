package linked_lists

func deleteMiddle(head *ListNode) *ListNode {
	size := 0
	iter := head
	for iter != nil {
		size++
		iter = iter.Next
	}

	mid := size / 2
	iter = head
	prev := &ListNode{}
	i := 0

	for i != mid {
		i++
		prev = iter
		iter = iter.Next
	}
	prev.Next = iter.Next

	if prev.Val == 0 {
		head = nil
	}
	return head
}

func Test2095() {
	a := deleteMiddle(create(1, 3, 4, 7, 1, 2, 6))
	a.print()
	a = deleteMiddle(create(1, 2, 3, 4))
	a.print()
	a = deleteMiddle(create(2, 1))
	a.print()
	a = deleteMiddle(create(2))
	a.print()
}
