package maxprofit

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
		fee    int
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
				prices: []int{1, 3, 2, 8, 4, 9},
				fee:    2,
			},
			want: 8,
		},
		testCase{
			name: "Test Case 1",
			args: args{
				prices: []int{1, 3, 7, 5, 10},
				fee:    3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
