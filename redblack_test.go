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
