package tmn

import "testing"

func Test_thirdMax(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{3, 2, 1},
			},
			want: 1,
		},
		testCase{
			name: "Test Case 2",
			args: args{
				nums: []int{1, 2},
			},
			want: 2,
		},
		testCase{
			name: "Test Case 3",
			args: args{
				nums: []int{2, 2, 3, 1},
			},
			want: 1,
		},
		testCase{
			name: "Test Case 4",
			args: args{
				nums: []int{1, 2, -2147483648},
			},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thirdMax(tt.args.nums); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
