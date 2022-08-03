package roman_number

func middleNode(head *ListNode) *ListNode {
	size := 0

	for child := head; child != nil; child = child.Next {
		size++
	}

	middle := size/2 + 1

	for child := head; child != nil; child = child.Next {
		middle--
		if middle == 0 {
			return child
		}
	}

	return nil
}
