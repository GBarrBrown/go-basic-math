package main

import "fmt"

func sum(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(sum(12, 3))
	fmt.Println(sum(8, -3))
	fmt.Println(sum(-8, 6))
}