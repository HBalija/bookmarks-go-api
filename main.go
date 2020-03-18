package main

import (
	"log"
	"net/http"

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

func getBookmarks(http.ResponseWriter, *http.Request) {
	log.Println("Get all bookmarks")
}

func addBookmark(http.ResponseWriter, *http.Request) {
	log.Println("Add new bookmark")
}

func getBookmark(http.ResponseWriter, *http.Request) {
	log.Println("Get one bookmark")
}

func updateBookmark(http.ResponseWriter, *http.Request) {
	log.Println("Update a bookmark")
}

func deleteBookmark(http.ResponseWriter, *http.Request) {
	log.Println("Delete a bookmark")
}
