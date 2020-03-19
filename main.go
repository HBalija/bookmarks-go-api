package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"bookmarks/models"

	"github.com/gorilla/mux"
	"github.com/lib/pq" // postgres driver
	"github.com/subosito/gotenv"
)

var db *sql.DB // db --> a pointer to global sql.DB type

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
	var b models.Bookmark
	var bs []models.Bookmark // empty slice

	rows, err := db.Query("SELECT * FROM bookmarks")
	logFatal(err)

	defer rows.Close() // close connection after function is executed

	log.Println(rows)

	for rows.Next() { // returns boolean
		err := rows.Scan(&b.ID, &b.Title, &b.URL)
		logFatal(err)

		bs = append(bs, b)
	}

	json.NewEncoder(w).Encode(bs)
}

func addBookmark(w http.ResponseWriter, r *http.Request) {
	var b models.Bookmark

	// decode request body and map value to "b" variable address
	json.NewDecoder(r.Body).Decode(&b)

	err := db.QueryRow(
		"INSERT INTO bookmarks (title, url) VALUES($1, $2) RETURNING id",
		b.Title, b.URL).Scan(&b.ID) // populate "b" objects id with created id

	if err != nil {
		json.NewEncoder(w).Encode(http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode("Created: " + strconv.Itoa(b.ID))
	}
}

func getBookmark(w http.ResponseWriter, r *http.Request) {
	var b models.Bookmark
	params := mux.Vars(r)

	row := db.QueryRow("select * from bookmarks where id=$1", params["id"])
	err := row.Scan(&b.ID, &b.Title, &b.URL)

	if err != nil {
		json.NewEncoder(w).Encode(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(b)
	}
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
