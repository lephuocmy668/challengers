package mps

import "testing"

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	type testCase struct {
		name string
		args args
		want int
	}

	tests := []testCase{
		testCase{
			name: "Test Case 1",
			args: args{
				grid: [][]int{
					[]int{1, 3, 1},
					[]int{1, 5, 1},
					[]int{4, 2, 1},
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
