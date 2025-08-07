package main

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "regular sentence",
			args: args{s: "Hello World"},
			want: 5,
		},
		{
			name: "trailing spaces",
			args: args{s: "   fly me   to   the moon  "},
			want: 4,
		},
		{
			name: "no trailing space",
			args: args{s: "luffy is still joyboy"},
			want: 6,
		},
		{
			name: "single word",
			args: args{s: "ChatGPT"},
			want: 7,
		},
		{
			name: "single word with trailing spaces",
			args: args{s: "AI     "},
			want: 2,
		},
		{
			name: "empty string",
			args: args{s: ""},
			want: 0,
		},
		{
			name: "only spaces",
			args: args{s: "     "},
			want: 0,
		},
		{
			name: "last word is one character",
			args: args{s: "Test x"},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
