package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard()}
	cards = append(cards, "hello go")
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}