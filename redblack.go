package redblack

type color bool

const (
	red   = false
	black = true
)

type node struct {
	color  color
	key    int
	parent *node
	left   *node
	right  *node
}

type tree struct {
	root *node
}

func (t *tree) rotateLeft(x *node) {
	var y, beta, parent *node
	y = x.right
	parent = x.parent

	beta = y.left
	if beta != nil {
		beta.parent = x
	}

	x.right = beta
	x.parent = y

	y.parent = parent
	y.left = x

	// Make sure to update the parent relationship
	if parent != nil {
		if parent.left == x {
			parent.left = y
		} else {
			parent.right = y
		}
	}

	// Don't forget to update the root of the tree
	if t.root == x {
		t.root = y
	}
}

func (t *tree) rotateRight(y *node) {
	var x, beta, parent *node
	x = y.left
	parent = y.parent

	beta = x.right
	if beta != nil {
		beta.parent = y
	}

	y.left = beta
	y.parent = x

	x.parent = parent
	x.right = y

	// Make sure to update the parent relationship
	if parent != nil {
		if parent.left == y {
			parent.left = x
		} else {
			parent.right = x
		}
	}
	// Don't forget to update the root of the tree
	if t.root == y {
		t.root = x
	}
}

// Does an in order traversal and collects all of the keys into a slice
func (t *tree) traverse() []int {
	keys := []int{}

	traverseHelper(t.root, &keys)
	return keys
}

func traverseHelper(n *node, keys *[]int) {
	if n != nil {
		traverseHelper(n.left, keys)
		*keys = append(*keys, n.key)
		traverseHelper(n.right, keys)
	}
}
