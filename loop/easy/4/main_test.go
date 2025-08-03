package main

import "testing"

func TestRomanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Simple I",
			args: args{s: "I"},
			want: 1,
		},
		{
			name: "Simple number",
			args: args{s: "III"},
			want: 3,
		},
		{
			name: "Number with subtraction (IV)",
			args: args{s: "IV"},
			want: 4,
		},
		{
			name: "Number with subtraction (IX)",
			args: args{s: "IX"},
			want: 9,
		},
		{
			name: "Number with mixed values (LVIII)",
			args: args{s: "LVIII"},
			want: 58,
		},
		{
			name: "Complex number (MCMXCIV)",
			args: args{s: "MCMXCIV"},
			want: 1994,
		},
		{
			name: "Maximum value (MMMCMXCIX)",
			args: args{s: "MMMCMXCIX"},
			want: 3999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanToInt(tt.args.s); got != tt.want {
				t.Errorf("RomanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
