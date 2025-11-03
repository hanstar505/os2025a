package main

import (
	"fmt"
	"greeting"
	"log"
	"week10/pkg/keyboard"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	fmt.Print("점수 입력 : ")
	score, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	if score >= 80 {
		fmt.Printf("%.1f점은 합격!", score)
	}
}
