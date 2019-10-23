package main

import (
	"practice/arrayfill"
	"practice/throwdice"
)

func main(){
	arrayfill.PrintArray(arrayfill.FillArray(5))
	throwdice.ThrowDice(1000);
}