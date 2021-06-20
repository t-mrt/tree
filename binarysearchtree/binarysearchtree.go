package binarysearchtree

type BinarySearchTreeNode struct {
	data  int
	left  *BinarySearchTreeNode
	right *BinarySearchTreeNode
}

func Insert(root *BinarySearchTreeNode, data int) *BinarySearchTreeNode {
	if root == nil {
		return &BinarySearchTreeNode{
			data:  data,
			left:  nil,
			right: nil,
		}
	}

	if data < root.data {
		root.left = Insert(root.left, data)
	} else {
		root.right = Insert(root.right, data)
	}
	return root
}

func Delete(root *BinarySearchTreeNode, data int) *BinarySearchTreeNode {
	if root == nil {
		return nil
	}

	if data < root.data {
		root.left = Delete(root.left, data)
	} else if data > root.data {
		root.right = Delete(root.right, data)
	} else {
		if root.left != nil && root.right != nil {
			m := FindMax(root.left)
			root.data = m
			root.left = Delete(root.left, root.data)
		} else {
			if root.left == nil {
				root = root.right
			} else {
				root = root.left
			}
		}
	}

	return root
}

func Find(root *BinarySearchTreeNode, target int) bool {
	if root == nil {
		return false
	}
	return (root.data == target) || Find(root.left, target) || Find(root.right, target)
}

func FindMax(root *BinarySearchTreeNode) int {

	if root.left == nil && root.right == nil {
		return root.data
	} else if root.left != nil && root.right == nil {
		return root.data
	} else if root.left == nil && root.right != nil {
		return FindMax(root.right)
	} else {
		return FindMax(root.right)
	}
}
