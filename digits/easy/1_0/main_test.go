package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	type args struct {
		x int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positive palindrome",
			args: args{x: 121},
			want: true,
		},
		{
			name: "Negative number",
			args: args{x: -121},
			want: false,
		},
		{
			name: "Not a palindrome",
			args: args{x: 10},
			want: false,
		},
		{
			name: "Single digit",
			args: args{x: 7},
			want: true,
		},
		{
			name: "Zero is a palindrome",
			args: args{x: 0},
			want: true,
		},
		{
			name: "Large palindrome",
			args: args{x: 123454321},
			want: true,
		},
		{
			name: "Large non-palindrome",
			args: args{x: 123456789},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.x); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
