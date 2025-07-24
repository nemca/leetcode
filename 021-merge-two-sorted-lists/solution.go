package merge_two_sorted_lists

import (
	"github.com/nemca/leetcode/models"
)

func mergeTwoLists(list1 *models.ListNode, list2 *models.ListNode) *models.ListNode {
	dummy := &models.ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next
}
