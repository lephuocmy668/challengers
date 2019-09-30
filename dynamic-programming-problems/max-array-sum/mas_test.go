package mas

import "testing"

func Test_maxSubsetSum(t *testing.T) {
	type args struct {
		arr []int32
	}
	type testCase struct {
		name string
		args args
		want int32
	}
	tests := []testCase{
		testCase{
			name: "Test Case 1",
			args: args{
				arr: []int32{3, 7, 4, 6, 5},
			},
			want: 13,
		},
		testCase{
			name: "Test Case 2",
			args: args{
				arr: []int32{2, 1, 5, 8, 4},
			},
			want: 11,
		},
		testCase{
			name: "Test Case 3",
			args: args{
				arr: []int32{3, 5, -7, 8, 10},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubsetSum(tt.args.arr); got != tt.want {
				t.Errorf("maxSubsetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
