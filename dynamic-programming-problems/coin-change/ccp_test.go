package ccp

import "testing"

func Test_getWays(t *testing.T) {
	type args struct {
		n int32
		c []int64
	}

	type testCase struct {
		name string
		args args
		want int64
	}
	tests := []testCase{
		testCase{
			name: "Test case 1",
			args: args{
				n: 4,
				c: []int64{1, 2, 3},
			},
			want: 4,
		},
		testCase{
			name: "Test case 1",
			args: args{
				n: 10,
				c: []int64{2, 5, 3, 6},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWays(tt.args.n, tt.args.c); got != tt.want {
				t.Errorf("getWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
