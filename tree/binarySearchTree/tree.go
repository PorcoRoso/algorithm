// 树：满树，完全二叉树，节点高度、节点深度、层，树的高度
// 树的遍历：前中后序遍历、层次遍历; 递归思想，递：A问题可以分为BC子问题，归：BC问题已解决，A问题怎么解决，终止条件！
// 树的复杂度：
package main

import "fmt"

type TreeNode struct {
	Data int
	Left *TreeNode
	Right *TreeNode
}

var arr []int

var res [][]int

// PreOrder leetcode 94
// Recursive
// 对比preOrder和preOrder1，前者是传引用（golang式写法），后者是传指针（C语言式写法）
// 推崇前者golang式写法，通过函数返回值的方式 https://blog.csdn.net/benben_2015/article/details/80884537
func preOrder(root *TreeNode) []int {
	if root != nil {
		preOrder(root.Left)
		arr = append(arr, root.Data)
		preOrder(root.Right)
	}

	return arr
}

func preOrder1(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	preOrder1(root.Left, arr)
	*arr = append(*arr, root.Data)
	preOrder1(root.Right, arr)
}

func levelOrder(root *TreeNode, level int) [][]int{
	if root != nil {

		if len(res) == level {
			var tmp []int
			res = append(res, tmp)
		}

		res[level] = append(res[level], root.Data)
		levelOrder(root.Left, level+1)
		levelOrder(root.Right, level+1)
	}
	return res
}

func main() {
	n3 := &TreeNode{
		3,
		nil,
		nil,
	}
	n2 := &TreeNode{
		2,
		n3,
		nil,
	}
	n1 := &TreeNode{
		1,
		nil,
		n2,
	}
	// 定义空切片
	//var arr = []int{}
	preOrder(n1)
	fmt.Println(arr[:])

	levelOrder(n1, 0)
	fmt.Println(res)
	res = nil
}




