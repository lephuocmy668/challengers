package ssm

import "testing"

func Test_sumSubarrayMins(t *testing.T) {
	type args struct {
		A []int
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
				A: []int{3, 1, 2, 4},
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSubarrayMins(tt.args.A); got != tt.want {
				t.Errorf("sumSubarrayMins() = %v, want %v", got, tt.want)
			}
		})
	}
}
