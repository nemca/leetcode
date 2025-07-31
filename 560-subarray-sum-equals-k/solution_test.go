package subarray_sum_equals_k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{1, 1, 1},
				k:    2,
			},
			want: 2,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{1, 2, 3},
				k:    3,
			},
			want: 2,
		},
		{
			name: "example 3",
			args: args{
				nums: []int{1, -1, 0},
				k:    0,
			},
			want: 3,
		},
		{
			name: "exmaple: 4",
			args: args{
				nums: []int{100, 1, 2, 3, 4},
				k:    6,
			},
			want: 1,
		},
		{
			name: "example 5",
			args: args{
				nums: []int{28, 54, 7, -70, 22, 65, -6},
				k:    100,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, subarraySumBrute(tt.args.nums, tt.args.k))
			assert.Equal(t, tt.want, subarraySum(tt.args.nums, tt.args.k))
		})
	}
}
