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
