package fibonacci_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "example 2",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "example 3",
			args: args{
				n: 4,
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fib(tt.args.n))
			assert.Equal(t, tt.want, fib2(tt.args.n))
		})
	}
}
