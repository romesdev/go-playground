package tour

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	list := &List[]{
		val: 42,
	}
	fmt.Println(list)
}
