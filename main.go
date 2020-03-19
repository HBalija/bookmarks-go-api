package main

import (
	"github.com/gorilla/mux" // postgres driver
	"github.com/subosito/gotenv"

	"log"
	"net/http"

	"bookmarks/drivers"
	"bookmarks/views"
)

func init() {
	gotenv.Load() // load env variables from ".env" file
}

func main() {

	// conect to db instance
	db := drivers.ConnectDB()

	// initialize router
	r := mux.NewRouter()

	// routes
	r.HandleFunc("/bookmarks/", views.GetBookmarks(db)).Methods("GET")
	r.HandleFunc("/bookmarks/", views.AddBookmark(db)).Methods("POST")
	r.HandleFunc("/bookmarks/", views.UpdateBookmark(db)).Methods("PUT")
	r.HandleFunc("/bookmarks/{id}/", views.GetBookmark(db)).Methods("GET")
	r.HandleFunc("/bookmarks/{id}/", views.RemoveBookmark(db)).Methods("DELETE")

	// run server on port 8000 and log error if any
	log.Fatal(http.ListenAndServe(":8000", r))
}
