package main

import (
	"fmt"
)

func main() {

	var N, M int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&N)
	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scan(&M)

	players := make([]int, N)
	for i := range players {
		players[i] = M
	}
}
