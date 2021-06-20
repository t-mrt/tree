package avltree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	r1 := Insert(nil, 5)
	assert.NotNil(t, r1)

	r1 = Insert(r1, 4)
	r1 = Insert(r1, 3)
	assert.Equal(t, 2, r1.height)
	assert.Equal(t, 4, r1.data)
	assert.Equal(t, 3, r1.left.data)
	assert.Equal(t, 5, r1.right.data)

	r1 = Insert(r1, 6)
	assert.Equal(t, 3, r1.height)
	assert.Equal(t, 4, r1.data)
	assert.Equal(t, 3, r1.left.data)
	assert.Equal(t, 5, r1.right.data)
	assert.Equal(t, 6, r1.right.right.data)

	r1 = Insert(r1, 7)
	assert.Equal(t, 3, r1.height)
	assert.Equal(t, 4, r1.data)
	assert.Equal(t, 3, r1.left.data)
	assert.Equal(t, 6, r1.right.data)
	assert.Equal(t, 5, r1.right.left.data)
	assert.Equal(t, 7, r1.right.right.data)
}
