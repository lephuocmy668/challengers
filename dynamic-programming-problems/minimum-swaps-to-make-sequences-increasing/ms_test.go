package ms

import "testing"

func Test_minSwap(t *testing.T) {
	type args struct {
		A []int
		B []int
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
				A: []int{1, 3, 5, 4},
				B: []int{1, 2, 3, 7},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwap(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("minSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
