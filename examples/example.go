package main

import (
	"fmt"
	"github.com/jasocox/figo"
)

const SIZE = 10

func main() {
	q := figo.New()

	fmt.Printf("Is empty: %t\n", q.IsEmpty())

	for i := 0; i < SIZE; i++ {
		q.Push(i)
	}

	fmt.Println("Iterating")
	for elem := q.Front(); elem != nil; elem = q.Next(elem) {
		fmt.Printf("%d ", elem.Value)
	}
	fmt.Println()

	fmt.Printf("Is empty: %t\n", q.IsEmpty())

	fmt.Println("Popping")
	for val := q.Pop(); val != nil; val = q.Pop() {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	fmt.Printf("Is empty: %t\n", q.IsEmpty())
}
