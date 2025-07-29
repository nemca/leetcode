package middle_of_the_linked_list

import (
	. "github.com/nemca/leetcode/models"
)

func middleNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	count := 1
	next := head
	ll := make(map[int]*ListNode)

	for next.Next != nil {
		count++
		next = next.Next
		ll[count] = next
	}

	if count%2 != 0 {
		count++
	} else {
		count += 2
	}

	mid := count / 2

	return ll[mid]
}

func middleNode2(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
