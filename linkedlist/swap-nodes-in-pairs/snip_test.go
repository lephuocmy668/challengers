package snip

import (
	"reflect"
	"testing"

	"github.com/lephuocmy668/challengers/linkedlist"
)

func Test_verify(t *testing.T) {
	node4 := linkedlist.ListNode{
		Val: 4,
	}
	node3 := linkedlist.ListNode{
		Val:  3,
		Next: &node4,
	}
	node2 := linkedlist.ListNode{
		Val:  2,
		Next: &node3,
	}
	node1 := linkedlist.ListNode{
		Val:  1,
		Next: &node2,
	}
	type args struct {
		head *linkedlist.ListNode
	}
	type testCase struct {
		name string
		args args
		want []int
	}
	tests := []testCase{
		testCase{
			name: "Test Case 1",
			args: args{
				head: &node1,
			},
			want: []int{2, 1, 4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verify(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verify() = %v, want %v", got, tt.want)
			}
		})
	}
}
