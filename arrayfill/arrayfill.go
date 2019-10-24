package arrayfill

import "fmt"

func FillArray(size int) [][]int {

	a := make([][]int, size)

	for i := 0; i < size; i++ {
		a[i] = make([]int, size)
		a[i][i] = 1
		a[i][size-1-i] = 1
	}

	return a

}

func PrintArray(a [][]int) {
	for _, v := range a {
		for _, f := range v {
			/*if f == 1 {
				fmt.Print("*")
			} else {*/
			fmt.Print(f)
			//}
			fmt.Print(" ")
		}
		fmt.Println("")
	}

}
