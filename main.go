package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "selamat malam"

	letters := strings.Split(sentence, "")
	for _, letter := range letters {
		fmt.Println(letter)
	}
	space := make(map[string]int)

	for _, letter := range letters {
		if letter != " " {
			space[letter]++
		}
	}
	fmt.Println(space)
}
