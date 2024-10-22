package internal

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)


func FetchFIDEProfile(fideID string) (*FIDEProfile, error) {
    url := fmt.Sprintf("https://ratings.fide.com/profile/%s", fideID)

    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch FIDE profile: %w", err)
    }
    defer resp.Body.Close()

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to parse FIDE profile HTML: %w", err)
    }

    name := doc.Find(".profile-top-title").First().Text()
    ratingStr := strings.TrimSpace(doc.Find(".profile-top-rating-data_gray").Contents().Last().Text())
    var rating int
    fmt.Sscanf(ratingStr, "%d", &rating)

    return &FIDEProfile{
        Name: name,
        Rating: rating,
        }, nil
}
