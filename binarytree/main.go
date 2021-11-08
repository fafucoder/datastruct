/**
二叉树
*/

package main

type BinaryTree struct {
	Root *BinaryTreeNode
}

type BinaryTreeNode struct {
	Value int
	Times int

	LeftNode  *BinaryTreeNode
	RightNode *BinaryTreeNode
}

// 添加树
func (t *BinaryTree) Add(value int) {
	if t.Root == nil {
		t.Root = &BinaryTreeNode{Value: value}
		return
	}

	t.Root.Add(value)
}

func (t *BinaryTree) FindMinTreeNode() *BinaryTreeNode {
	if t.Root == nil {
		return nil
	}

	return t.Root.FindMinTreeNode()
}

func (t *BinaryTree) FindMaxTreeNode() *BinaryTreeNode {
	if t.Root == nil {
		return nil
	}

	return t.Root.FindMaxTreeNode()
}

func (t *BinaryTree) Find(value int) *BinaryTreeNode {
	if t.Root == nil {
		return nil
	}

	return t.Root.Find(value)
}

// 添加树节点
func (b *BinaryTreeNode) Add(value int) {
	if value < b.Value {
		if b.LeftNode == nil {
			b.LeftNode = &BinaryTreeNode{
				Value: value,
			}
			return
		}

		b.LeftNode.Add(value)
	} else if value > b.Value {
		if b.RightNode == nil {
			b.RightNode = &BinaryTreeNode{
				Value: value,
			}
			return
		}

		b.RightNode.Add(value)
	} else {
		b.Times += 1
	}
}

func (b *BinaryTreeNode) FindMinTreeNode() *BinaryTreeNode {
	if b.LeftNode == nil {
		return b
	}

	return b.LeftNode.FindMinTreeNode()
}

func (b *BinaryTreeNode) FindMaxTreeNode() *BinaryTreeNode {
	if b.RightNode == nil {
		return b
	}

	return b.RightNode.FindMaxTreeNode()
}

func (b *BinaryTreeNode) Find(value int) *BinaryTreeNode {
	if b.Value == value {
		return b
	}
	if b.Value < value {
		if b.LeftNode == nil {
			return nil
		}
		return b.LeftNode.Find(value)
	}

	if b.RightNode == nil {
		return nil
	}

	return b.RightNode.Find(value)
}

func main() {

}
