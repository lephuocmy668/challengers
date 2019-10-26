package kcms

import "testing"

func Test_kConcatenationMaxSum(t *testing.T) {
	type args struct {
		arr []int
		k   int
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
				arr: []int{1, -2, 1},
				k:   5,
			},
			want: 2,
		},
		testCase{
			name: "Test Case 2",
			args: args{
				arr: []int{1, 2},
				k:   3,
			},
			want: 9,
		},
		testCase{
			name: "Test Case 3",
			args: args{
				arr: []int{1, 0, 4, 1, 4},
				k:   4,
			},
			want: 40,
		},
		testCase{
			name: "Test Case 3",
			args: args{
				arr: []int{-5, 4, -4, -3, 5, -3},
				k:   3,
			},
			want: 5,
		},
		testCase{
			name: "Test Case 4",
			args: args{
				arr: []int{1, 2},
				k:   3,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kConcatenationMaxSum(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("kConcatenationMaxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
