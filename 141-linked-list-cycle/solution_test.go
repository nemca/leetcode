package linked_list_cycle

import (
	"testing"

	"github.com/nemca/leetcode/models"
	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	type args struct {
		head *models.ListNode
	}

	zero := &models.ListNode{
		Val:  0,
		Next: nil,
	}

	two := &models.ListNode{
		Val:  2,
		Next: zero,
	}

	three := &models.ListNode{
		Val:  3,
		Next: two,
	}

	four := &models.ListNode{
		Val:  -4,
		Next: two,
	}

	zero.Next = four

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				head: three,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, hasCycle(tt.args.head))
		})
	}
}
