package main

import "fmt"

func main() {
	subjects := [4]string{"Go", "Javascript", "Python", "linux"}
	subjectsSlice := subjects[:3] //slicing
	//subjects[0] = "Java"
	subjectsSlice[0] = "Java"
	subjectsSlice = append(subjectsSlice, "Go")
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}
