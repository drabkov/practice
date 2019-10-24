package arrayfill

import "fmt"

func CreateArray(size int) [][]string {
	a := make([][]string, size)
	for i := 0; i < size; i++ {
		a[i] = make([]string, size)
	}
	return a
}

func FillArray(a [][]string, symbol string) [][]string {
	size := len(a)
	for i := 0; i < size; i++ {
		a[i] = make([]string, size)
		a[i][i] = symbol
		a[i][size-1-i] = symbol
	}
	return a
}

func PrintArray(a [][]string) {
	for _, v := range a {
		for _, f := range v {
			if f == "" {
				fmt.Print("0")
			} else {
				fmt.Print(f)
			}
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}
