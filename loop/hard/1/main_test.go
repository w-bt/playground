package main

import "testing"

func TestNumberToWords(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Single digit",
			args: args{num: 7},
			want: "Seven",
		},
		{
			name: "Double digits",
			args: args{num: 42},
			want: "Forty Two",
		},
		{
			name: "Hundreds",
			args: args{num: 123},
			want: "One Hundred Twenty Three",
		},
		{
			name: "Thousands",
			args: args{num: 12345},
			want: "Twelve Thousand Three Hundred Forty Five",
		},
		{
			name: "Millions",
			args: args{num: 1234567},
			want: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
		},
		{
			name: "Zero",
			args: args{num: 0},
			want: "Zero",
		},
		{
			name: "Max 32-bit integer",
			args: args{num: 2147483647},
			want: "Two Billion One Hundred Forty Seven Million Four Hundred Eighty Three Thousand Six Hundred Forty Seven",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberToWords(tt.args.num); got != tt.want {
				t.Errorf("NumberToWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
