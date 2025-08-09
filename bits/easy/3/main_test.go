package main

import "testing"

func TestAddBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Basic addition without carry",
			args: args{
				a: "11",
				b: "1",
			},
			want: "100",
		},
		{
			name: "Basic addition with multiple carries",
			args: args{
				a: "1010",
				b: "1011",
			},
			want: "10101",
		},
		{
			name: "Carry through all digits",
			args: args{
				a: "111",
				b: "1",
			},
			want: "1000",
		},
		{
			name: "Different lengths",
			args: args{
				a: "1",
				b: "11101",
			},
			want: "11110",
		},
		{
			name: "Both zeros",
			args: args{
				a: "0",
				b: "0",
			},
			want: "0",
		},
		{
			name: "Longer carry chain",
			args: args{
				a: "101111",
				b: "000111",
			},
			want: "110110",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AddBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
