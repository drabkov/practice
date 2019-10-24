package arrayfill

import "fmt"

func CreateArray(size int, symbol string) [][]string {
	a := make([][]string, size)
	for i := 0; i < size; i++ {
		a[i] = make([]string, size)
		for j := 0; j < size; j++ {
			a[i][j] = symbol
		}
	}
	return a
}

func FillDiagonals(a [][]string, symbol string) [][]string {
	size := len(a)
	for i := 0; i < size; i++ {
		a[i][i], a[i][size-1-i] = symbol, symbol
	}
	return a
}

func PrintArray(a [][]string) {
	for _, v := range a {
		fmt.Println(v)
	}
}
