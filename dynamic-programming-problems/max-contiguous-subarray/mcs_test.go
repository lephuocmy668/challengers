package mcs

import (
	"reflect"
	"testing"
)

func Test_maxSubarray(t *testing.T) {
	type args struct {
		arr []int32
	}
	type testCase struct {
		name string
		args args
		want []int32
	}
	tests := []testCase{
		testCase{
			name: "Test Case 1",
			args: args{
				arr: []int32{-2, -3, -1, -4, -6},
			},
			want: []int32{-1, -1},
		},
		testCase{
			name: "Test Case 2",
			args: args{
				arr: []int32{2, -1, 2, 3, 4, -5},
			},
			want: []int32{10, 11},
		},
		testCase{
			name: "Test Case 3",
			args: args{
				arr: []int32{1},
			},
			want: []int32{1, 1},
		},
		testCase{
			name: "Test Case 4",
			args: args{
				arr: []int32{-1, -2, -3, -4, -5, -6},
			},
			want: []int32{-1, -1},
		},
		testCase{
			name: "Test Case 5",
			args: args{
				arr: []int32{1, -2},
			},
			want: []int32{1, 1},
		},
		testCase{
			name: "Test Case 6",
			args: args{
				arr: []int32{1, 2, 3},
			},
			want: []int32{6, 6},
		},
		testCase{
			name: "Test Case 7",
			args: args{
				arr: []int32{-10},
			},
			want: []int32{-10, -10},
		},
		testCase{
			name: "Test Case 8",
			args: args{
				arr: []int32{1, -1, -1, -1, -1, 5},
			},
			want: []int32{5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarray(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
