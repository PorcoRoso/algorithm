/*
	binary search tree.
	Node
	Tree: node & method
	Method: Insert, Search, Remove, Min,Max, InOrderTraverse,PreOrderTraverse,PostOrderTraverse, String
*/
package main

import (
	"sync"
	"fmt"
)

type Node struct {
	data int
	left *Node
	right *Node
}

type ItemBinarySearchTree struct {
	node *Node
	lock sync.Mutex
}

func (tree *ItemBinarySearchTree) Insert(data int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	newNode := &Node{data, nil, nil}

	if tree.node == nil {
		tree.node = newNode
	} else {
		insertNode(data, tree.node)
	}
}

func insertNode(data int, node *Node) {
	if data < node.data {
		if node.left == nil {
			node.left = &Node{data, nil, nil}
		} else {
			insertNode(data, node.left)
		}

	} else {
		if node.right == nil {
			node.right = &Node{data, nil, nil}
		} else {
			insertNode(data, node.right)
		}
	}
}

// 后序遍历打印树结构
func (tree *ItemBinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	if tree.node == nil {
		println("Tree is empty")
		return
	}
	stringify(tree.node, 0)
	println("----------------------------")
}
func stringify(node *Node, level int) {
	if node == nil {
		return
	}

	format := ""
	for i := 0; i < level; i++ {
		format += "\t" // 根据节点的深度决定缩进长度
	}
	format += "----[ "
	level++
	// 先递归打印左子树
	stringify(node.left, level)
	// 打印值
	fmt.Printf(format+"%d\n", node.data)
	/// 再递归打印右子树
	stringify(node.right, level)
}

func main() {
	tree := &ItemBinarySearchTree{}
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(10)
	tree.Insert(6)
	tree.Insert(2)
	tree.String()
}
