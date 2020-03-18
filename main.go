package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// MODEL

type bookmark struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

var bs []bookmark

func main() {

	// add static data
	bs = append(bs, bookmark{ID: 1, Title: "Golang", URL: "https://golang.org/"},
		bookmark{ID: 2, Title: "Python", URL: "https://www.python.org/"},
		bookmark{ID: 3, Title: "DRF", URL: "https://www.django-rest-framework.org/"},
		bookmark{ID: 4, Title: "Angular", URL: "https://angular.io/"},
		bookmark{ID: 5, Title: "React", URL: "https://reactjs.org/"},
	)

	// initialize router
	r := mux.NewRouter()

	// routes with handlers
	r.HandleFunc("/bookmarks/", getBookmarks).Methods("GET")
	r.HandleFunc("/bookmarks/", addBookmark).Methods("POST")
	r.HandleFunc("/bookmarks/{id}/", getBookmark).Methods("GET")
	r.HandleFunc("/bookmarks/{id}/", updateBookmark).Methods("PUT")
	r.HandleFunc("/bookmarks/{id}/", deleteBookmark).Methods("DELETE")

	// run server on port 8000 and log error if any
	log.Fatal(http.ListenAndServe(":8000", r))
}

// HANDLERS

func getBookmarks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(bs)
}

func addBookmark(w http.ResponseWriter, r *http.Request) {
	var b bookmark
	// decode pointer response body into "b" memory address
	json.NewDecoder(r.Body).Decode(&b)

	bs = append(bs, b)

	// return response (201)
	json.NewEncoder(w).Encode(http.StatusCreated)
}

func getBookmark(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"]) // Ascii to int

	for _, v := range bs {
		if v.ID == id {
			// return response
			json.NewEncoder(w).Encode(v)
		}
	}
}

func updateBookmark(http.ResponseWriter, *http.Request) {
	log.Println("Update a bookmark")
}

func deleteBookmark(http.ResponseWriter, *http.Request) {
	log.Println("Delete a bookmark")
}
