package main

import (
	"fmt"

	"github.com/chrissollitto/cps/iterator"
)

type TreeNode[T any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type TreeNodeInt = TreeNode[int]

func main() {
	exampleTree := &TreeNodeInt{Val: 7, Left: &TreeNodeInt{Val: 3}, Right: &TreeNodeInt{Val: 15, Left: &TreeNodeInt{Val: 9}, Right: &TreeNodeInt{Val: 20}}}

	dfsIterator := &iterator.CPSIterator[TreeNodeInt, int]{}
	dfsIterator.Init(exampleTree,
		func(node *TreeNodeInt, cont func() int) int {
			return dfsIterator.Run(node.Left, func() int {
				return dfsIterator.Suspend(node.Val, func() int {
					return dfsIterator.Run(node.Right, cont)
				})
			})
		})

	for dfsIterator.HasNext() {
		fmt.Println(dfsIterator.Next())
	}
}
