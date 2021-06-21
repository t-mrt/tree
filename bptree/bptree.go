package bptree

import (
	"errors"
)

// order
const leafNodeOrder = 3
const NonLeafNodeOrder = leafNodeOrder + 1

type BP struct {
	root  interface{}
	heght int
}

type BPNonLeafNode struct {
	key      [NonLeafNodeOrder]string
	children [NonLeafNodeOrder]interface{} // [](*BPNonLeafNode or *BPLeafNode or nil)
}

// maybe inefficiency
func (bp *BPNonLeafNode) Length() int {
	ret := 0
	for _, v := range bp.key {
		if v != "" {
			ret++
		}
	}
	return ret + 1
}

type BPLeafNode struct {
	key  [leafNodeOrder]string
	data [leafNodeOrder]string
	next *BPLeafNode
}

func (bp *BPLeafNode) Length() int {
	ret := 0
	for _, v := range bp.key {
		if v != "" {
			ret++
		}
	}
	return ret
}

func NewBP() *BP {

	bpln := &BPLeafNode{
		key:  [leafNodeOrder]string{},
		data: [leafNodeOrder]string{},
		next: nil,
	}

	return &BP{
		heght: 1,
		root:  bpln,
	}
}

func (bp *BP) Insert(key string, data string) {

	if key == "" {
		panic("key must not be empty string")
	}

	var ancestor []*BPNonLeafNode
	var insertedLeafNode *BPLeafNode
	node := bp.root

	for {

		if bpln, ok := node.(*BPLeafNode); ok {
			insertedLeafNode = bpln

			index := 0
			for {
				if bpln.Length() <= index {
					break
				}
				if bpln.key[index] > key {
					copy(bpln.key[index+1:], bpln.key[index:])
					copy(bpln.data[index+1:], bpln.data[index:])
					break
				}
				index++
			}
			bpln.key[index] = key
			bpln.data[index] = data

			break
		} else if bpnln, ok := node.(*BPNonLeafNode); ok {
			ancestor = append(ancestor, bpnln)
			index := 0
			for {
				if bpnln.Length() <= index || key < bpnln.key[index] {
					break
				}
				index++
			}
			node = bpnln.children[index-1]
		} else {
			panic("impossible")
		}
	}

	bp.rearrange(ancestor, insertedLeafNode, key)

}

func (bp *BP) rearrange(ancestor []*BPNonLeafNode, insertedLeafNode *BPLeafNode, key string) {

	if len(ancestor) > 0 {
		parent := ancestor[len(ancestor)-1]
		bp.rearrangeLeafNode(insertedLeafNode, parent)
	} else {
		bp.rearrangeLeafNode(insertedLeafNode, nil)
		return
	}

	for i := len(ancestor) - 1; 0 <= i; i-- {
		if ancestor[i].Length() == NonLeafNodeOrder {
			var parent *BPNonLeafNode
			if i != 0 {
				parent = ancestor[i-1]
			}
			bp.rearrangeNonLeafNode(ancestor[i], parent)
		}
	}
}

func (bp *BP) rearrangeLeafNode(leafNode *BPLeafNode, parentNode *BPNonLeafNode) {
	if leafNode.Length() < leafNodeOrder {
		return
	}

	climbKey := leafNode.key[1]
	leftBpln := leafNode
	rightBpln := &BPLeafNode{
		key:  [leafNodeOrder]string{},
		data: [leafNodeOrder]string{},
	}
	copy(rightBpln.key[:], leftBpln.key[1:leafNodeOrder])
	copy(rightBpln.data[:], leftBpln.data[1:leafNodeOrder])

	for _, v := range []int{2} {
		rightBpln.key[v] = ""
		rightBpln.data[v] = ""
	}

	for _, v := range []int{1, 2} {
		leftBpln.data[v] = ""
		leftBpln.key[v] = ""
	}

	rightBpln.next = leafNode.next
	leafNode.next = rightBpln

	if parentNode == nil {
		newBpnln := &BPNonLeafNode{
			key:      [NonLeafNodeOrder]string{},
			children: [NonLeafNodeOrder]interface{}{},
		}

		newBpnln.key[0] = ""
		newBpnln.key[1] = rightBpln.key[0]

		newBpnln.children[0] = leftBpln
		newBpnln.children[1] = rightBpln

		bp.root = newBpnln
		bp.heght++
	} else {
		parentBpnln := parentNode
		insertIndex := 0
		for {
			if parentBpnln.Length() <= insertIndex {
				break
			}
			if parentBpnln.key[insertIndex] > climbKey {
				copy(parentBpnln.key[insertIndex:], parentBpnln.key[insertIndex-1:])
				copy(parentBpnln.children[insertIndex:], parentBpnln.children[insertIndex-1:])
				break
			}
			insertIndex++
		}
		parentBpnln.key[insertIndex] = climbKey
		parentBpnln.children[insertIndex] = rightBpln
	}
}

func (bp *BP) rearrangeNonLeafNode(nonLeafNode *BPNonLeafNode, parentNode *BPNonLeafNode) {

	// split nonLeafNode into leftBpnln + rightBpnln
	// (reuse nonLeafNode for rightBpnln)

	rightBpnln := nonLeafNode
	climbKey := rightBpnln.key[NonLeafNodeOrder/2]

	leftBpnln := &BPNonLeafNode{
		key:      [NonLeafNodeOrder]string{},
		children: [NonLeafNodeOrder]interface{}{},
	}

	copy(leftBpnln.key[:], rightBpnln.key[:2])
	copy(leftBpnln.children[:], rightBpnln.children[:2])

	copy(rightBpnln.key[1:], rightBpnln.key[3:4]) // index=0 is "" key
	copy(rightBpnln.children[:], rightBpnln.children[2:4])

	for _, v := range []int{2, 3} {
		rightBpnln.key[v] = ""
		rightBpnln.children[v] = nil
	}

	if parentNode == nil {
		root := &BPNonLeafNode{
			key:      [NonLeafNodeOrder]string{},
			children: [NonLeafNodeOrder]interface{}{},
		}

		root.key[0] = ""
		root.key[1] = climbKey

		root.children[0] = leftBpnln
		root.children[1] = rightBpnln

		bp.root = root
		bp.heght++
	} else {
		parentBpnln := parentNode
		insertIndex := 0
		for {
			if parentBpnln.Length() <= insertIndex {
				break
			}
			if parentBpnln.key[insertIndex] > climbKey {
				copy(parentBpnln.key[insertIndex:], parentBpnln.key[insertIndex-1:])
				copy(parentBpnln.children[insertIndex:], parentBpnln.children[insertIndex-1:])
				break
			}
			insertIndex++
		}
		parentBpnln.key[insertIndex] = climbKey
		parentBpnln.children[insertIndex-1] = leftBpnln
		parentBpnln.children[insertIndex] = rightBpnln
	}
}

func (bp *BP) Find(key string) (string, error) {

	node := bp.root
	for {
		if bpln, ok := node.(*BPLeafNode); ok {
			for i := 0; i < bpln.Length(); i++ {
				if bpln.key[i] == key {
					return bpln.data[i], nil
				}
			}
			return "", errors.New("not found")
		} else if bpnln, ok := node.(*BPNonLeafNode); ok {
			for i := 0; i < bpnln.Length(); i++ {
				if i != 0 && key < bpnln.key[i] {
					break
				}
				node = bpnln.children[i]
			}
		} else {
			panic("impossible")
		}
	}
}
