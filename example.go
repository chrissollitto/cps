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

type Fib struct {
	n1, n2 int
}

type Nbonacci struct {
	n []int
}

func main() {
	exampleTree := &TreeNodeInt{Val: 7, Left: &TreeNodeInt{Val: 3}, Right: &TreeNodeInt{Val: 15, Left: &TreeNodeInt{Val: 9}, Right: &TreeNodeInt{Val: 20}}}

	fmt.Println("----- Inorder Traversal ------")

	dfsInorderIter := &iterator.CPSIterator[TreeNodeInt, int]{}
	dfsInorderIter.Init(exampleTree,
		func(node *TreeNodeInt, cont func() int) int {
			return dfsInorderIter.Run(node.Left, func() int {
				return dfsInorderIter.Suspend(node.Val, func() int {
					return dfsInorderIter.Run(node.Right, cont)
				})
			})
		})
	for dfsInorderIter.HasNext() {
		fmt.Println(dfsInorderIter.Next())
	}

	fmt.Println("----- Preorder Traversal ------")

	dfsPreorderIter := &iterator.CPSIterator[TreeNodeInt, int]{}
	dfsPreorderIter.Init(exampleTree,
		func(node *TreeNodeInt, cont func() int) int {
			return dfsPreorderIter.Suspend(node.Val, func() int {
				return dfsPreorderIter.Run(node.Left, func() int {
					return dfsPreorderIter.Run(node.Right, cont)
				})
			})
		})
	for dfsPreorderIter.HasNext() {
		fmt.Println(dfsPreorderIter.Next())
	}

	fmt.Println("----- Postorder Traversal ------")

	dfsPostorderIter := &iterator.CPSIterator[TreeNodeInt, int]{}
	dfsPostorderIter.Init(exampleTree,
		func(node *TreeNodeInt, cont func() int) int {
			return dfsPostorderIter.Run(node.Left, func() int {
				return dfsPostorderIter.Run(node.Right, func() int {
					return dfsPostorderIter.Suspend(node.Val, cont)
				})
			})
		})
	for dfsPostorderIter.HasNext() {
		fmt.Println(dfsPostorderIter.Next())
	}

	fmt.Println("----- Fibonacci Sequence ------")
	fibIter := &iterator.CPSIterator[Fib, int]{}
	fibIter.Init(&Fib{0, 1},
		func(f *Fib, cont func() int) int {
			return fibIter.Suspend(f.n1, func() int {
				return fibIter.Run(&Fib{f.n2, f.n1 + f.n2}, cont)
			})
		})

	for i := 0; i < 20; i++ {
		fmt.Println(fibIter.Next())
	}

	fmt.Println("----- Nbonacci Sequence ------")
	nbonacciIter := &iterator.CPSIterator[Nbonacci, int]{}
	nbonacciIter.Init(&Nbonacci{[]int{0, 0, 1}},
		func(f *Nbonacci, cont func() int) int {
			return nbonacciIter.Suspend(f.n[0], func() int {
				sum := 0
				for i, x := range f.n {
					sum += x
					if i > 0 {
						f.n[i-1] = x
					}
				}
				f.n[len(f.n)-1] = sum
				return nbonacciIter.Run(f, cont)
			})
		})

	for i := 0; i < 20; i++ {
		fmt.Println(nbonacciIter.Next())
	}

}
