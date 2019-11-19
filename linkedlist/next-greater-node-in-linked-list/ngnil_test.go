package ngnil

import (
	"reflect"
	"testing"

	"github.com/lephuocmy668/challengers/linkedlist"
)

// Input: [2,1,5,5]
// Output: [5,5,0,0]
// Example 2:
func Test_nextLargerNodes(t *testing.T) {
	node4 := linkedlist.ListNode{
		Val: 5,
	}
	node3 := linkedlist.ListNode{
		Val:  5,
		Next: &node4,
	}
	node2 := linkedlist.ListNode{
		Val:  1,
		Next: &node3,
	}
	node1 := linkedlist.ListNode{
		Val:  2,
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
			want: []int{5, 5, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextLargerNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextLargerNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Input: [2,7,4,3,5]
// Output: [7,0,5,5,0]
// Example 3:
func Test_nextLargerNodes2(t *testing.T) {
	node5 := linkedlist.ListNode{
		Val: 5,
	}
	node4 := linkedlist.ListNode{
		Val:  3,
		Next: &node5,
	}
	node3 := linkedlist.ListNode{
		Val:  4,
		Next: &node4,
	}
	node2 := linkedlist.ListNode{
		Val:  7,
		Next: &node3,
	}
	node1 := linkedlist.ListNode{
		Val:  2,
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
			want: []int{7, 0, 5, 5, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextLargerNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextLargerNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Input: [1,7,5,1,9,2,5,1]
// Output: [7,9,9,9,0,5,0,0]
func Test_nextLargerNodes3(t *testing.T) {
	node8 := linkedlist.ListNode{
		Val: 1,
	}
	node7 := linkedlist.ListNode{
		Val:  5,
		Next: &node8,
	}
	node6 := linkedlist.ListNode{
		Val:  2,
		Next: &node7,
	}
	node5 := linkedlist.ListNode{
		Val:  9,
		Next: &node6,
	}
	node4 := linkedlist.ListNode{
		Val:  1,
		Next: &node5,
	}
	node3 := linkedlist.ListNode{
		Val:  5,
		Next: &node4,
	}
	node2 := linkedlist.ListNode{
		Val:  7,
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
			want: []int{7, 9, 9, 9, 0, 5, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextLargerNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextLargerNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
