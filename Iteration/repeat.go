package main

import (
	"fmt"
	"log"
	"strings"
)

func Repeat(character string, number int) string {
	var repeated strings.Builder
	for i := 0; i < number; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func main() {
	fmt.Println("Enter a number for how many charcters you want to print: ")
	var number int
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Repeat("a", number))
}
