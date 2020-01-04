package mw

import "testing"

func Test_maximumSwap(t *testing.T) {
	type args struct {
		num int
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
				num: 2736,
			},
			want: 7236,
		},
		testCase{
			name: "Test Case 2",
			args: args{
				num: 1993,
			},
			want: 9913,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSwap(tt.args.num); got != tt.want {
				t.Errorf("maximumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
