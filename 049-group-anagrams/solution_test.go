package group_anagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	type args struct {
		strs []string
	}

	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "example 1",
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name: "example 2",
			args: args{
				strs: []string{""},
			},
			want: [][]string{{""}},
		},
		{
			name: "example 3",
			args: args{
				strs: []string{"a"},
			},
			want: [][]string{{"a"}},
		},
		{
			name: "exmaple 4",
			args: args{
				strs: []string{"a", "aa", "aaa", "b"},
			},
			want: [][]string{{"a"}, {"aa"}, {"aaa"}, {"b"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, groupAnagrams(tt.args.strs))
			assert.Equal(t, tt.want, groupAnagrams2(tt.args.strs))
		})
	}
}
