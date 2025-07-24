package reverse_linked_list

import "github.com/nemca/leetcode/models"

func reverseList(head *models.ListNode) *models.ListNode {
	if head == nil {
		return nil
	}

	var prev *models.ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
