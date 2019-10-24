package main

import (
	"fmt"
	"practice/arrayfill"
	"practice/throwdice"
	"sort"
)

func main() {
	fmt.Println("task 1 :")
	arrayfill.PrintArray(arrayfill.FillDiagonals(arrayfill.CreateArray(5, "0"), "1"))

	fmt.Println("task 2 :")
	numberShots := 1000
	m := make(map[int]int)

	for i := 0; i < numberShots; i++ {
		m[throwdice.RandInt()+throwdice.RandInt()]++
	}

	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d - %2.1f %%\n", k, float64(m[k])/float64(numberShots)*100)
	}
}
