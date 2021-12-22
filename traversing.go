package traversing

import (
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var stack []*tree.Tree
	current := t

	for current != nil || len(stack) > 0 {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ch <- current.Value
			current = current.Right
		}
	}

	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		v1, okV1 := <-c1
		v2, okV2 := <-c2
		if okV1 && okV2 {
			if v1 != v2 {
				return false
			}
		}
		if !okV1 && !okV2 {
			return true
		}
	}
}
