package dicegame

import (
	"testing"
)

type GameParams struct {
	N int // Jumlah pemain
	M int // Jumlah dadu
}

func TestPlayGame(t *testing.T) {
	tests := []struct {
		name   string
		params GameParams
		want   struct {
			winner int
			score  int
		}
	}{
		{
			name: "2 pemain & 4 dadu",
			params: GameParams{
				N: 2,
				M: 4,
			},
			want: struct {
				winner int
				score  int
			}{
				winner: 1,
				score:  0,
			},
		},
		{
			name: "2 pemain & 2 dadu",
			params: GameParams{
				N: 2,
				M: 2,
			},
			want: struct {
				winner int
				score  int
			}{
				winner: 1,
				score:  0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			winner, score := PlayGame(tt.params.N, tt.params.M)
			if winner != tt.want.winner || score != tt.want.score {
				t.Errorf("PlayGame() = (%v, %v), want (%v, %v)", winner, score, tt.want.winner, tt.want.score)
			}
		})
	}
}
