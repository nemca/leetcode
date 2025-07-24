package reverse_linked_list

import (
	"testing"

	"github.com/nemca/leetcode/models"
	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	type args struct {
		head *models.ListNode
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
			want: &models.ListNode{
				Val: 5,
				Next: &models.ListNode{
					Val: 4,
					Next: &models.ListNode{
						Val: 3,
						Next: &models.ListNode{
							Val: 2,
							Next: &models.ListNode{
								Val:  1,
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
				head: &models.ListNode{
					Val: 1,
					Next: &models.ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
			want: &models.ListNode{
				Val: 2,
				Next: &models.ListNode{
					Val:  1,
					Next: nil,
				},
			},
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
