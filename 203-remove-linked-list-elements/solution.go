package remove_linked_list_elements

import (
	"github.com/nemca/leetcode/models"
)

func removeElements(head *models.ListNode, val int) *models.ListNode {
	dummy := &models.ListNode{}
	tail := dummy

	for head != nil {
		if head.Val != val {
			tail.Next = &models.ListNode{
				Val:  head.Val,
				Next: nil,
			}
			tail = tail.Next
		}
		head = head.Next
	}

	return dummy.Next
}
