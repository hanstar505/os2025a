package main

import (
	"fmt"
	"greeting"
	"log"

	"github.com/os2025a.git/keyboard"
	//"week10/pkg/keyboard"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	fmt.Print("실수 입력 : ")
	n, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.1f점은 합격!\n", n)
}
