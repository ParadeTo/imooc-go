package main

import (
	. "./tree"
	"fmt"
)

type MyTreeNode struct {
	node *TreeNode
}

func (myNode *MyTreeNode) postOrder () {
	if myNode == nil || myNode.node == nil {
		return
	}
	(&MyTreeNode{myNode.node.Left}).postOrder()
	(&MyTreeNode{myNode.node.Right}).postOrder()
	myNode.node.Print()
}

func main() {

	var root TreeNode

	root = TreeNode{Value:3}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{5, nil, nil}
	root.Right.Left = new(TreeNode)

	//
	//t := &MyTreeNode{&root}
	//t.postOrder()
	//
	////root.print()
	////
	////n := treeNode{}
	////n.left.print()
	//root.Traverse()

	//q := queue.Queue{1}
	//q.Push("123")
	//fmt.Println(q.Pop())
	//fmt.Println(q.Pop())
	//fmt.Println(q.IsEmpty())

	//fmt.Println(q.Pop())

	//fmt.Println(q.Pop())

	// use channel to traverse
	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}

	a := make([]int, 5)
	fmt.Print(a[4])
}