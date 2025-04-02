package binarytree

// Comparable is an interface that defines a method for comparing
// two objects of the same type for equality.
type Comparable interface {
	Equals(other Comparable) bool
}

// Node represents a node in a binary tree.
//
// Structures implementing this interface must provide methods to access
// the left and right children of the node, as well as a method to check
// for equality with another node. The Left() and Right() methods should
// return nil if the node does not have a left or right child, respectively.
type Node[T Comparable] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

// Equals checks if two binary tree nodes are equal.
//
// It compares the structure and values of the nodes recursively.
// If both nodes are nil, they are considered equal.
// If one node is nil and the other is not, they are not equal.
// If both nodes are not nil, it checks if their values are equal and
// recursively checks their left and right children for equality.
//
// This function is prone to stack overflow for large trees due to deep recursion.
// Please use FastEquals for large trees.
func Equals[T Comparable](a, b *Node[T]) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if !a.Value.Equals(b.Value) {
		return false
	}
	return Equals(a.Left, b.Left) && Equals(a.Right, b.Right)
}

// FastEquals checks if two binary tree nodes are equal.
//
// It compares the structure and values of the nodes iteratively.
// It is the pretty version of the Equals function.
func FastEquals[T Comparable](a, b *Node[T]) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	la := a
	lb := b
	ra := a
	rb := b
	for {
		for {
			if la == nil && lb == nil {
				break
			}
			if la == nil || lb == nil {
				return false
			}
			if !la.Value.Equals(lb.Value) {
				return false
			}
			la = la.Left
			lb = lb.Left
		}
		if ra == nil && rb == nil {
			break
		}
		ra = ra.Right
		rb = rb.Right
		la = ra
		lb = rb
	}
	return true
}
