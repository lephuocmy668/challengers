package lds

import (
	"reflect"
	"testing"
)

func Test_largestDivisibleSubset(t *testing.T) {
	type args struct {
		nums []int
	}

	type testCase struct {
		name string
		args args
		want []int
	}

	tests := []testCase{
		testCase{
			name: "largestDivisibleSubset of empty set",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
		testCase{
			name: "largestDivisibleSubset should return largest divisuble subset",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: []int{1, 2},
		},
		testCase{
			name: "largestDivisibleSubset should return largest divisuble subset",
			args: args{
				nums: []int{546, 547},
			},
			want: []int{546},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestDivisibleSubset(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestDivisibleSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
