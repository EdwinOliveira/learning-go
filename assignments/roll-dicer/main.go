package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for range 50 {
		diceNumberOne := rand.Intn(6)
		diceNumberTwo := rand.Intn(6)

		switch(diceNumberOne + diceNumberTwo) {
			case 7: {
				fmt.Println("NATURAL")
				break;
			}
			case 11: {
				fmt.Println("NATURAL")
				break;
			}
			case 2: {
				fmt.Println("SNAKE-EYES-CRAPS")
				break;
			}
			case 3: {
				fmt.Println("LOSS-CRAPS")
				break;
			}
			case 12: {
				fmt.Println("LOSS-CRAPS")
				break;
			}
			default: {
				fmt.Println("NEUTRAL")
			}
		}
	}

}