package bptree2

import (
	"fmt"
	"sync"
)

const nonLeadNodeOrder = 7
const margin = 1

const nonLeadNodeSize = nonLeadNodeOrder + margin
const leadNodeSize = 5 + margin

type childNode interface {
	Insert(key string) *keyChild
	Find(key string) string
	NeedLock() bool // this method is for parent
	RLock()
	RUnlock()
	Lock()
	Unlock()
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
	mu    sync.RWMutex
}

type leafNode struct {
	pairs [leadNodeSize]*keyData
	mu    sync.RWMutex
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

	bpt.root.Lock()
	p := bpt.root.Insert(key)

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
	bpt.root.RLock()

	return bpt.root.Find(key)
}

func (n *nonLeafNode) Insert(key string) *keyChild {

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

	node.Lock()

	if node.NeedLock() {
		defer n.Unlock()
	} else {
		// technically, we need to release all the locks on the ancestors.
		n.Unlock()
	}

	p := node.Insert(key)

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

func (n *nonLeafNode) Find(key string) string {

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
	node.RLock()
	n.RUnlock()

	return node.Find(key)
}

func (n *nonLeafNode) split() *keyChild {

	if n.needSplit() {
		return nil
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

func (n *nonLeafNode) needSplit() bool {
	for _, v := range n.pairs {
		if v == nil {
			return true
		}
	}
	return false
}

func (n *nonLeafNode) NeedLock() bool {

	oc := 0
	for _, v := range n.pairs {
		if v == nil {
			break
		}
		oc++
	}
	return oc == nonLeadNodeSize-1
}

func (n *leafNode) Find(key string) string {

	defer n.RUnlock()

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

func (n *leafNode) Insert(key string) *keyChild {

	defer n.Unlock()

	for i, v := range n.pairs {
		// first Insert
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

	if n.needSplit() {
		return nil
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

func (n *leafNode) needSplit() bool {
	for _, v := range n.pairs {
		if v == nil {
			return true
		}
	}
	return false
}

func (n *leafNode) NeedLock() bool {

	oc := 0
	for _, v := range n.pairs {
		if v == nil {
			break
		}
		oc++
	}
	return oc == leadNodeSize-1
}

func (n *leafNode) RLock() {
	n.mu.RLock()
}
func (n *leafNode) Lock() {
	n.mu.Lock()
}
func (n *leafNode) RUnlock() {
	n.mu.RUnlock()
}
func (n *leafNode) Unlock() {
	n.mu.Unlock()
}

func (n *nonLeafNode) RLock() {
	n.mu.RLock()
}
func (n *nonLeafNode) Lock() {
	n.mu.Lock()
}
func (n *nonLeafNode) RUnlock() {
	n.mu.RUnlock()
}
func (n *nonLeafNode) Unlock() {
	n.mu.Unlock()
}
