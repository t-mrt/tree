package bptree2

import "fmt"

const nonLeadNodeOrder = 7
const margin = 1

const nonLeadNodeSize = nonLeadNodeOrder + margin
const leadNodeSize = 5 + margin

type childNode interface {
	insert(key string) *keyChild
	find(key string) string
}

type keyChild struct {
	key   string
	child childNode
}

type keyData struct {
	key  string
	data string
}

type nonLeafNode struct {
	pairs [nonLeadNodeSize]*keyChild
}

type leafNode struct {
	pairs [leadNodeSize]*keyData
}

type BPTree struct {
	root childNode
}

func New() *BPTree {

	ln := &leafNode{
		pairs: [leadNodeSize]*keyData{},
	}

	p := [nonLeadNodeSize]*keyChild{}
	p[0] = &keyChild{
		key:   "",
		child: ln,
	}

	return &BPTree{
		root: &nonLeafNode{
			pairs: p,
		}}
}

func (bpt *BPTree) Insert(key string) {

	p := bpt.root.insert(key)

	if p != nil {
		n := &nonLeafNode{
			pairs: [nonLeadNodeSize]*keyChild{},
		}
		n.pairs[0] = &keyChild{
			key:   "",
			child: bpt.root,
		}
		n.pairs[1] = p
		bpt.root = n
	}
}

func (bpt *BPTree) Find(key string) string {

	return bpt.root.find(key)
}

func (n *nonLeafNode) insert(key string) *keyChild {

	var node childNode
	for _, v := range n.pairs {
		if v == nil {
			break
		}
		if key < v.key {
			break
		}
		node = v.child
	}
	p := node.insert(key)

	if p != nil {
		for i, v := range n.pairs {
			if v == nil {
				n.pairs[i] = p
				break
			}
			if key < v.key {
				copy(n.pairs[i:], n.pairs[i-1:])
				n.pairs[i] = p
				break
			}
		}
	}

	return n.split()
}

func (n *nonLeafNode) find(key string) string {

	var node childNode

	for i, v := range n.pairs {
		if i == 0 && v == nil {
			break
		}
		if v == nil {
			node = n.pairs[i-1].child
			break
		}
		if i == 0 {
			node = v.child
		}
		if key < v.key {
			node = n.pairs[i-1].child
			break
		}
	}

	return node.find(key)
}

func (n *nonLeafNode) split() *keyChild {

	for _, v := range n.pairs {
		if v == nil {
			return nil
		}
	}

	leftNode := n
	rightNode := &nonLeafNode{
		pairs: [nonLeadNodeSize]*keyChild{},
	}

	p := &keyChild{
		key:   n.pairs[len(n.pairs)/2].key,
		child: rightNode,
	}

	copy(rightNode.pairs[:], leftNode.pairs[nonLeadNodeSize/2:])
	rightNode.pairs[0].key = ""

	zero := [nonLeadNodeSize]*keyChild{}
	for i := 0; i < nonLeadNodeSize; i++ {
		zero[i] = nil
	}

	copy(leftNode.pairs[nonLeadNodeSize/2:], zero[:])

	return p
}

func (n *leafNode) insert(key string) *keyChild {
	for i, v := range n.pairs {
		// first insert
		if v == nil {
			n.pairs[i] = &keyData{
				key:  key,
				data: fmt.Sprintf("d-%s", key),
			}
			break
		}
		if key < v.key {
			copy(n.pairs[i+1:], n.pairs[i:])
			n.pairs[i] = &keyData{
				key:  key,
				data: fmt.Sprintf("d-%s", key),
			}
			break
		}
	}
	return n.split()
}

func (n *leafNode) split() *keyChild {

	for _, v := range n.pairs {
		if v == nil {
			return nil
		}
	}

	leftNode := n
	rightNode := &leafNode{
		pairs: [leadNodeSize]*keyData{},
	}

	p := &keyChild{
		key:   n.pairs[nonLeadNodeSize/2-1].key,
		child: rightNode,
	}

	copy(rightNode.pairs[:], leftNode.pairs[nonLeadNodeSize/2-1:])
	zero := [5]*keyData{
		nil, nil, nil, nil, nil,
	}
	copy(leftNode.pairs[nonLeadNodeSize/2-1:], zero[:])

	return p
}

func (n *leafNode) find(key string) string {

	for _, v := range n.pairs {
		if v == nil {
			return ""
		}
		if v.key == key {
			return v.data
		}
	}

	return ""
}
