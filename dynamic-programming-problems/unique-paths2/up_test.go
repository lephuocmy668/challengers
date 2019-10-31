package up

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	type testCase struct {
		name string
		args args
		want int
	}
	tests := []testCase{
		testCase{
			name: "Test case 1",
			args: args{
				obstacleGrid: [][]int{
					[]int{0, 0, 0},
					[]int{0, 1, 0},
					[]int{0, 0, 0},
				},
			},
			want: 2,
		},
		testCase{
			name: "Test case 2",
			args: args{
				obstacleGrid: [][]int{
					[]int{0, 0},
					[]int{1, 1},
					[]int{0, 0},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
