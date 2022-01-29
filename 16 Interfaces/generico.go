package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	fmt.Println(1, 2, "string", false, true, float64(123456))

	mapa := map[interface{}]interface{}{
		1:            "string",
		float32(100): true,
		"String":     "string",
		true:         float64(12),
	}

	fmt.Println(mapa)
}
