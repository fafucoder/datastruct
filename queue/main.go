/**
队列：先进先出
*/

package main

import (
	"fmt"
	"sync"
)

type LinkQueue struct {
	root *LinkNode
	size int

	lock sync.Mutex
}

type LinkNode struct {
	Data string
	Next *LinkNode
}

func (q *LinkQueue) Push(value string) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.root == nil {
		q.root = &LinkNode{Data: value}
		q.size += 1
		return
	}

	node := new(LinkNode)
	node.Data = value

	currentNode := q.root
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = node

	q.size += 1
}

func (q *LinkQueue) Pop() string {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.size == 0 {
		return ""
	}

	currentNode := q.root
	q.root = currentNode.Next
	q.size = q.size - 1

	return currentNode.Data
}

func (q *LinkQueue) Size() int {
	return q.size
}

func main() {
	linkQueue := new(LinkQueue)
	linkQueue.Push("cat")
	linkQueue.Push("dog")
	linkQueue.Push("hen")

	fmt.Println("size:", linkQueue.Size())
	fmt.Println("pop:", linkQueue.Pop())
	fmt.Println("pop:", linkQueue.Pop())
	fmt.Println("size:", linkQueue.Size())
}
