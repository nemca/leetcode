package models

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// String implements fmt.Stringer interface.
func (l *ListNode) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for node := l; node != nil; node = node.Next {
		sb.WriteString(fmt.Sprintf("%d", node.Val))
		if node.Next != nil {
			sb.WriteString(" -> ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

// MakeSingleLinkedList generates single linjed list from slice of integers.
func MakeSingleLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, val := range nums[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}
