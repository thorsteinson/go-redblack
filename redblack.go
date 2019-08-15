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
	var y, alpha, beta, gamma *node
	y = x.right

	alpha = x.left
	beta = y.left
	gamma = y.right

	x.left = alpha
	x.right = beta
	y.left = x
	y.right = gamma

	y.parent = x.parent
	x.parent = y

	// Don't forget to update the root of the tree
	if t.root == x {
		t.root = y
	}
}

func (t *tree) rotateRight(y *node) {
	var x, alpha, beta, gamma *node
	x = y.left

	alpha = x.left
	beta = x.right
	gamma = y.right

	x.left = alpha
	x.right = y
	y.left = beta
	y.right = gamma

	x.parent = y.parent
	y.parent = x

	// Don't forget to update the root of the tree
	if t.root == y {
		t.root = x
	}
}

// Does an in order traversal and collects all of the nodes into a slice
func (t *tree) traverse() []*node {
	nodes := []*node{}

	traverseHelper(t.root, &nodes)
	return nodes
}

func traverseHelper(n *node, nodeList *[]*node) {
	if n != nil {
		traverseHelper(n.left, nodeList)
		*nodeList = append(*nodeList, n)
		traverseHelper(n.right, nodeList)
	}
}
