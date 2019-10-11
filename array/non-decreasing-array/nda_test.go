package nda

import "testing"

func Test_checkPossibility(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{4, 2, 3},
			},
			want: true,
		},
		testCase{
			name: "Test Case 2",
			args: args{
				nums: []int{4, 2, 1},
			},
			want: false,
		},
		testCase{
			name: "Test Case 3",
			args: args{
				nums: []int{2, 3, 3, 2, 4},
			},
			want: true,
		},
		testCase{
			name: "Test Case 4",
			args: args{
				nums: []int{3, 4, 2, 3},
			},
			want: false,
		},
		testCase{
			name: "Test Case 4",
			args: args{
				nums: []int{-1, 4, 2, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPossibility(tt.args.nums); got != tt.want {
				t.Errorf("checkPossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
