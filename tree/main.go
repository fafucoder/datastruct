/**
树
*/

package main

import "fmt"

type TreeNode struct {
	Data      string
	LeftNode  *TreeNode
	RightNode *TreeNode
}

// 先续遍历，根左右
func (t *TreeNode) PreOrder() {
	if t == nil {
		return
	}

	fmt.Print(t.Data, " ")
	if t.LeftNode != nil {
		t.LeftNode.PreOrder()
	}

	if t.RightNode != nil {
		t.RightNode.PreOrder()
	}
}

// 中续遍历 左根右
func (t *TreeNode) MinOrder() {
	if t == nil {
		return
	}

	if t.LeftNode != nil {
		t.LeftNode.MinOrder()
	}

	fmt.Print(t.Data, " ")

	if t.RightNode != nil {
		t.RightNode.MinOrder()
	}
}

// 后续遍历 左右根
func (t *TreeNode) AfterOrder() {
	if t == nil {
		return
	}

	if t.LeftNode != nil {
		t.LeftNode.MinOrder()
	}

	if t.RightNode != nil {
		t.RightNode.MinOrder()
	}

	fmt.Print(t.Data, " ")
}

func main() {
	t := &TreeNode{Data: "A"}
	t.LeftNode = &TreeNode{Data: "B"}
	t.RightNode = &TreeNode{Data: "C"}
	t.LeftNode.LeftNode = &TreeNode{Data: "D"}
	t.LeftNode.RightNode = &TreeNode{Data: "E"}
	t.RightNode.LeftNode = &TreeNode{Data: "F"}
	t.RightNode.RightNode = &TreeNode{Data: "G"}

	fmt.Println("先序排序：")
	t.PreOrder()
	fmt.Println("\n中序排序：")
	t.MinOrder()
	fmt.Println("\n后序排序")
	t.AfterOrder()
}
