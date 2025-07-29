package remove_linked_list_elements

import (
	"testing"

	"github.com/nemca/leetcode/models"
	. "github.com/nemca/leetcode/models"
	"github.com/stretchr/testify/assert"
)

func TestRemoveElements(t *testing.T) {
	type args struct {
		head *models.ListNode
		val  int
	}

	tests := []struct {
		name string
		args args
		want *models.ListNode
	}{
		{
			name: "example 1",
			args: args{
				head: MakeSingleLinkedList([]int{1, 2, 6, 3, 4, 5, 6}),
				val:  6,
			},
			want: MakeSingleLinkedList([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "example 2",
			args: args{
				head: nil,
				val:  1,
			},
			want: nil,
		},
		{
			name: "example 3",
			args: args{
				head: MakeSingleLinkedList([]int{7, 7, 7, 7}),
				val:  7,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeElements(tt.args.head, tt.args.val))
		})
	}
}
