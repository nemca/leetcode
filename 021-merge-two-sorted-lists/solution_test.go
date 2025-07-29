package merge_two_sorted_lists

import (
	"testing"

	. "github.com/nemca/leetcode/models"
	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example 1",
			args: args{
				list1: MakeSingleLinkedList([]int{1, 2, 4}),
				list2: MakeSingleLinkedList([]int{1, 3, 4}),
			},
			want: MakeSingleLinkedList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name: "example 2",
			args: args{
				list1: nil,
				list2: nil,
			},
			want: nil,
		},
		{
			name: "example 3",
			args: args{
				list1: nil,
				list2: MakeSingleLinkedList([]int{0}),
			},
			want: MakeSingleLinkedList([]int{0}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, mergeTwoLists(tt.args.list1, tt.args.list2))
		})
	}
}
