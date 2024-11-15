package dataStructure

import (
	"errors"
	"fmt"
)

type keyed interface {
	name() string
}

type BinaryNode[K keyed] struct {
	data  K              // to store the data item
	left  *BinaryNode[K] // pointer to point to left node
	right *BinaryNode[K] // pointer to point to right node
}

type BST[K keyed] struct {
	root *BinaryNode[K]
	size int
}

func NewBST[K keyed]() *BST[K] {
	tree := BST[K]{nil, 0}
	return &tree
}

func (bst *BST[K]) insertNode(t **BinaryNode[K], item K) error {
	if *t == nil {
		newNode := &BinaryNode[K]{
			data:  item,
			left:  nil,
			right: nil,
		}
		*t = newNode
		return nil
	}
	if item.name() < (*t).data.name() {
		bst.insertNode(&((*t).left), item)
	} else if item.name() > (*t).data.name() {
		bst.insertNode(&((*t).right), item)
	} else {
		// If item is already in the tree, replace item.
		return errors.New("item already in BST")
	}
	bst.size++
	return nil
}

func (bst *BST[K]) Insert(item K) error {
	// item := Item{name, x, y, s}
	err := bst.insertNode(&bst.root, item)
	bst.size++
	return err
}

func (bst *BST[K]) inOrderTraverse(t *BinaryNode[K]) {
	if t != nil {
		bst.inOrderTraverse(t.left)
		fmt.Println(t.data)
		bst.inOrderTraverse(t.right)
	}
}

func (bst *BST[K]) preOrderTraverseToSlice(t *BinaryNode[K], slice *[]BinaryNode[K]) {
	if t != nil {
		*slice = append(*slice, *t)
		bst.preOrderTraverseToSlice(t.left, slice)
		bst.preOrderTraverseToSlice(t.right, slice)
	}
}

func (bst *BST[K]) PreOrderTraverseToSlice() *[]BinaryNode[K] {
	inorder := make([]BinaryNode[K], 0, bst.findSize())
	bst.preOrderTraverseToSlice(bst.root, &inorder)
	return &inorder
}

func (bst *BST[K]) InOrder() {
	bst.inOrderTraverse(bst.root)
}

func (bst *BST[K]) preOrderTraverse(t *BinaryNode[K]) {
	if t != nil {
		fmt.Println(t.data)
		bst.preOrderTraverse(t.left)
		bst.preOrderTraverse(t.right)
	}
}

func (bst *BST[K]) PreOrder() {
	bst.preOrderTraverse(bst.root)
}

// func (bst *BST) getItem(t *BinaryNode) string {
// 	for t.right != nil {
// 		t = t.right
// 	}
// 	return t.item
// }

func (bst *BST[K]) searchFunction(t *BinaryNode[K], target string) (*BinaryNode[K], error) {
	// Search for item, if cannot find, return root with an error
	if t.data.name() == target {
		return t, nil
	} else if t.left == nil && t.right == nil {
		return bst.root, errors.New("did not find item")
	} else if target < t.data.name() {
		return bst.searchFunction(t.left, target)
	} else {
		return bst.searchFunction(t.right, target)
	}
}

func (bst *BST[K]) Search(target string) (keyed, error) {
	out, e := bst.searchFunction(bst.root, target)
	return out.data, e
}

func (bst *BST[K]) findSize() int {
	return bst.size
}

// Activity #4: Count no of nodes function for the tree
func (bst *BST[K]) NoOfNode() int {
	return bst.findSize()
}

// Activity #5: Remove function for the tree
func (bst *BST[K]) searchRemoveFunction(t **BinaryNode[K], target string) (*BinaryNode[K], bool, error) {
	if (*t) == nil {
		return nil, false, errors.New("Item not found, cannot be deleted")
	} else if (*t).data.name() == target {
		bst.size--
		// Case: no children
		if (*t).left == nil && (*t).right == nil {
			return nil, true, nil
			// Case: Left child only
		} else if (*t).left != nil && (*t).right == nil {
			return (*t).left, true, nil
			// Case: Right child only
		} else if (*t).left == nil && (*t).right != nil {
			return (*t).right, true, nil
			// Case: both children
			// plan is to go left then right all the way
		} else {
			successor := bst.findSuccessor(&(*t).left)
			return successor, true, nil
		}
	} else if target < (*t).data.name() {
		newChild, changes, e := bst.searchRemoveFunction(&(*t).left, target)
		if e == nil && changes {
			(*t).left = newChild
		}
		return nil, false, e
	} else if target > (*t).data.name() {
		newChild, changes, e := bst.searchRemoveFunction(&(*t).right, target)
		if e == nil && changes {
			(*t).right = newChild
		}
		return nil, false, e
	} else {
		return nil, false, errors.New("unknown error")
	}
}

func (bst *BST[K]) findSuccessor(t **BinaryNode[K]) *BinaryNode[K] {
	if (*t).right == nil {
		return *t
	} else {
		return bst.findSuccessor(&(*t).right)
	}
}

func (bst *BST[K]) Remove(target string) error {
	newChild, changes, e := bst.searchRemoveFunction(&bst.root, target)
	// To handle case where root is the target node
	if e == nil && changes {
		bst.root = newChild
	}
	return e
}
