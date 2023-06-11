package main

import (
	"encoding/json"
	"fmt"
	"log"

	// "math/random"
	"net/http"
	// "strconv"
	"github.com/gorilla/mux"
)

// struct je kao objekat u JS-u (key + value parovi)
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// * zvezdica je pointer na Director struct

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// w je response
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// isto kao foreach u JS-u
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

// Get single movie
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Kada ne koristimo index, moramo staviti _ (blank identifier)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// Svaki go file/program ima func main
func main() {
	// Declares and defines variable at the same time : = (zajedno)

	r := mux.NewRouter()

	// & referenca adrese objekta Directora
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "Luka", Lastname: "Golubovic"}})
	movies = append(movies, Movie{ID: "2", Isbn: "482726", Title: "Movie Two", Director: &Director{Firstname: "Katarina", Lastname: "Golubovic"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
