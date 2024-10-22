package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

    "github.com/craftidev/riskfactortournament/internal"
)



func serveFrontend(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func generateGrid(
    w http.ResponseWriter,
    players []internal.Player,
    numPrizes int,
) {
    numPlayers := len(players)
    for pIx := range players {
        games := make([]internal.Game, 0, numPlayers - 1)
        gameNumber := 1
        for gIx := 0; gIx < numPlayers; gIx++ {
            if pIx == gIx {
                continue
            }
            games = append(games,
                internal.Game{
                    Number: gameNumber,
                    Opponent: players[gIx],
                    PlayerResult: nil,
                },
            )
            gameNumber++
        }
        players[pIx].Games = games
    }

    gridData := internal.GridData{
        Players: players,
        NumPrizes: numPrizes,
    }

    tmpl := template.Must(template.ParseFiles("templates/grid.html"))
    err := tmpl.Execute(w, gridData)
    if err != nil {
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
        return
    }
}

func initGrid(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Unable to parse form", http.StatusBadRequest)
        return
    }

    numPlayers, errPlayers := strconv.Atoi(r.FormValue("numplayers"))
    numPrizes, errPrizes := strconv.Atoi(r.FormValue("numprizes"))
    if errPlayers != nil || errPrizes != nil {
        log.Print("Unable to parse players or prizes input")
    }

    players := make([]internal.Player, numPlayers)
    for i := 0; i < numPlayers; i++ {
        players[i] = internal.Player{
            Name: "Player " + strconv.Itoa(i + 1),
            Chance: 100 / float64(numPlayers),
            Standing: i,
        }
    }

    generateGrid(w, players, numPrizes)
}

func updateGrid(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err!= nil {
        http.Error(w, "Unable to parse form", http.StatusBadRequest)
        return
    }

    numPlayers := len(r.Form) -1
    players := make([]internal.Player, numPlayers)

    for i := 0; i < numPlayers; i++ {
        fideStr := r.FormValue(fmt.Sprintf("playerFIDE_%d", i))
        profile, err := internal.FetchFIDEProfile(fideStr)
        if err != nil {
            http.Error(w, "Failed to fetch FIDE profile", http.StatusInternalServerError)
            return
        }

        fide, err := strconv.Atoi(fideStr)
        if err != nil {
            http.Error(w, "Invalid FIDE input", http.StatusBadRequest)
        }
        players[i].Name = profile.Name
        players[i].FIDE = fide
        players[i].Rating = profile.Rating
    }

    numPrizesStr := r.FormValue("numPrizes")
    numPrizes, err := strconv.Atoi(numPrizesStr)
    if err != nil {
        http.Error(w, "Unable to read number of prizes at this stage", http.StatusBadRequest)
    }

    generateGrid(w, players, numPrizes)
}

func main() {
	http.HandleFunc("/", serveFrontend)
	http.HandleFunc("/grid", initGrid)
	http.HandleFunc("/update-grid", updateGrid)

    // Serve static files
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server started at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
