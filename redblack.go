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
	value  interface{}
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

// Put will put the value v with the associated key k. If k is already
// present, this will overwrite that value. This operation is
// considered idempotent.
func (t *tree) Put(k int, v interface{}) {
	n, p := descend(t.root, nil, k)
	newNode := node{color: red, key: k, parent: p, value: v}
	if p == nil {
		// Must be root
		t.root = &newNode
	} else if n != nil {
		// Value already present
		n.value = v
	} else {
		if k < p.key {
			p.left = &newNode
		} else {
			p.right = &newNode
		}
	}

	var z *node
	if n != nil {
		z = n
	} else {
		z = &newNode
	}
	fixupInsertion(t, z)
}

// Recursively descends down a tree looking for where to insert a
// given node. Returns the pointer to that location AND the parent of
// the given node
func descend(n *node, p *node, k int) (node *node, parent *node) {
	if n == nil {
		return n, p
	}
	if n.key == k {
		return n, nil
	} else if k < n.key {
		return descend(n.left, n, k)
	} else {
		return descend(n.right, n, k)
	}
}

func checkColor(n *node) color {
	if n == nil {
		return black
	}
	return n.color
}

// Called after performing a new insertion. This will restore balance
// to the tree and esnure the properties of the redblack tree are
// all true
func fixupInsertion(t *tree, n *node) {
	// When we begin this loop, the following conditions are true:
	// - The node is red
	// - If the parent is nil, it's color is black
	// - If the tree is violating any of the RBTree properties then it
	// can only be violating at most one of them.
	for checkColor(n.parent) == red {
		// Checks which side of the tree we're on
		if n.parent == n.parent.parent.left {
			// y is the uncle of node n, which may be nil
			y := n.parent.parent.right
			if checkColor(y) == red {
				// Occurs when both parent and uncle are red
				n.parent.color = black
				y.color = black
				n.parent.parent.color = red
				n = n.parent.parent
				// By setting the colors in this way, we ensure that
				// we'll go through another iteration of the loop. We
				// move the node pointer to it's grandfather. This
				// branch ensures that the no red children of the
				// subtree have red nodes as children
			} else {
				if n == n.parent.right {
					// Occurs when the uncle of n (y) is black, and n is
					// right child
					n = n.parent
					// Perform a rotation immediately, which is needed
					// for the next steps
					t.rotateLeft(n)
				}
				// Occurs when the uncle is black, and n is a left
				// child. If it wasn't a left chiild, we just made it one.
				n.parent.color = black
				n.parent.parent.color = red
				t.rotateRight(n.parent.parent)
				// The loop should terminate at this point, since n
				// will now be black
			}
		} else {
			// This is symetrical to above. The only difference is
			// that the directions right and left are swapped.
			y := n.parent.parent.left
			if checkColor(y) == red {
				n.parent.color = black
				y.color = black
				n.parent.parent.color = red
				n = n.parent.parent
			} else {
				if n == n.parent.left {
					n = n.parent
					t.rotateRight(n)
				}
			n.parent.color = black
			n.parent.parent.color = red
			t.rotateLeft(n.parent.parent)
			}
		}
	}
	t.root.color = black
}

func (t *tree) Get(k int) (v interface{}, ok bool) {
	n, _ := descend(t.root, nil, k)

	if n != nil {
		v = n.value
		ok = true
	} else {
		ok = false
	}
	return v, ok
}

func New() *tree {
	return &tree{root: nil}
}
