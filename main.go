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

	fmt.Println("Pemain:", players)
	fmt.Println("Dadu awal:", dice)
}
