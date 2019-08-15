package redblack

import (
	"reflect"
	"testing"
)

func TestTraversal(t *testing.T) {
	n1 := node{key: 1}
	n2 := node{key: 2}
	n3 := node{key: 3}
	n4 := node{key: 4}

	tr := &tree{root: &n2}
	n2.left = &n1
	n2.right = &n3
	n3.right = &n4

	expectedOrder := []*node{&n1, &n2, &n3, &n4}
	if !reflect.DeepEqual(expectedOrder, tr.traverse()) {
		t.Error("Improper ordering found in traversal")
	}
}

// genTestTree is a helper function that generates a balanced size 15
// binary tree for testing various cases, The tree is generated with
// keys in order for easily working with the traverse function
//
// The in order traversal of this tree results in [n1..n15]
//
// This also returns a list of pointers to non leaf nodes, and leaf
// nodes so we test out rotation and panic behaviour
func genTestTree() (t *tree, nonLeaves []*node, leaves []*node) {
	n1 := node{key: 1}
	n2 := node{key: 2}
	n3 := node{key: 3}
	n4 := node{key: 4}
	n5 := node{key: 5}
	n6 := node{key: 6}
	n7 := node{key: 7}
	n8 := node{key: 8}
	n9 := node{key: 9}
	n10 := node{key: 10}
	n11 := node{key: 11}
	n12 := node{key: 12}
	n13 := node{key: 13}
	n14 := node{key: 14}
	n15 := node{key: 15}

	t = &tree{root: &n8}
	n8.left = &n4
	n8.right = &n14
	n4.left = &n2
	n4.right = &n6
	n2.left = &n1
	n2.right = &n3
	n6.left = &n5
	n6.right = &n7
	n12.left = &n10
	n12.right = &n14
	n10.left = &n9
	n10.right = &n11
	n14.left = &n13
	n14.right = &n15

	return t, []*node{&n8, &n4, &n12, &n2, &n6, &n10, &n14}, []*node{&n1, &n3, &n5, &n7, &n9, &n11, &n13, &n15}
}

func deepEqualTree(t1 *tree, t2 *tree) bool {
	return deepEqualHelper(t1.root, t2.root)
}

func deepEqualHelper(n1 *node, n2 *node) bool {
	if n1 == nil && n2 == nil {
		return true
	} else if n1.key != n2.key {
		return false
	}
	return deepEqualHelper(n1.left, n2.left) &&
		deepEqualHelper(n1.right, n2.right)
}

// The functions for rotation should be invertible, so calling left
// and right successfully should modify the structure such that it was
// never modified at all
func TestRotationInversion(t *testing.T) {
	tree, nonLeaves, _ := genTestTree()

	for _, node := range nonLeaves {
		child := node.right

		tree.rotateLeft(node)
		tree.rotateRight(child)

		if newTree, _, _ := genTestTree(); !deepEqualTree(tree, newTree) {
			t.Errorf("Node with key %v broke inversion", node.key)
		}
	}
}

func TestLeftRotation(t *testing.T) {
	n1 := node{}
	n2 := node{}
	n3 := node{}
	n4 := node{}

	tr := &tree{root: &n1}
	n1.right = &n2
	n2.left = &n3
	n2.right = &n4

	expectedOrder := []*node{&n1, &n3, &n2, &n4}
	if !reflect.DeepEqual(expectedOrder, tr.traverse()) {
		t.Error("Improper left rotation")
	}
}

func TestRightRotation(t *testing.T) {
	n1 := node{}
	n2 := node{}
	n3 := node{}
	n4 := node{}

	tr := &tree{root: &n1}
	n1.right = &n2
	n2.left = &n3
	n2.right = &n4

	expectedOrder := []*node{&n1, &n3, &n2, &n4}
	if !reflect.DeepEqual(expectedOrder, tr.traverse()) {
		t.Error("Improper right rotation")
	}
}

func willPanic(f func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	f()
	return panicked
}

// When rotate is called with on a node of a tree that is a leaf, we
// should panic
func TestRotationLeafPanic(t *testing.T) {
	tree, _, leaves := genTestTree()
	for _, leaf := range leaves {
		if !willPanic(func() {
			tree.rotateLeft(leaf)
		}) {
			t.Error("Left rotation failed to panic on leaf node")
		}
		if !willPanic(func() {
			tree.rotateRight(leaf)
		}) {
			t.Error("Right rotation failed to panic on leaf node")
		}
	}
}
