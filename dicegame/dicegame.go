package dicegame

import (
	"fmt"
	"math/rand"
)

func PlayGame(n, m int) (int, int) {

	players := make([]int, n)
	for i := range players {
		players[i] = m
	}

	round := 1
	for {
		fmt.Printf("===\nGiliran %d lempar dadu:\n", round)

		for i, player := range players {
			if player == 0 {
				continue
			}

			daduDilempar := make([]int, player)
			for j := 0; j < player; j++ {
				daduDilempar[j] = rand.Intn(6) + 1
			}

			fmt.Printf("Pemain #%d (%d): %v\n", i+1, player, daduDilempar)

			// Evaluasi hasil lemparan dadu
			for j := 0; j < len(daduDilempar); {
				if daduDilempar[j] == 1 {
					if j < len(daduDilempar)-1 {
						players[i+1]++
					} else {
						players[0]++
					}
					copy(daduDilempar[j:], daduDilempar[j+1:])
					daduDilempar = daduDilempar[:len(daduDilempar)-1]
				} else {
					j++
				}
			}

			players[i] = len(daduDilempar)

			fmt.Printf("Setelah evaluasi:\n")
			fmt.Printf("Pemain #%d (%d): %v\n", i+1, len(daduDilempar), daduDilempar)
		}

		// Cek apakah hanya ada satu pemain yang memiliki dadu
		sisaPlayers := 0
		winnerIndex := -1
		for i, player := range players {
			if player > 0 {
				sisaPlayers++
				winnerIndex = i
			}
		}

		if sisaPlayers == 1 {
			fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu.\n", winnerIndex+1)
			return winnerIndex + 1, players[winnerIndex]
		}

		round++
	}
}
