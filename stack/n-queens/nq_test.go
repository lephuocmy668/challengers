package main

import (
	"reflect"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	type testCase struct {
		name string
		args args
		want [][]string
	}

	tests := []testCase{
		testCase{
			name: "Test Case 1",
			args: args{
				n: 4,
			},
			want: [][]string{
				[]string{".Q..", "...Q", "Q...", "..Q."},
				[]string{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		testCase{
			name: "Test Case 2",
			args: args{
				n: 6,
			},
			want: [][]string{
				[]string{".Q....", "...Q..", ".....Q", "Q.....", "..Q...", "....Q."},
				[]string{"..Q...", ".....Q", ".Q....", "....Q.", "Q.....", "...Q.."},
				[]string{"...Q..", "Q.....", "....Q.", ".Q....", ".....Q", "..Q..."},
				[]string{"....Q.", "..Q...", "Q.....", ".....Q", "...Q..", ".Q...."},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
