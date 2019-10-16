package dw

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
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
				s: "12",
			},
			want: 2,
		},
		testCase{
			name: "Test Case 2",
			args: args{
				s: "10",
			},
			want: 1,
		},
		testCase{
			name: "Test Case 3",
			args: args{
				s: "101",
			},
			want: 1,
		},
		testCase{
			name: "Test Case 4",
			args: args{
				s: "17",
			},
			want: 2,
		},
		testCase{
			name: "Test Case 5",
			args: args{
				s: "1212",
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
