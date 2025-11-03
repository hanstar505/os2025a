package main

import "fmt"

func main() {
	arrauBool := [3]bool{true, false, true} //배열 리터럴
	var arrayInt [3]int
	fmt.Println(arrauBool[1]) //zero value
	arrayInt[1] = 2           // arrayInt[1] + 1
	fmt.Println(arrayInt[1])  //zero value + 2
}
