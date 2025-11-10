package main

import "fmt"

func main() {
	subjects := []string{"Go", "Javascript", "Python", "linux"}
	subjectsSlice := subjects[1:3] //slicing
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}
