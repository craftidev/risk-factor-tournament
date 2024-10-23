package internal

import (
	"fmt"
	"math"
	"math/rand"
)


func expectedScore(eloA, eloB int) float64 {
    adjustedDiff := .84 * float64(eloB - eloA)
    return 1 / (1 + math.Pow(10, adjustedDiff / 400))
}

func simulateGame(playerA, playerB Player) bool {
    expectedA := expectedScore(playerA.Rating, playerB.Rating)
    result := rand.Float64()

    if result < expectedA {
        return true
    } else {
        return false
    }
}

func simulateTournament(players []Player) {
    for pIx := 0; pIx < len(players); pIx++ {
        for gIx := 0; gIx < len(players); gIx++ {
            if pIx == gIx {
                continue
            }
            isResultWinForA := simulateGame(players[pIx], players[gIx])
            // TEST
            fmt.Printf(
                "Player %s vs Player %s -> %v\n",
                players[pIx].Name,
                players[gIx].Name,
                isResultWinForA,
            )
        }
    }
}
