package sas

import "testing"

func Test_substrings(t *testing.T) {
	type args struct {
		n string
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
				n: "16",
			},
			want: 23,
		},
		testCase{
			name: "Test Case 2",
			args: args{
				n: "123",
			},
			want: 164,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := substrings(tt.args.n); got != tt.want {
				t.Errorf("substrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
