package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func generateLinkList(intArray []int) *ListNode {
	sentinel := &ListNode{
		Next: nil,
	}
	last := sentinel
	for _, value := range intArray {
		head := &ListNode{
			Val:  value,
			Next: nil,
		}
		last.Next = head
		last = head
	}
	return sentinel.Next
}

func TestReverseKGroup(t *testing.T) {
	testCases := []struct {
		name  string
		input struct {
			values []int
			k      int
		}
		expect []int
	}{
		{
			name: "case 1",
			input: struct {
				values []int
				k      int
			}{
				values: []int{1, 2, 3, 4, 5},
				k:      2,
			},
			expect: []int{2, 1, 4, 3, 5},
		},
		{
			name: "case 2",
			input: struct {
				values []int
				k      int
			}{
				values: []int{1, 2, 3, 4, 5},
				k:      3,
			},
			expect: []int{3, 2, 1, 4, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputLinkedList := generateLinkList(tc.input.values)
			output := reverseKGroup(inputLinkedList, tc.input.k)
			outputIntArray := make([]int, 0)
			for output != nil {
				outputIntArray = append(outputIntArray, output.Val)
				output = output.Next
			}
			assert.Equal(t, tc.expect, outputIntArray)
		})
	}
}
