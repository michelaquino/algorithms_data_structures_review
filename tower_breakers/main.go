package main

import (
	"fmt"
)

func main() {
	var n int32 = 10
	var m int32 = 796723

	fmt.Println(towerBreakers(n, m))
}

func towerBreakers(n int32, m int32) int32 {
	// if height == 1, the game is immediately over
	//     => p.1 has no moves; p.2 wins
	// if towers are divisible by 2,
	//     => 2 mimics p.1's moves; p.2 wins
	if n%2 == 0 || m == 1 {
		return 2
	}

	// otherwise, if towers is odd, p.1 can
	// take the first tower down to 1, effectively
	// making themselves p.2 and the tower numbers even,
	// which means they are now in the position to mimic
	// the original p.2's moves for the remaining towers
	// and win the game
	//    => the original p.1 wins
	return 1
}
