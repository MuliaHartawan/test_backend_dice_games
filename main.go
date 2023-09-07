package main

import (
	"fmt"
	"main/dicegame"
)

func main() {

	var N, M int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&N)
	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scan(&M)

	winner, score := dicegame.PlayGame(N, M)

	fmt.Printf("Game dimenangkan oleh pemain #%d dengan skor %d.\n", winner, score)
}
