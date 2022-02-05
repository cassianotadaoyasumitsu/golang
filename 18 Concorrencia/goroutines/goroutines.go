package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá Mundo.")
	escrever("Programando em GO")
}

func escrever(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
