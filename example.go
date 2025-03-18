package main

import (
	"fmt"

	"github.com/chrissollitto/cps/iterator"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	exampleTree := &TreeNode{Val: 7, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}}

	dfsIterator := &iterator.CPSIterator[TreeNode, int]{}
	dfsIterator.Init(exampleTree,
		func(node *TreeNode, cont func() int) int {
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
