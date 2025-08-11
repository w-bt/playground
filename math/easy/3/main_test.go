package main

import "testing"

func TestIsPowerOfThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "n=1-true",
			args: args{n: 1},
			want: true,
		},
		{
			name: "n=0-false",
			args: args{n: 0},
			want: false,
		},
		{
			name: "negative-1-false",
			args: args{n: -1},
			want: false,
		},
		{
			name: "negative-3-false",
			args: args{n: -3},
			want: false,
		},
		{
			name: "n=2-false",
			args: args{n: 2},
			want: false,
		},
		{
			name: "n=3-true",
			args: args{n: 3},
			want: true,
		},
		{
			name: "n=4-false",
			args: args{n: 4},
			want: false,
		},
		{
			name: "n=9-true",
			args: args{n: 9},
			want: true,
		},
		{
			name: "n=12-false-multiple-of-3-not-power",
			args: args{n: 12},
			want: false,
		},
		{
			name: "n=27-true",
			args: args{n: 27},
			want: true,
		},
		{
			name: "n=45-false-multiple-of-3-not-power",
			args: args{n: 45},
			want: false,
		},
		{
			name: "n=81-true",
			args: args{n: 81},
			want: true,
		},
		{
			name: "n=243-true",
			args: args{n: 243},
			want: true,
		},
		{
			name: "n=244-false-nearby",
			args: args{n: 244},
			want: false,
		},
		{
			name: "n=59049-true-3^10",
			args: args{n: 59049},
			want: true,
		},
		{
			name: "n=1594323-true-3^13",
			args: args{n: 1594323},
			want: true,
		},
		{
			name: "n=1162261467-true-3^19-max-int32-power",
			args: args{n: 1162261467},
			want: true,
		},
		{
			name: "n=1162261466-false-just-below-3^19",
			args: args{n: 1162261466},
			want: false,
		},
		{
			name: "n=1162261468-false-just-above-3^19",
			args: args{n: 1162261468},
			want: false,
		},
		{
			name: "n=1000000000-false-large-non-power",
			args: args{n: 1000000000},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPowerOfThree(tt.args.n); got != tt.want {
				t.Errorf("IsPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
