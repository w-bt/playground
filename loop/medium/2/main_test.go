package main

import "testing"

func TestIntToRoman(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1 - 3 becomes III",
			args: args{num: 3},
			want: "III",
		},
		{
			name: "Example 2 - 58 becomes LVIII",
			args: args{num: 58},
			want: "LVIII", // L = 50, V = 5, III = 3
		},
		{
			name: "Example 3 - 1994 becomes MCMXCIV",
			args: args{num: 1994},
			want: "MCMXCIV", // M=1000, CM=900, XC=90, IV=4
		},
		{
			name: "Smallest number",
			args: args{num: 1},
			want: "I",
		},
		{
			name: "Largest number",
			args: args{num: 3999},
			want: "MMMCMXCIX",
		},
		{
			name: "Number with subtraction in middle (944)",
			args: args{num: 944},
			want: "CMXLIV", // CM=900, XL=40, IV=4
		},
		{
			name: "Number without subtraction (1666)",
			args: args{num: 1666},
			want: "MDCLXVI", // M=1000, D=500, C=100, L=50, X=10, V=5, I=1
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToRoman(tt.args.num); got != tt.want {
				t.Errorf("IntToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
