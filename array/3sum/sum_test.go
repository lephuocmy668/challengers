package sum

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	type testCase struct {
		name string
		args args
		want [][]int
	}
	tests := []testCase{
		testCase{
			name: "Case 1",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{
				[]int{-1, -1, 2},
				[]int{-1, 0, 1},
			},
		},
		testCase{
			name: "Case 2",
			args: args{
				nums: []int{-2, 0, 0, 2, 2},
			},
			want: [][]int{
				[]int{-2, 0, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
