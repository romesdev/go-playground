package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkNodes(t, ch)
	close(ch)
}

func walkNodes(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	ch <- t.Value

	if t.Right != nil {
		walkNodes(t.Right, ch)
	}

	if t.Left != nil {
		walkNodes(t.Left, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	chOne := make(chan int)
	chTwo := make(chan int)

	go Walk(t1, chOne)
	go Walk(t2, chTwo)

	for v := range chOne {
		if v != <-chTwo {
			return false
		}
	}

	/*

		valuesOne := make([]int, 10)
		valuesTwo := make([]int, 10)

		i := 0
		j := 0

		for v := range chOne {
			valuesOne[i] = v
			i++
		}

		for v := range chTwo {
			valuesTwo[j] = v
			j++
		}

		fmt.Println(valuesOne, valuesTwo)

		for i, _ := range valuesOne {
			if valuesOne[i] != valuesTwo[i] {
				return false
			}
		}
	*/

	return true
}

func main() {
	k := 10
	t := tree.New(k)

	/*
		ch := make(chan int, k)

		go Walk(tree.New(k), ch)

		for v := range ch {
			fmt.Println(v)
		}
	*/

	// same
	fmt.Println(Same(t, t))
	// different
	fmt.Println(Same(tree.New(k), tree.New(k)))

}
