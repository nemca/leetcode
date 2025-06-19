package valid_parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid parenthesis",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "valid parenthesis",
			args: args{
				s: "((",
			},
			want: false,
		},
		{
			name: "valid parenthesis",
			args: args{
				s: "",
			},
			want: true,
		},
		{
			name: "valid parenthesis",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "valid parenthesis",
			args: args{
				s: "(]",
			},
			want: false,
		},
		{
			name: "valid parenthesis",
			args: args{
				s: "([)]",
			},
			want: false,
		},
		{
			name: "valid parenthesis",
			args: args{
				s: "{[]}",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isValid(tt.args.s))
		})
	}
}
