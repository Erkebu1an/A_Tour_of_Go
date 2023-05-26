package main

import (
	"fmt"
)

func main() {
	n := 0
	PointOne(&n)
	fmt.Println(n)
}
func PointOne(n *int) {
	*n = 1
}
func UltimatePointOne(n ***int) {
	***n = 0
}
