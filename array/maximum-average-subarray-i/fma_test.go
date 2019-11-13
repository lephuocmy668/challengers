package fma

import "testing"

func Test_findMaxAverage(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	type testCase struct {
		name string
		args args
		want float64
	}

	tests := []testCase{
		testCase{
			name: "Test Case 1",
			args: args{
				nums: []int{1, 12, -5, -6, 50, 3},
				k:    4,
			},
			want: 12.75,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxAverage(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findMaxAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
