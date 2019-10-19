package css

import "testing"

func Test_checkSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	type testCase struct {
		name string
		args args
		want bool
	}
	tests := []testCase{
		testCase{
			name: "Test Case 1",
			args: args{
				nums: []int{23, 2, 4, 6, 7},
				k:    6,
			}, want: true,
		},
		testCase{
			name: "Test Case 2",
			args: args{
				nums: []int{23, 2, 4, 6, 7},
				k:    0,
			}, want: false,
		},
		testCase{
			name: "Test Case 3",
			args: args{
				nums: []int{0, 0},
				k:    0,
			}, want: true,
		},
		testCase{
			name: "Test Case 4",
			args: args{
				nums: []int{0, 1, 0},
				k:    0,
			}, want: false,
		},
		testCase{
			name: "Test Case 5",
			args: args{
				nums: []int{0},
				k:    -1,
			}, want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("checkSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
