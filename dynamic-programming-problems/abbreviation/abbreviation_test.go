package abbreviation

import "testing"

func Test_abbreviation(t *testing.T) {
	type args struct {
		a string
		b string
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
				a: "daBcd",
				b: "ABC",
			},
			want: "YES",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abbreviation(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("abbreviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
