package merge_two_sorted_lists

import (
	"testing"

	"github.com/nemca/leetcode/models"
	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	type args struct {
		list1 *models.ListNode
		list2 *models.ListNode
	}

	tests := []struct {
		name string
		args args
		want *models.ListNode
	}{
		{
			name: "example 1",
			args: args{
				list1: &models.ListNode{
					Val: 1,
					Next: &models.ListNode{
						Val: 2,
						Next: &models.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				list2: &models.ListNode{
					Val: 1,
					Next: &models.ListNode{
						Val: 3,
						Next: &models.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &models.ListNode{
				Val: 1,
				Next: &models.ListNode{
					Val: 1,
					Next: &models.ListNode{
						Val: 2,
						Next: &models.ListNode{
							Val: 3,
							Next: &models.ListNode{
								Val: 4,
								Next: &models.ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
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
				list2: &models.ListNode{
					Val:  0,
					Next: nil,
				},
			},
			want: &models.ListNode{
				Val:  0,
				Next: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, mergeTwoLists(tt.args.list1, tt.args.list2))
		})
	}
}
