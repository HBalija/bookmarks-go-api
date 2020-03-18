package main

import (
	"bookmarks/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var bs []models.Bookmark

func main() {

	// add static data
	bs = append(bs, models.Bookmark{ID: 1, Title: "Golang", URL: "https://golang.org/"},
		models.Bookmark{ID: 2, Title: "Python", URL: "https://www.python.org/"},
		models.Bookmark{ID: 3, Title: "DRF", URL: "https://www.django-rest-framework.org/"},
		models.Bookmark{ID: 4, Title: "Angular", URL: "https://angular.io/"},
		models.Bookmark{ID: 5, Title: "React", URL: "https://reactjs.org/"},
	)

	// initialize router
	r := mux.NewRouter()

	// routes with handlers
	r.HandleFunc("/bookmarks/", getBookmarks).Methods("GET")
	r.HandleFunc("/bookmarks/", addBookmark).Methods("POST")
	r.HandleFunc("/bookmarks/", updateBookmark).Methods("PUT")
	r.HandleFunc("/bookmarks/{id}/", getBookmark).Methods("GET")
	r.HandleFunc("/bookmarks/{id}/", removeBookmark).Methods("DELETE")

	// run server on port 8000 and log error if any
	log.Fatal(http.ListenAndServe(":8000", r))
}

// HANDLERS

func getBookmarks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(bs)
}

func addBookmark(w http.ResponseWriter, r *http.Request) {
	var b models.Bookmark
	// decode pointer response body into "b" memory address
	json.NewDecoder(r.Body).Decode(&b)
	bs = append(bs, b)

	// return status 201
	json.NewEncoder(w).Encode(http.StatusCreated)
}

func getBookmark(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"]) // Ascii to int

	for _, v := range bs {
		if v.ID == id {
			// return response
			json.NewEncoder(w).Encode(v)
			break
		}
	}
}

func updateBookmark(w http.ResponseWriter, r *http.Request) {
	var b models.Bookmark
	json.NewDecoder(r.Body).Decode(&b)

	for i, v := range bs {
		if v.ID == b.ID {
			bs[i] = b
			// return status 200
			json.NewEncoder(w).Encode(http.StatusOK)
			break
		}
	}
}

func removeBookmark(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, v := range bs {
		if v.ID == id {
			bs = append(bs[:i], bs[(i+1):]...) // slice till "i" and add values from "i" to end
			break
		}
	}
	// return status 204
	json.NewEncoder(w).Encode(http.StatusNoContent)
}
