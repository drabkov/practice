package main

import (
	"fmt"
	"practice/arrayfill"
	"practice/throwdice"
)

func main() {
	fmt.Println("task 1 :")
	arrayfill.PrintArray(arrayfill.FillArray(arrayfill.CreateArray(5), "1"))

	fmt.Println("task 2 :")
	numberShots := 1000
	m := throwdice.ThrowDice(numberShots)
	for i, v := range m {
		fmt.Printf("%d - %2.1f %%\n", i, float64(v)/float64(numberShots)*100)
	}
}
