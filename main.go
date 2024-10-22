package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)


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
    numPrizes, errPrizes := strconv.Atoi(r.FormValue("numprizes"))
    if errPlayers != nil || errPrizes != nil {
        log.Print("Unable to parse players or prizes input")
    }

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

    fmt.Fprintf(w, "<table>")
    fmt.Fprintf(w, "<tr><th>Players</th><th>FIDE number</th>")

    for i := 1; i < numPlayers; i++ {
        fmt.Fprintf(w, "<th>Game %d</th>", i)
    }
    fmt.Fprintf(w, "</tr>")

    for i := 0; i < numPlayers; i++ {
        fmt.Fprint(w, "<tr")
        if i < numPrizes {
            fmt.Fprint(w, " class='prize'")
        }
        fmt.Fprintf(w, `
            ><td><td id="playername_%v">Player %v</td>
            <td><input type="number"></td>
        `, i, i)
        for j := 0; j < numPlayers; j++ {
            fmt.Fprintf(w, "<td class='score'></td>")
        }
        fmt.Fprintf(w, "</tr>")
    }

    fmt.Fprintf(w, "</table>")
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
