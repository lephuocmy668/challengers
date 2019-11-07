package ms

import "testing"

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
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
				matrix: [][]byte{
					[]byte("10100"),
					[]byte("10111"),
					[]byte("11111"),
					[]byte("10010"),
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.args.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
