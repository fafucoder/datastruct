/**
栈: 先进后出
*/

package main

import (
	"fmt"
	"sync"
)

type LinkStack struct {
	root *LinkNode
	size int

	lock sync.Mutex
}
type LinkNode struct {
	Data string
	Next *LinkNode
}

func (l *LinkStack) Push(value string) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.root == nil {
		l.root = &LinkNode{Data: value}
		l.size += 1
		return
	}

	linkNode := &LinkNode{
		Data: value,
		Next: l.root,
	}

	l.root = linkNode
	l.size += 1
}

func (l *LinkStack) Pop() string {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.root == nil {
		return ""
	}

	topNode := l.root
	l.root = topNode.Next
	l.size = l.size - 1

	return topNode.Data
}

func (l *LinkStack) Size() int {
	return l.size
}

func main() {
	linkStack := new(LinkStack)
	linkStack.Push("cat")
	linkStack.Push("dog")
	linkStack.Push("hen")
	fmt.Println("size:", linkStack.Size())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("size:", linkStack.Size())
}
