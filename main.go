package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "5894", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "5944", Title: "Movie Two", Director: &Director{FirstName: "Tom", LastName: "Chloe"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	PORT := ":8000"
	fmt.Printf("Server started on port %v", PORT)

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	err := http.ListenAndServe(PORT, r)
	if err != nil {
		log.Fatal(err)
	}
}
