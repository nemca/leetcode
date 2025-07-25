package remove_linked_list_elements

import (
	"testing"

	"github.com/nemca/leetcode/models"
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
				head: &models.ListNode{
					Val: 1,
					Next: &models.ListNode{
						Val: 2,
						Next: &models.ListNode{
							Val: 6,
							Next: &models.ListNode{
								Val: 3,
								Next: &models.ListNode{
									Val: 4,
									Next: &models.ListNode{
										Val: 5,
										Next: &models.ListNode{
											Val:  6,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
				val: 6,
			},
			want: &models.ListNode{
				Val: 1,
				Next: &models.ListNode{
					Val: 2,
					Next: &models.ListNode{
						Val: 3,
						Next: &models.ListNode{
							Val: 4,
							Next: &models.ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
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
				head: &models.ListNode{
					Val: 7,
					Next: &models.ListNode{
						Val: 7,
						Next: &models.ListNode{
							Val: 7,
							Next: &models.ListNode{
								Val:  7,
								Next: nil,
							},
						},
					},
				},
				val: 7,
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
