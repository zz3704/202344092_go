package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("정수 입력 : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	if inputScore >= 90 {
		grade := "A grade"
	} else {
		grade := "under grade"
	}
	fmt.Println(inputScore)
}
