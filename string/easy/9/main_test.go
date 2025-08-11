package main

import "testing"

func TestIsIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1: egg and add",
			args: args{s: "egg", t: "add"},
			want: true,
		},
		{
			name: "Example 2: foo and bar",
			args: args{s: "foo", t: "bar"},
			want: false,
		},
		{
			name: "Example 3: paper and title",
			args: args{s: "paper", t: "title"},
			want: true,
		},
		{
			name: "Same single char",
			args: args{s: "a", t: "a"},
			want: true,
		},
		{
			name: "Different single char",
			args: args{s: "a", t: "b"},
			want: true, // Still isomorphic (single char can map to another single char)
		},
		{
			name: "Repeated char maps to different",
			args: args{s: "ab", t: "aa"},
			want: false,
		},
		{
			name: "Repeated char maps to different",
			args: args{s: "bbbaaaba", t: "aaabbbba"},
			want: false,
		},
		{
			name: "Special chars mapping",
			args: args{s: "ab!a", t: "cd#c"},
			want: true,
		},
		{
			name: "Case sensitive check",
			args: args{s: "Paper", t: "Title"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
