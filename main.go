package main

import (
	"fmt"
	"practice/arrayfill"
	"practice/throwdice"
)

func main() {
	fmt.Println("task 1 :")
	arrayfill.PrintArray(arrayfill.FillDiagonals(arrayfill.CreateArray(5, "0"), "1"))

	fmt.Println("task 2 :")
	numberShots := 1000
	m := make(map[int]int)

	for i := 0; i < numberShots; i++ {
		m[throwdice.RandInt()+throwdice.RandInt()] += 1
	}
	for i, v := range m {
		fmt.Printf("%d - %2.1f %%\n", i, float64(v)/float64(numberShots)*100)
	}
}
