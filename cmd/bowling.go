package main

import (
	"fmt"

	b "bowling_game"
)

func main() {
	g := b.NewBowlingGame()
	for i := 0; i < 13; i++ {
		err := g.AcceptRoll(10)
		if err != nil {
			fmt.Printf("acceptRoll failed with error: %s", err.Error())
			fmt.Println("")
			continue
		}
		fmt.Printf("Remaining rolls in current frame: %d", g.GetRemainingRollsForCurrentFrame())
		fmt.Println("")
		fmt.Printf("Current score: %d", g.GetScore())
		fmt.Println("")
		fmt.Printf("Game finished: %t", g.Finished())
		fmt.Println("")
	}
}
