package main

import (
	"fmt"
	"reflect"
)

func main() {
	// casting
	var length float64 = 1.2
	var width float64 = 2
	fmt.Println("면적은", length*width)
	fmt.Println("길이 > 너비?", length > float64(width))
	fmt.Println("형변환", reflect.TypeOf(float64(length)))
	fmt.Println("원본", reflect.TypeOf(length))
}
