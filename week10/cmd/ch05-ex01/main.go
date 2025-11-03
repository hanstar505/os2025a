package main

import "fmt"

func main() {
	var arrauBool [3]bool
	var arrayInt [3]int
	fmt.Println(arrauBool[1]) //zero value
	arrayInt[1]++             // arrayInt[1] + 1
	arrayInt[1]++             // arrayInt[1] + 1
	fmt.Println(arrayInt[1])  //zero value + 2
}
