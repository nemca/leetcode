package linked_list_cycle

import "github.com/nemca/leetcode/models"

func hasCycle(head *models.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}

		if fast.Next == nil {
			fast = nil
		} else {
			fast = fast.Next.Next
		}

		if slow.Next == nil {
			slow = nil
		} else {
			slow = slow.Next
		}
	}

	return false
}
