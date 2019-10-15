package lg

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	type testCase struct {
		name string
		args args
		want string
	}
	tests := []testCase{
		testCase{
			name: "Test Case 1",
			args: args{
				s: "abcdcdc",
			},
			want: "cdcdc",
		},
		testCase{
			name: "Test Case 2",
			args: args{
				s: "aaa",
			},
			want: "aaa",
		},
		testCase{
			name: "Test Case 3",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		testCase{
			name: "Test Case 4",
			args: args{
				s: "",
			},
			want: "",
		},
		testCase{
			name: "Test Case 5",
			args: args{
				s: "aaaa",
			},
			want: "aaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
