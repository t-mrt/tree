package avltree

import "math"

type AVLTreeNode struct {
	data   int
	left   *AVLTreeNode
	right  *AVLTreeNode
	height int
}

func max(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func height(n *AVLTreeNode) int {
	if n == nil {
		return 0
	}
	return n.height
}

func SingleRotateRight(y *AVLTreeNode) *AVLTreeNode {
	x := y.left
	t := x.right

	x.right = y
	y.left = t

	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1

	return x
}

func SingleRotateLeft(x *AVLTreeNode) *AVLTreeNode {
	y := x.right
	t := y.left

	y.left = x
	x.right = t

	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1

	return y
}

func getBalance(root *AVLTreeNode) int {
	if root == nil {
		return 0
	}
	lh := 0
	if root.left != nil {
		lh = root.left.height
	}
	rh := 0
	if root.right != nil {
		rh = root.right.height
	}
	return lh - rh
}

func Insert(root *AVLTreeNode, data int) *AVLTreeNode {

	if root == nil {
		return &AVLTreeNode{
			data:   data,
			left:   nil,
			right:  nil,
			height: 1,
		}
	}

	if data < root.data {
		root.left = Insert(root.left, data)
	} else if data > root.data {
		root.right = Insert(root.right, data)
	} else {
		return root
	}
	lh := 0
	if root.left != nil {
		lh = root.left.height
	}
	rh := 0
	if root.right != nil {
		rh = root.right.height
	}

	root.height = max(lh, rh) + 1

	b := getBalance(root)

	if b > 1 {
		if data < root.left.data {
			return SingleRotateRight(root)
		} else if data > root.left.data {
			root.left = SingleRotateLeft(root.left)
			return SingleRotateRight(root)
		}
	}

	if b < -1 {
		if data > root.right.data {
			return SingleRotateLeft(root)
		} else if data < root.right.data {
			root.right = SingleRotateRight(root.right)
			return SingleRotateLeft(root)
		}
	}
	return root
}
