package generictree

type GenericTreeNode struct {
	data          int
	firstChildren *GenericTreeNode
	nextSibling   *GenericTreeNode
}

func FindSum(root *GenericTreeNode) int {
	if root == nil {
		return 0
	}

	return root.data + FindSum(root.firstChildren) + FindSum(root.nextSibling)

}
