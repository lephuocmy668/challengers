package ssek

import "testing"

func Test_subarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
				k:    17,
				nums: []int{4, 5, 8, 7, 0, 3, 1, -8, 25},
			},
			want: 1,
		},
		testCase{
			name: "Test Case 1",
			args: args{
				k:    2,
				nums: []int{1, 1, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
