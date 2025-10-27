package main

import (
	"fmt"
	"log"
)

func GetFloat() (float64, error) {

}

func main() {
	fmt.Print("점수 입력 :")
	score, err := GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	status := ""
	if score >= 60 {
		status = "합격"
	} else {
		status = "불합격"
	}
	fmt.Printf("%.2f점은 %s\n", score, status)
}
