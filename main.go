package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var N, M int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&N)
	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scan(&M)

	players := make([]int, N)
	dice := make([]int, M)

	for i := range dice {
		dice[i] = rand.Intn(6) + 1
	}

	round := 1
	for {
		fmt.Printf("==================\nGiliran %d lempar dadu:\n", round)

		for i := range players {
			if players[i] == 0 {
				continue
			}

			// Lepaskan dadu dari pemain
			playerDice := make([]int, players[i])
			copy(playerDice, dice)

			// Evaluasi hasil lemparan dadu
			for j := 0; j < len(playerDice); {
				if playerDice[j] == 6 {
					players[i]++
					playerDice = append(playerDice[:j], playerDice[j+1:]...)
				} else if playerDice[j] == 1 {
					if i == len(players)-1 {
						players[0]++
					} else {
						players[i+1]++
					}
					playerDice = append(playerDice[:j], playerDice[j+1:]...)
				} else {
					j++
				}
			}

			// Perbarui jumlah dadu pemain
			players[i] = len(playerDice)

			fmt.Printf("Setelah evaluasi:\n")
			fmt.Printf("Pemain #%d (%d): %v\n", i+1, players[i], playerDice)

		}

		// TODO: Cek apakah permainan berakhir
		// TODO: Jika hanya ada satu pemain dengan dadu, umumkan pemenang dan akhiri permainan

		round++
	}
}
