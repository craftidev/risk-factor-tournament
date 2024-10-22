package internal

// Main
type Player struct {
    Name   string
    FIDE   int
    Rating int
    Chance float64
    Standing int
    Games []Game
}

type Game struct {
    Number int
    Opponent Player
    PlayerResult *float64
}

type GridData struct {
    Players []Player
    NumPrizes int
}

// Fetch
type FIDEProfile struct {
    Name  string
    Rating int
}
