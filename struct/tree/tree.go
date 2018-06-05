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
	node.TraverseFunc(func(node *TreeNode) {
		node.Print()
	})
}

func (node *TreeNode) TraverseFunc(f func(node *TreeNode)) {
	if node == nil {
		return
	}
	node.Left.Traverse()
	f(node)
	node.Right.Traverse()
}

func (node *TreeNode) TraverseWithChannel() chan *TreeNode {
	out := make(chan *TreeNode)
	go func() {
		node.TraverseFunc(func(node *TreeNode) {
			out <- node
		})
		close(out)
	}()
	return out
}