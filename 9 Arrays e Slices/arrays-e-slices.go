package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]string{"Position 1", "Position 2", "Position 3", "Position 4", "Position 5"}

	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{6, 7, 8, 9, 10}
	slice = append(slice, 11)
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice2 := array2[1:3]
	fmt.Println(slice2)
	array2[1] = "Position diff"
	fmt.Println(slice2)

	// Arrays Internos

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	slice4 := make([]float32, 5)
	slice4 = append(slice4, 10)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // length
	fmt.Println(cap(slice4)) // capacidade

}
