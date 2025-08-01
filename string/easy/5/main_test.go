package main

import "testing"

func TestLongestWord(t *testing.T) {
	type args struct {
		sen string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Simple input with clear longest word",
			args: args{sen: "Hello world123 567"},
			want: "world123",
		},
		{
			name: "Two longest words, return first",
			args: args{sen: "abc defg hijk lmn"},
			want: "defg", // defg and hijk are same length, defg comes first
		},
		{
			name: "Longest word with numbers",
			args: args{sen: "funny1234 haha2023 ok"},
			want: "funny1234",
		},
		{
			name: "Sentence with punctuation",
			args: args{sen: "Wow! This... is, amazing!"},
			want: "amazing",
		},
		{
			name: "Single word",
			args: args{sen: "onlyword"},
			want: "onlyword",
		},
		{
			name: "Words with equal length and punctuation",
			args: args{sen: "a! bb? cc... dd"},
			want: "bb",
		},
		{
			name: "Empty-like input with punctuation only between words",
			args: args{sen: "x! y? z."},
			want: "x",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestWord(tt.args.sen); got != tt.want {
				t.Errorf("LongestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
