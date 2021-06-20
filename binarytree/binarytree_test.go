package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateTree() *BinaryTreeNode {
	ts := []*BinaryTreeNode{}

	for i := 1; i <= 7; i++ {
		ts = append(ts, &BinaryTreeNode{data: i})
	}
	ts[0].left = ts[1]
	ts[0].right = ts[2]
	ts[1].left = ts[3]
	ts[1].right = ts[4]
	ts[2].left = ts[5]
	ts[2].right = ts[6]

	return ts[0]
}

func TestPreOder(t *testing.T) {

	root := CreateTree()

	assert.Equal(t, []int{1, 2, 4, 5, 3, 6, 7}, PreOrder(root))
}

func TestInOder(t *testing.T) {

	root := CreateTree()

	assert.Equal(t, []int{4, 2, 5, 1, 6, 3, 7}, InOrder(root))
}

func TestPostOder(t *testing.T) {

	root := CreateTree()

	assert.Equal(t, []int{4, 5, 2, 6, 7, 3, 1}, PostOrder(root))
}

func TestLevelOder(t *testing.T) {

	root := CreateTree()

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, LevelOrder(root))
}

func TestFindMax(t *testing.T) {

	root := CreateTree()
	max, e := FindMax(root)
	assert.Nil(t, e)
	assert.Equal(t, 7, max)
}

func TestSearchUsingLebelOder(t *testing.T) {

	root := CreateTree()
	for i := 1; i <= 7; i++ {
		assert.True(t, SearchUsingLebelOder(root, i))
	}

	assert.False(t, SearchUsingLebelOder(root, 0))
	assert.False(t, SearchUsingLebelOder(root, 8))
}
