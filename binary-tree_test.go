package gotree

import (
	"fmt"
	"testing"
)

func TestCreateBinaryTree(t *testing.T) {
	var testCases = []struct {
		a              []int
		expectedHeight int
	}{
		{a: nil, expectedHeight: 0},
		{a: []int{}, expectedHeight: 0},
		{a: []int{1}, expectedHeight: 0},
		{a: []int{1, 2}, expectedHeight: 1},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.a), func(t *testing.T) {

			root := CreateBinaryTree(tc.a)
			tree := BinaryTree{Root: root}
			result := tree.Height()
			if tc.expectedHeight != result {
				t.Errorf("Expected height %d but received %d", tc.expectedHeight, result)
			}
		})
	}
}
