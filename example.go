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

	// Inorder Traversal
	dfsInorderIter := &iterator.CPSIterator[TreeNodeInt, int]{}
	dfsInorderIter.Init(exampleTree,
		func(node *TreeNodeInt, cont func() int) int {
			return dfsInorderIter.Run(node.Left, func() int {
				return dfsInorderIter.Suspend(node.Val, func() int {
					return dfsInorderIter.Run(node.Right, cont)
				})
			})
		})
	fmt.Println("----- Inorder Traversal ------")
	for dfsInorderIter.HasNext() {
		fmt.Println(dfsInorderIter.Next())
	}

	// Preorder Traversal
	dfsPreorderIter := &iterator.CPSIterator[TreeNodeInt, int]{}
	dfsPreorderIter.Init(exampleTree,
		func(node *TreeNodeInt, cont func() int) int {
			return dfsPreorderIter.Suspend(node.Val, func() int {
				return dfsPreorderIter.Run(node.Left, func() int {
					return dfsPreorderIter.Run(node.Right, cont)
				})
			})
		})
	fmt.Println("----- Preorder Traversal ------")
	for dfsPreorderIter.HasNext() {
		fmt.Println(dfsPreorderIter.Next())
	}

	// Postorder Traversal
	dfsPostorderIter := &iterator.CPSIterator[TreeNodeInt, int]{}
	dfsPostorderIter.Init(exampleTree,
		func(node *TreeNodeInt, cont func() int) int {
			return dfsPostorderIter.Run(node.Left, func() int {
				return dfsPostorderIter.Run(node.Right, func() int {
					return dfsPostorderIter.Suspend(node.Val, cont)
				})
			})
		})
	fmt.Println("----- Postorder Traversal ------")
	for dfsPostorderIter.HasNext() {
		fmt.Println(dfsPostorderIter.Next())
	}
}
