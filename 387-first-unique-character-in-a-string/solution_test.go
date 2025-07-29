package first_unique_character_in_a_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				s: "leetcode",
			},
			want: 0,
		},
		{
			name: "example 2",
			args: args{
				s: "loveleetcode",
			},
			want: 2,
		},
		{
			name: "example 3",
			args: args{
				s: "aabb",
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, firstUniqChar(tt.args.s))
			assert.Equal(t, tt.want, firstUniqChar2(tt.args.s))
			assert.Equal(t, tt.want, firstUniqChar3(tt.args.s))
		})
	}
}
