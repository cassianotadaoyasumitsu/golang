package main

import "fmt"

var n int

func init() {
	fmt.Println("Func init executatda")
	n = 10
}

func main() {
	fmt.Println("Func main executada")
	fmt.Println(n)
}
