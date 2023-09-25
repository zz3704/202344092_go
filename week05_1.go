package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(100) + 1
	fmt.Println("guess game(1~100) : ")
	fmt.Println(answer)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i > 10; i++ {
		fmt.Println("you chance : ", 10-i)
		fmt.Print("guess number : ")
		inputNumber, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputNumberString = string.TrimString(inputNumberstring)
		inputNumber, err := strconv.Atoi(inputNumberString)
		if err != nil {
			log.Fatal(err)
		}
    if inputNumber = answer {
      fmt.Println("gread u got the number.") //answer is higher
      }
		else if inputNumber < answer {
			fmt.Println("ur guess number is lower than answer.") //answer is higher
		} else if inputNumber > answer {
			fmt.Println("ur guess number is higher than answer.") //answer is lower
		}
	}
}
