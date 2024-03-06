package linked_lists

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	iter := head
	for iter.Next != nil {
		nextElem := iter.Next
		gcdNode := ListNode{
			Val:  gcd(iter.Val, nextElem.Val),
			Next: nextElem,
		}
		iter.Next = &gcdNode
		iter = nextElem
	}
	return head
}
