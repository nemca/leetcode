package middle_of_the_linked_list

import (
	"testing"

	. "github.com/nemca/leetcode/models"
	"github.com/stretchr/testify/assert"
)

func TestMiddleNode(t *testing.T) {
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
			want: MakeSingleLinkedList([]int{3, 4, 5}),
		},
		{
			name: "example 2",
			args: args{
				head: MakeSingleLinkedList([]int{1, 2, 3, 4, 5, 6}),
			},
			want: MakeSingleLinkedList([]int{4, 5, 6}),
		},
		{
			name: "example 3",
			args: args{
				head: MakeSingleLinkedList([]int{1}),
			},
			want: MakeSingleLinkedList([]int{1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, middleNode(tt.args.head))
		})
	}
}

func TestMiddleNode2(t *testing.T) {
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
			want: MakeSingleLinkedList([]int{3, 4, 5}),
		},
		{
			name: "example 2",
			args: args{
				head: MakeSingleLinkedList([]int{1, 2, 3, 4, 5, 6}),
			},
			want: MakeSingleLinkedList([]int{4, 5, 6}),
		},
		{
			name: "example 3",
			args: args{
				head: MakeSingleLinkedList([]int{1}),
			},
			want: MakeSingleLinkedList([]int{1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, middleNode(tt.args.head))
		})
	}
}
