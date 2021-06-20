package binarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {

	r1 := Insert(nil, 1)
	Insert(r1, 2)
	assert.Equal(t, 1, r1.data)
	assert.Nil(t, r1.left)
	assert.Equal(t, 2, r1.right.data)

	r2 := Insert(nil, 2)
	Insert(r2, 1)
	assert.Equal(t, 2, r2.data)
	assert.Nil(t, r2.right)
	assert.Equal(t, 1, r2.left.data)

	r3 := Insert(nil, 2)
	Insert(r2, 3)
	Insert(r2, 4)
	assert.Equal(t, 2, r3.data)
	assert.Equal(t, 3, r2.right.data)
	assert.Equal(t, 4, r2.right.right.data)

	r4 := Insert(nil, 1)
	Insert(r4, 1)
	Insert(r4, 1)
	assert.Equal(t, 1, r4.data)
	assert.Equal(t, 1, r4.right.data)
	assert.Equal(t, 1, r4.right.right.data)

}

func TestFind(t *testing.T) {

	r1 := Insert(nil, 1)
	Insert(r1, 2)
	Insert(r1, 6)
	Insert(r1, 8)

	assert.True(t, Find(r1, 1))
	assert.True(t, Find(r1, 2))
	assert.False(t, Find(r1, 3))
	assert.True(t, Find(r1, 6))
	assert.True(t, Find(r1, 8))

}

func TestFindMax(t *testing.T) {

	r1 := Insert(nil, 1)
	Insert(r1, 8)
	Insert(r1, 2)
	Insert(r1, 6)

	assert.Equal(t, 8, FindMax(r1))

}

func TestDelete(t *testing.T) {

	r1 := Insert(nil, 1)
	Insert(r1, 2)
	Insert(r1, 4)
	Delete(r1, 4)

	assert.Nil(t, r1.right.right)

	r2 := Insert(nil, 2)
	Insert(r2, 1)
	Insert(r2, 3)
	Delete(r2, 2)

	assert.Equal(t, 1, r2.data)
	assert.Nil(t, r2.left)
	assert.Equal(t, 3, r2.right.data)

}
