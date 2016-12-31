/*
 * Example of how to insert a value in to a slice.!
 */

package main

import (
	"fmt"
)

// Insert value V in position I of slice S.
func insert(s []int, i int, v int) {
	copy(s[i+1:], s[i:])
	s[i] = v
}

func main() {
	s := make([]int, 5, 5)
	s[0] = 0
	s[1] = 1
	s[2] = 2

	insert(s, 1, 4)
	fmt.Printf("%v\n", s)
}
