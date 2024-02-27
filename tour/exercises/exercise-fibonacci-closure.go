package tour

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		c := a
		a = b
		b = b + c
		return c
	}
}

func main() {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
