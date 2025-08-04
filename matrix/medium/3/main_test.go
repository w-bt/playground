package main

import "testing"

func TestWordSearch(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1 - ABCCED",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCCED",
			},
			want: true,
		},
		{
			name: "Example 2 - SEE",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "SEE",
			},
			want: true,
		},
		{
			name: "Example 3 - ABCB",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCB",
			},
			want: false,
		},
		{
			name: "Single letter match",
			args: args{
				board: [][]byte{
					{'A'},
				},
				word: "A",
			},
			want: true,
		},
		{
			name: "Single letter no match",
			args: args{
				board: [][]byte{
					{'A'},
				},
				word: "B",
			},
			want: false,
		},
		{
			name: "Diagonal not allowed",
			args: args{
				board: [][]byte{
					{'A', 'B'},
					{'C', 'D'},
				},
				word: "AD",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordSearch(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("WordSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
