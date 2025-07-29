package remove_duplicates_from_sorted_list

import (
	"testing"

	. "github.com/nemca/leetcode/models"
	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example 1",
			args: args{
				head: MakeSingleLinkedList([]int{1, 1, 2}),
			},
			want: MakeSingleLinkedList([]int{1, 2}),
		},
		{
			name: "example 2",
			args: args{
				head: MakeSingleLinkedList([]int{1, 1, 2, 3, 3}),
			},
			want: MakeSingleLinkedList([]int{1, 2, 3}),
		},
		{
			name: "example 3",
			args: args{
				head: MakeSingleLinkedList([]int{1, 1, 1}),
			},
			want: MakeSingleLinkedList([]int{1}),
		},
		{
			name: "empty list",
			args: args{
				head: MakeSingleLinkedList([]int{}),
			},
			want: MakeSingleLinkedList([]int{}),
		},
		{
			name: "single list node",
			args: args{
				head: MakeSingleLinkedList([]int{1}),
			},
			want: MakeSingleLinkedList([]int{1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, deleteDuplicates(tt.args.head))
		})
	}
}
