package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/disiqueira/gotree"
)

func graph(name string, decks, cards int) gotree.Tree {
	tree := gotree.New(name)
	for i := 0; i < decks; i++ {
		deckName := fmt.Sprintf("Deck %d", i+1)
		deck := tree.Add(deckName)
		for j := 0; j < cards; j++ {
			cardName := fmt.Sprintf("Card %d", j+1)
			deck.Add(cardName)
		}
	}
	return tree
}

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1:]
		if args[0] == "graph" {
			if len(args) == 4 {
				decks, err := strconv.Atoi(args[2]) //error check
				cards, err2 := strconv.Atoi(args[3])
				if err != nil {
					log.Fatal(err)
					fmt.Println("Error: please enter int for deckNumber")
				} else if err2 != nil {
					log.Fatal(err2)
					fmt.Println("Error: please enter int for cardNumber")
				} else {
					tree := graph(args[1], decks, cards)
					fmt.Println(tree.Print())
				}
			} else {
				fmt.Println("Error: need treeName, deckNumber, and cardNumber")
			}
		} else {
			fmt.Println("Error: invalid function argument. Use 'graph' to implement graph function")
		}
	} else {
		fmt.Println("Error: no arguments given")
	}
}
