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

func (t *tree) transplant(u *node, v *node) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v !=nil {
		v.parent = u.parent
	}
}

func minimum(n *node) *node {
	if n.left == nil {
		return n
	}
	return minimum(n.left)
}

func (t *tree) Delete(k int) {
	// first find the node with the given key
	n, _ := descend(t.root, nil, k)
	p := n.parent

	y := n
	yColorOriginal := y.color
	var x *node
	if n.left == nil {
		x = n.right
		t.transplant(n, n.right)
	} else if n.right == nil {
		x = n.left
		t.transplant(n, n.left)
	} else {
		y = minimum(n.right)
		yColorOriginal = y.color
		x = y.right
		if y.parent == n {
			x.parent = y
		} else {
			t.transplant(n, y)
			y.right = n.right
			y.right.parent = y
		}
		t.transplant(n, y)
		y.left = n.left
		y.left.parent = y
		y.color = n.color
	}
	if yColorOriginal == black {
		t.fixupDeletion(x, p)
	}
}

func (t *tree) fixupDeletion(x *node, p *node) {

	for x != t.root && checkColor(x) == black {
		if x == p.left {
			w := p.right
			if checkColor(w) == red {
				w.color = black
				p.color = red
				t.rotateLeft(p)
				w = p.right
			}
			if checkColor(w.left) == black && checkColor(w.right) == black {
				w.color = red
				x = p
				p = x.parent
			} else {
				if checkColor(w.right) == black {
					w.left.color = black
					w.color = red
					t.rotateRight(w)
					w = p.right
				}
				w.color = p.color
				p.color = black
				w.right.color = black
				t.rotateLeft(p)
				x = t.root
			}
		} else {
			w := p.left
			if checkColor(w) == red {
				w.color = black
				p.color = red
				t.rotateRight(p)
				w = p.left
			}
			if checkColor(w.right) == black && checkColor(w.left) == black {
				w.color = red
				x = p
				p = x.parent
			} else {
				if checkColor(w.left) == black {
					w.right.color = black
					w.color = red
					t.rotateLeft(w)
					w = p.left
				}
				w.color = p.color
				p.color = black
				w.left.color = black
				t.rotateRight(p)
				x = t.root
			}
		}
	}
	if x != nil {
		x.color = black
	}
}
