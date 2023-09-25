package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

func main() {
	fmt.Print("정수 입력 : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n') //option2
	if err !=nil{  //conditionals
		log.Fatal(err)
	}
	log.Fatal(err)
	fmt.Println(inputScore)
}
