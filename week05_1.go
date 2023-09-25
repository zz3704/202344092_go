package main

import (
	"time"
	"fmt"
)

func main()  {
	brokenWords :="cs r?ck~"
	replaceWords :=string.NewReplacer("?","o")
	fixeWords := replaceWords.Replace(brokenWords)
	fmt.Println(fixedWords)
}