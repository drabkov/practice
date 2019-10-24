package main

import (
	"fmt"
	"practice/arrayfill"
	"practice/throwdice"
)

func main() {
	fmt.Println("task 1 :")
	arrayfill.PrintArray(arrayfill.FillArray(5))
	fmt.Println("task 2 :")
	throwdice.ThrowDice(1000)
}
