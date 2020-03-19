package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"bookmarks/models"

	"github.com/gorilla/mux"
	"github.com/lib/pq" // postgres driver
	"github.com/subosito/gotenv"
)

var bs []models.Bookmark // empty slice
var db *sql.DB           // db --> a pointer to global sql.DB type

func init() {
	gotenv.Load() // load env variables from ".env" file
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	pgURL, err := pq.ParseURL(os.Getenv("DATABASE_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgURL)
	logFatal(err)

	err = db.Ping() // check connection to db and log error
	logFatal(err)

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
	var bookmark models.Bookmark
	rows, err := db.Query("select * from bookmarks")
	logFatal(err)

	defer rows.Close() // close connection after function is executed

	for rows.Next() { // returns boolean
		err := rows.Scan(&bookmark.ID, &bookmark.Title, &bookmark.URL)
		logFatal(err)

		bs = append(bs, bookmark)
	}

	json.NewEncoder(w).Encode(bs)
}

func addBookmark(w http.ResponseWriter, r *http.Request) {
	// var b models.Bookmark
	// // decode pointer response body into "b" memory address
	// json.NewDecoder(r.Body).Decode(&b)
	// bs = append(bs, b)

	// // return status 201
	// json.NewEncoder(w).Encode(http.StatusCreated)
}

func getBookmark(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// id, _ := strconv.Atoi(params["id"]) // Ascii to int

	// for _, v := range bs {
	// 	if v.ID == id {
	// 		// return response
	// 		json.NewEncoder(w).Encode(v)
	// 		break
	// 	}
	// }
}

func updateBookmark(w http.ResponseWriter, r *http.Request) {
	// var b models.Bookmark
	// json.NewDecoder(r.Body).Decode(&b)

	// for i, v := range bs {
	// 	if v.ID == b.ID {
	// 		bs[i] = b
	// 		// return status 200
	// 		json.NewEncoder(w).Encode(http.StatusOK)
	// 		break
	// 	}
	// }
}

func removeBookmark(w http.ResponseWriter, r *http.Request) {
	// 	params := mux.Vars(r)
	// 	id, _ := strconv.Atoi(params["id"])

	// 	for i, v := range bs {
	// 		if v.ID == id {
	// 			bs = append(bs[:i], bs[(i+1):]...) // slice till "i" and add values from "i" to end
	// 			break
	// 		}
	// 	}
	// 	// return status 204
	// 	json.NewEncoder(w).Encode(http.StatusNoContent)
}
