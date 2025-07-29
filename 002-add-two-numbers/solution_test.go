package add_two_numbers

import (
	"testing"

	. "github.com/nemca/leetcode/models"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example 1",
			args: args{
				l1: MakeSingleLinkedList([]int{2, 4, 3}),
				l2: MakeSingleLinkedList([]int{5, 6, 4}),
			},
			want: MakeSingleLinkedList([]int{7, 0, 8}),
		},
		{
			name: "example 2",
			args: args{
				l1: MakeSingleLinkedList([]int{0}),
				l2: MakeSingleLinkedList([]int{0}),
			},
			want: MakeSingleLinkedList([]int{0}),
		},
		{
			name: "example 3",
			args: args{
				l1: MakeSingleLinkedList([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: MakeSingleLinkedList([]int{9, 9, 9, 9}),
			},
			want: MakeSingleLinkedList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
		{
			name: "example 4",
			args: args{
				l1: MakeSingleLinkedList([]int{9}),
				l2: MakeSingleLinkedList([]int{1, 9, 9, 9, 9, 9, 9, 9, 9, 9}),
			},
			want: MakeSingleLinkedList([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		},
		{
			name: "example 5",
			args: args{
				l1: MakeSingleLinkedList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
				l2: MakeSingleLinkedList([]int{5, 6, 4}),
			},
			want: MakeSingleLinkedList([]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, addTwoNumbers(tt.args.l1, tt.args.l2))
		})
	}
}
