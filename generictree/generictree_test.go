package generictree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateGenericTree() *GenericTreeNode {
	ts := []*GenericTreeNode{}

	for i := 1; i <= 7; i++ {
		ts = append(ts, &GenericTreeNode{data: i})
	}
	ts[0].nextSibling = ts[1]
	ts[1].nextSibling = ts[2]
	ts[2].firstChildren = ts[5]
	ts[2].nextSibling = ts[3]
	ts[3].nextSibling = ts[4]
	ts[4].firstChildren = ts[6]

	return ts[0]
}

func TestFindSum(t *testing.T) {

	root := CreateGenericTree()
	sum := FindSum(root)
	assert.Equal(t, 28, sum)
}
