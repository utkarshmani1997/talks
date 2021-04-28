package main

import "fmt"

func Sum(x, y int) int {
	return x + y
}

func main() {
	m, n := 3, 4
	fmt.Println(Sum(m, n))
}
