package t_bank

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
				s: "Rumpelstiltskin",
			},
			want: 1,
		},
		{
			name: "exmaple 2",
			args: args{
				s: "Rabbits under moonlight play, elves laugh, singing tales in late summer, kittens in nests.",
			},
			want: 1,
		},
		{
			name: "empty string",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "double target",
			args: args{
				s: "RumpelstiltskinRumpelstiltskin",
			},
			want: 2,
		},
		{
			name: "triple target",
			args: args{
				s: "RumpelstiltskinRumpelstiltskinRumpelstiltskin",
			},
			want: 3,
		},
		{
			name: "with spaces",
			args: args{
				s: "Rumpelstiltskin Rumpelstiltskin",
			},
			want: 2,
		},
		{
			name: "with punctuations",
			args: args{
				s: "Rumpelstiltskin,Rumpelstiltskin.Rumpelstiltskin;Rumpelstiltskin.",
			},
			want: 4,
		},
		{
			name: "not enought",
			args: args{
				s: "Rumpelstiltski",
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countWord(tt.args.s))
			assert.Equal(t, tt.want, countWord2(tt.args.s))
		})
	}
}
