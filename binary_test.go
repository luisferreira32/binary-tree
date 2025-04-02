package binarytree

import (
	"fmt"
	"testing"
	"time"
)

type mockComparable struct {
	value int
}

func (m mockComparable) Equals(other Comparable) bool {
	o, ok := other.(mockComparable)
	return ok && m.value == o.value
}

func Test_Equals(t *testing.T) {
	tests := []struct {
		name     string
		tree1    *Node[mockComparable]
		tree2    *Node[mockComparable]
		expected bool
	}{
		{
			name: "Equal trees",
			tree1: &Node[mockComparable]{
				Value: mockComparable{value: 1},
				Left:  &Node[mockComparable]{Value: mockComparable{value: 2}},
				Right: &Node[mockComparable]{Value: mockComparable{value: 3}},
			},
			tree2: &Node[mockComparable]{
				Value: mockComparable{value: 1},
				Left:  &Node[mockComparable]{Value: mockComparable{value: 2}},
				Right: &Node[mockComparable]{Value: mockComparable{value: 3}},
			},
			expected: true,
		},
		{
			name: "Different trees",
			tree1: &Node[mockComparable]{
				Value: mockComparable{value: 1},
				Left:  &Node[mockComparable]{Value: mockComparable{value: 2}},
				Right: &Node[mockComparable]{Value: mockComparable{value: 3}},
			},
			tree2: &Node[mockComparable]{
				Value: mockComparable{value: 1},
				Left:  &Node[mockComparable]{Value: mockComparable{value: 4}},
				Right: &Node[mockComparable]{Value: mockComparable{value: 3}},
			},
			expected: false,
		},
		{
			name: "Uneven trees",
			tree1: &Node[mockComparable]{
				Value: mockComparable{value: 1},
				Left:  &Node[mockComparable]{Value: mockComparable{value: 2}},
			},
			tree2: &Node[mockComparable]{
				Value: mockComparable{value: 1},
				Left:  &Node[mockComparable]{Value: mockComparable{value: 2}},
				Right: &Node[mockComparable]{Value: mockComparable{value: 3}},
			},
			expected: false,
		},
		{
			name: "One tree is nil",
			tree1: &Node[mockComparable]{
				Value: mockComparable{value: 1},
				Left:  &Node[mockComparable]{Value: mockComparable{value: 2}},
				Right: &Node[mockComparable]{Value: mockComparable{value: 3}},
			},
			tree2:    nil,
			expected: false,
		},
		{
			name:     "Both trees are nil",
			tree1:    nil,
			tree2:    nil,
			expected: true,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			if result := Equals(testcase.tree1, testcase.tree2); result != testcase.expected {
				t.Errorf("Equals failed: expected %v, got %v", testcase.expected, result)
			}
		})
	}
}

func Test_FastEquals(t *testing.T) {
	tests := []struct {
		name     string
		tree1    *Node[mockComparable]
		tree2    *Node[mockComparable]
		expected bool
	}{
		{
			name: "Equal trees",
			tree1: &Node[mockComparable]{
				Value: mockComparable{value: 1},
				Left:  &Node[mockComparable]{Value: mockComparable{value: 2}},
				Right: &Node[mockComparable]{Value: mockComparable{value: 3}},
			},
			tree2: &Node[mockComparable]{
				Value: mockComparable{value: 1},
				Left:  &Node[mockComparable]{Value: mockComparable{value: 2}},
				Right: &Node[mockComparable]{Value: mockComparable{value: 3}},
			},
			expected: true,
		},
		{
			name: "Different trees",
			tree1: &Node[mockComparable]{
				Value: mockComparable{value: 1},
				Left:  &Node[mockComparable]{Value: mockComparable{value: 2}},
				Right: &Node[mockComparable]{Value: mockComparable{value: 3}},
			},
			tree2: &Node[mockComparable]{
				Value: mockComparable{value: 1},
				Left:  &Node[mockComparable]{Value: mockComparable{value: 4}},
				Right: &Node[mockComparable]{Value: mockComparable{value: 3}},
			},
			expected: false,
		},
		{
			name: "Uneven trees on right branch",
			tree1: &Node[mockComparable]{
				Value: mockComparable{value: 1},
				Left:  &Node[mockComparable]{Value: mockComparable{value: 2}},
			},
			tree2: &Node[mockComparable]{
				Value: mockComparable{value: 1},
				Left:  &Node[mockComparable]{Value: mockComparable{value: 2}},
				Right: &Node[mockComparable]{Value: mockComparable{value: 3}},
			},
			expected: false,
		},
		{
			name: "Uneven trees on left branch",
			tree1: &Node[mockComparable]{
				Value: mockComparable{value: 1},
				Left:  &Node[mockComparable]{Value: mockComparable{value: 2}},
				Right: &Node[mockComparable]{Value: mockComparable{value: 3}},
			},
			tree2: &Node[mockComparable]{
				Value: mockComparable{value: 1},
				Left:  &Node[mockComparable]{Value: mockComparable{value: 2}},
			},
			expected: false,
		},
		{
			name: "One tree is nil",
			tree1: &Node[mockComparable]{
				Value: mockComparable{value: 1},
				Left:  &Node[mockComparable]{Value: mockComparable{value: 2}},
				Right: &Node[mockComparable]{Value: mockComparable{value: 3}},
			},
			tree2:    nil,
			expected: false,
		},
		{
			name:     "Both trees are nil",
			tree1:    nil,
			tree2:    nil,
			expected: true,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			if result := FastEquals(testcase.tree1, testcase.tree2); result != testcase.expected {
				t.Errorf("FastEquals failed: expected %v, got %v", testcase.expected, result)
			}
		})
	}
}

// generateTree generates a binary tree with `size` nodes, where each node's value is its position in the tree.
func generateTree(size int) *Node[mockComparable] {
	if size == 0 {
		return nil
	}
	nodes := make([]*Node[mockComparable], size)
	for i := 0; i < size; i++ {
		nodes[i] = &Node[mockComparable]{Value: mockComparable{value: i}}
	}
	for i := 0; i < size; i++ {
		leftIndex := 2*i + 1
		rightIndex := 2*i + 2
		if leftIndex < size {
			nodes[i].Left = nodes[leftIndex]
		}
		if rightIndex < size {
			nodes[i].Right = nodes[rightIndex]
		}
	}
	return nodes[0]
}

func Test_Performance(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected bool
	}{
		{"10      nodes", 10, true},
		{"100     nodes", 100, true},
		{"1000    nodes", 1000, true},
		{"10000   nodes", 10000, true},
		{"100000  nodes", 100000, true},
		{"1000000 nodes", 1000000, true},
	}

	for _, testcase := range tests {
		t.Run(fmt.Sprintf("Equals_%s", testcase.name), func(t *testing.T) {
			tree1 := generateTree(testcase.size)
			tree2 := generateTree(testcase.size)

			start := time.Now()
			result := Equals(tree1, tree2)
			duration := time.Since(start)

			if result != testcase.expected {
				t.Errorf("Equals failed: expected %v, got %v", testcase.expected, result)
			}
			fmt.Printf("Equals (%s): %v\n", testcase.name, duration)
		})
	}

	for _, testcase := range tests {
		t.Run(fmt.Sprintf("FastEquals_%s", testcase.name), func(t *testing.T) {
			tree1 := generateTree(testcase.size)
			tree2 := generateTree(testcase.size)

			start := time.Now()
			result := FastEquals(tree1, tree2)
			duration := time.Since(start)

			if result != testcase.expected {
				t.Errorf("FastEquals failed: expected %v, got %v", testcase.expected, result)
			}
			fmt.Printf("FastEquals (%s): %v\n", testcase.name, duration)
		})
	}
}
