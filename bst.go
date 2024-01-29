package bst

import (
	"errors"
	"fmt"
	"sync"
)

type Node[T any] struct {
	key   uint64
	data  T
	left  *Node[T]
	right *Node[T]
}

type Tree[T any] struct {
	root  *Node[T]
	count int
	mu    *sync.RWMutex
	locks bool
}

// by default tree is locked, use the unlook method to disable the mutex
func New[T any]() *Tree[T] {
	return &Tree[T]{
		mu:    new(sync.RWMutex),
		locks: true}
}

func (tr *Tree[T]) Lock() *Tree[T] {
	tr.locks = true
	tr.mu = new(sync.RWMutex)
	return tr
}

func (tr *Tree[T]) Unlock() *Tree[T] {
	tr.locks = false
	tr.mu = nil
	return tr
}

func (tr *Tree[T]) Insert(key uint64, data T) error {

	if _, found := tr.Search(key); !found {

		if !tr.locks {
			tr.root = tr.insertByNode(tr.root, key, data)
			tr.count++
			return nil
		}

		tr.mu.Lock()

		tr.root = tr.insertByNode(tr.root, key, data)
		tr.count++

		tr.mu.Unlock()
		return nil
	} else {
		return errors.New("repeated key")
	}

}

func (tr *Tree[T]) insertByNode( /*root*/ node *Node[T], key uint64, data T) *Node[T] {
	if node == nil {
		return &Node[T]{key: key,
			data: data,
		}
	}

	if node.key > key {
		node.left = tr.insertByNode(node.left, key, data)
	}

	if node.key < key {
		node.right = tr.insertByNode(node.right, key, data)
	}

	return node
}

func (tr *Tree[T]) InOrderTraversalByNode(node *Node[T]) {
	if node == nil {
		return
	}

	tr.InOrderTraversalByNode(node.left)
	fmt.Println("key:", node.key, "   ", "value:", node.data)
	tr.InOrderTraversalByNode(node.right)

}

func (tr *Tree[T]) Search(key uint64) (*Node[T], bool) {

	return tr.searchByNode(tr.root, key)

}

func (tr *Tree[T]) searchByNode(node *Node[T], key uint64) (*Node[T], bool) {

	if node == nil {
		return nil, false
	}

	if node.key > key {
		return tr.searchByNode(node.left, key)

	}

	if node.key < key {
		return tr.searchByNode(node.right, key)

	}

	return node, true /*
		equvalent to if n.key == key {return n, true}
	*/

}

// Remove returns the Root, root == the whole tree cause we can access everything using the root
func (tr *Tree[T]) Remove(key uint64) (*Node[T], error) {

	if _, found := tr.Search(key); found {
		if !tr.locks {
			root := tr.removeByNode(tr.root, key)
			tr.count--
			return root, nil
		}

		tr.mu.Lock()

		root := tr.removeByNode(tr.root, key)
		tr.count--

		tr.mu.Unlock()
		return root, nil
	} else {
		return nil, errors.New("non-existent key")
	}
}

func (tr *Tree[T]) removeByNode(node *Node[T], key uint64) *Node[T] {
	if node == nil {
		return nil
	}

	if key > node.key {
		node.right = tr.removeByNode(node.right, key)
	}

	if key < node.key {
		node.left = tr.removeByNode(node.left, key)
	}

	if key == node.key {
		if node.IsLeaf() {
			return nil
		}

		if node.left == nil {
			return node.right

		} else if node.right == nil {
			return node.left
		} else /*if none of the right and left nodes are nil*/ {
			minKey := node.right.Min() //find the min key of the right sub-tree
			node.key = minKey
			temp, _ := tr.Search(minKey)
			node.data = temp.data
			node.right = tr.removeByNode(node.right, minKey)
		}
	}

	return node

}

func (n *Node[T]) IsLeaf() bool {

	return n.left == nil && n.right == nil

}

// minimum key in a sub-tree
func (n *Node[T]) Min() uint64 {
	if n.left == nil {
		return n.key
	}

	return n.left.Min()

}

func (n *Node[T]) Max() uint64 {
	if n.right == nil {
		return n.key
	}

	return n.right.Max()

}

/*
	func (n Node[T]) MinALT() uint64 {
		for n.left != nil {
			n.left = n.left.left
		}
		return n.key
	}
*/
/*
	func (n *Node[T]) MaxALT() uint64 {
		if n.right != nil {
			n.right.Min()
		}

		return n.key

}
*/
