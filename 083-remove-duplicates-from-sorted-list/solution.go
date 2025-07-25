package remove_duplicates_from_sorted_list

import (
	"github.com/nemca/leetcode/models"
)

func deleteDuplicates(head *models.ListNode) *models.ListNode {
	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
			continue
		}
		curr = curr.Next
	}

	return head
}
