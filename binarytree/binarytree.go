package binarytree

import "errors"

type BinaryTreeNode struct {
	data  int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func PreOrder(root *BinaryTreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{}
	ret = append(ret, root.data)

	ret = append(ret, PreOrder(root.left)...)
	ret = append(ret, PreOrder(root.right)...)

	return ret
}

func InOrder(root *BinaryTreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{}

	ret = append(ret, InOrder(root.left)...)
	ret = append(ret, root.data)
	ret = append(ret, InOrder(root.right)...)

	return ret
}

func PostOrder(root *BinaryTreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{}

	ret = append(ret, PostOrder(root.left)...)
	ret = append(ret, PostOrder(root.right)...)
	ret = append(ret, root.data)

	return ret
}

func LevelOrder(root *BinaryTreeNode) []int {
	q := []*BinaryTreeNode{}

	ret := []int{}

	if root == nil {
		return []int{}
	}
	q = append(q, root)

	for {
		temp := q[0]
		q = q[1:]

		ret = append(ret, temp.data)

		if temp.left != nil {
			q = append(q, temp.left)
		}

		if temp.right != nil {
			q = append(q, temp.right)
		}

		if len(q) == 0 {
			break
		}

	}
	return ret
}

func FindMax(root *BinaryTreeNode) (int, error) {
	if root == nil {
		return -1, errors.New("root is nil")
	}
	var max int
	l, err := FindMax(root.left)
	if err == nil {
		max = l
	}
	if root.data > max {
		max = root.data
	}
	r, err := FindMax(root.right)
	if err == nil && r > max {
		max = r
	}

	return max, nil
}

func SearchUsingLebelOder(root *BinaryTreeNode, target int) bool {
	q := []*BinaryTreeNode{}

	ret := false

	if root == nil {
		return ret
	}
	q = append(q, root)

	for {
		temp := q[0]
		q = q[1:]

		if temp.data == target {
			ret = true
			return ret
		}

		if temp.left != nil {
			q = append(q, temp.left)
		}

		if temp.right != nil {
			q = append(q, temp.right)
		}

		if len(q) == 0 {
			break
		}

	}
	return ret
}
