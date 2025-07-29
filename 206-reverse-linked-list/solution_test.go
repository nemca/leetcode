package reverse_linked_list

import (
	"testing"

	. "github.com/nemca/leetcode/models"
	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
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
				head: MakeSingleLinkedList([]int{1, 2, 3, 4, 5}),
			},
			want: MakeSingleLinkedList([]int{5, 4, 3, 2, 1}),
		},
		{
			name: "example 2",
			args: args{
				head: MakeSingleLinkedList([]int{1, 2}),
			},
			want: MakeSingleLinkedList([]int{2, 1}),
		},
		{
			name: "example 3",
			args: args{
				head: nil,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseList(tt.args.head))
		})
	}
}
