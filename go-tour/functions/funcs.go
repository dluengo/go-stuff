package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first_call, second_call := true, true
	curr, last := 1, 0
	return func() int {
		if first_call {
			first_call = false
			return 0
		} else if second_call {
			second_call = false
			return 1
		} else {
			aux := curr
			curr += last
			last = aux
			return curr
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

