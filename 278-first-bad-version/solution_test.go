package first_bad_version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstBadVersion(t *testing.T) {
	type args struct {
		n   int
		bad int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				n:   5,
				bad: 4,
			},
			want: 4,
		},
		{
			name: "example 2",
			args: args{
				n:   1,
				bad: 1,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bad = tt.args.bad
			assert.Equal(t, tt.want, firstBadVersion(tt.args.n))
		})
	}
}
