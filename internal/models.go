package internal

// Main
type Player struct {
    ID      int
    Name   string
    FIDE   int
    Rating int
    Chance float64
    Score int
    Standing int
}

type Game struct {
    ID int
    PlayerOneID int
    PlayerTwoID int
    Winer int // -1, not played, 1 Player1, 2 Player2, 3 Draw
}

type GridData struct {
    Players []Player
    Games []Game
    NumPrizes int
}

// Fetch
type FIDEProfile struct {
    Name  string
    Rating int
}
