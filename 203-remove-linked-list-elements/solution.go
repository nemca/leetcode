package remove_linked_list_elements

import (
	. "github.com/nemca/leetcode/models"
)

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for head != nil {
		if head.Val != val {
			tail.Next = &ListNode{
				Val:  head.Val,
				Next: nil,
			}
			tail = tail.Next
		}
		head = head.Next
	}

	return dummy.Next
}
