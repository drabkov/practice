package main

import (
	"fmt"
	"practice2/list"
)

func main() {
	collection := list.Collection{}
	for i := 0; i < 10; i++ {
		collection.Add(i)
	}
	collection.Print()
	collection.Remove(4)
	collection.Print()
	fmt.Println(collection.Get(3))
	fmt.Println(collection.First())
	fmt.Println(collection.Last())
	fmt.Println(collection.Length())
}
