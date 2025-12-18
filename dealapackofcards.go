package piscine

import (
	"fmt"
)

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		fmt.Printf("Player %d: ", i+1)

		// Each player gets 3 cards
		start := i * 3
		end := start + 3

		for j := start; j < end; j++ {
			if j == end-1 {
				fmt.Printf("%d", deck[j])
			} else {
				fmt.Printf("%d, ", deck[j])
			}
		}
		fmt.Printf("\n")
	}
}
