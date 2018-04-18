package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func createNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func (node *TreeNode) Print() {
	if node == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println(node.Value)
}

func (node *TreeNode) SetValue(value int) {
	node.Value = value
}

func (node *TreeNode) Traverse() {
	node.TranverseFunc(func(node *TreeNode) {
		node.Print()
	})
}

func (node *TreeNode) TranverseFunc(f func(node *TreeNode)) {
	if node == nil {
		return
	}
	node.Left.Traverse()
	f(node)
	node.Right.Traverse()
}