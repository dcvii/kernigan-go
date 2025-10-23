package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int
	Actors []string
}

func main() {

	movies := []Movie{
		{Title: "Casablanca", Year: 1942, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "The Godfather", Year: 1972, Actors: []string{"Marlon Brando", "Al Pacino"}},
	}

	movies = append(movies, Movie{Title: "The Shawshank Redemption", Year: 1994, Actors: []string{"Tim Robbins", "Morgan Freeman"}})

	movies = append(movies, Movie{Title: "The Dark Knight", Year: 2008, Actors: []string{"Christian Bale", "Heath Ledger"}})

	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		log.Fatalf("Error marshaling JSON: %v", err)
		return
	}

	fmt.Println(string(data))

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		log.Fatalf("Error unmarshaling JSON: %v", err)

	}
	fmt.Println(titles)

	// fmt.Println()
	// for _, title := range titles {
	// 	fmt.Println(title.Title)
	// }

}
