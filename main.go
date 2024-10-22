package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)


type Player struct {
    Name   string
    FIDE   *int
    Chance float64
}

type GridData struct {
    Players []Player
}

func serveFrontend(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func generateGrid(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Unable to parse form", http.StatusBadRequest)
        return
    }

    numPlayers, errPlayers := strconv.Atoi(r.FormValue("numplayers"))
    if errPlayers != nil {
        log.Print("Unable to parse players input")
    }
    // numPrizes, errPrizes := strconv.Atoi(r.FormValue("numprizes"))
    // if errPlayers != nil || errPrizes != nil {
    //     log.Print("Unable to parse players or prizes input")
    // }




    players := make([]Player, numPlayers)
    for i := 0; i < numPlayers; i++ {
        players[i] = Player{
            Name: "Player " + strconv.Itoa(i + 1),
            Chance: 100 / float64(numPlayers),
        }
    }
    gridData := GridData{
        Players: players,
    }

    tmpl := template.Must(template.ParseFiles("templates/grid.html"))
    err = tmpl.Execute(w, gridData)
    if err != nil {
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
        return
    }



    // w.Header().Set("Content-Type", "text/html; charset=UTF-8")

    // fmt.Fprintf(w, "<form hx-post='/update' hx-target='#grid' hx-swap='innerHTML'>")
    // fmt.Fprintf(w, "<table>")
    // fmt.Fprintf(w, "<tr><th>Players</th><th>FIDE number</th><th>Chances</th>")

    // for i := 1; i < numPlayers; i++ {
    //     fmt.Fprintf(w, "<th>Game %d</th>", i)
    // }
    // fmt.Fprintf(w, "</tr>")

    // for i := 1; i <= numPlayers; i++ {
    //     fmt.Fprint(w, "<tr")
    //     if i <= numPrizes {
    //         fmt.Fprint(w, " class='prize'")
    //     }
    //     fmt.Fprintf(w, `
    //         ><td id="playername_%v">Player %v</td>
    //         <td><input type="number" id="payerid_%v"></td>
    //         <td><span id="chances_%v">%.2f<span>%%</td>
    //     `, i, i, i, i, 100 / float64(numPlayers) * float64(numPrizes))
    //     for j := 1; j < numPlayers; j++ {
    //         fmt.Fprintf(w, `
    //             <td id="player_%v_game_%v" class='score'>
    //                 <a href="#">✅</a>
    //                 <a href="#">½</a>
    //                 <a href="#">❌</a>
    //             </td>
    //         `, i, j)
    //     }
    //     fmt.Fprintf(w, "</tr>")
    // }

    // fmt.Fprintf(w, "</table>")
    // fmt.Fprintf(w, "<button type='submit'>Update grid</button>")
    // fmt.Fprintf(w, "</form>")
}

func main() {
	http.HandleFunc("/", serveFrontend)
	http.HandleFunc("/grid", generateGrid)

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
