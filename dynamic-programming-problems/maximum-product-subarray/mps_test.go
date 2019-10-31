package mps

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	type unitTest struct {
		name string
		args args
		want int
	}
	tests := []unitTest{
		unitTest{
			name: "Test Case 1",
			args: args{
				nums: []int{2, 3, -2, 4},
			},
			want: 6,
		},
		// unitTest{
		// 	name: "Test Case 2",
		// 	args: args{
		// 		nums: []int{-2, 3},
		// 	},
		// 	want: 3,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
