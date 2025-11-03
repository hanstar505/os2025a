package main

import (
	"fmt"
)

func main() {
	numbers := [3]int{-9, 11, 7}
	for i, number := range numbers {
		fmt.Println(i, number)
	}

	//fmt.Println(reflect.TypeOf(arrayBool))
	//fmt.Printf("%v\n", arrayBool)
	//fmt.Printf("%v\n", arrayInt)
}
