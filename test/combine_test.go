package jgstr_test

import (
	"testing"

	jgstr "github.com/Jinglever/go-string"
)

func TestCombineSortedWithSep(t *testing.T) {
	type args struct {
		sep  string
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				sep:  ",",
				strs: []string{"a", "b", "c"},
			},
			want: "a,b,c",
		},
		{
			name: "test2",
			args: args{
				sep:  ",",
				strs: []string{"c", "b", "a"},
			},
			want: "a,b,c",
		},
		{
			name: "test3",
			args: args{
				sep:  ",",
				strs: []string{"a", "b", "c", "a"},
			},
			want: "a,a,b,c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jgstr.CombineSortedWithSep(tt.args.sep, tt.args.strs); got != tt.want {
				t.Errorf("CombineSortedWithSep() = %v, want %v", got, tt.want)
			}
		})
	}
}
