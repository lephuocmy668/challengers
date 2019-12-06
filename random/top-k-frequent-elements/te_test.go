package te

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	type testCase struct {
		name string
		args args
		want []int
	}
	tests := []testCase{
		testCase{
			name: "Test Case 1",
			args: args{
				nums: []int{1, 2, 2, 2, 2, 1, 2, 4, 5, 6, 4, 4, 5, 9, 0, 7, 8, 11},
				k:    4,
			},
			want: []int{2, 4, 5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
