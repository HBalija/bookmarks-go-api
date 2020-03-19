package views

import (
	"github.com/gorilla/mux"

	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"bookmarks/drivers"
	"bookmarks/models"
)

// type BookmarkViews struct{}

// GetBookmarks handler for list bookmarks
func GetBookmarks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b models.Bookmark
		var bs []models.Bookmark // empty slice

		rows, err := db.Query("SELECT * FROM bookmarks")
		drivers.LogFatal(err)

		defer rows.Close() // close connection after function is executed

		for rows.Next() { // returns boolean
			err := rows.Scan(&b.ID, &b.Title, &b.URL)
			drivers.LogFatal(err)

			bs = append(bs, b)
		}

		json.NewEncoder(w).Encode(bs)
	}
}

// AddBookmark handler for creating bookmarks
func AddBookmark(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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
}

// GetBookmark handler for retrieving bookmarks
func GetBookmark(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}

// UpdateBookmark handler for updating bookmarks
func UpdateBookmark(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b models.Bookmark
		json.NewDecoder(r.Body).Decode(&b)

		_, err := db.Exec("UPDATE bookmarks SET title=$1, url=$2 WHERE id=$3",
			b.Title, b.URL, b.ID)

		if err != nil {
			json.NewEncoder(w).Encode(http.StatusBadRequest)
		} else {
			json.NewEncoder(w).Encode(b)
		}
	}
}

// RemoveBookmark handler for deleting bookmarks
func RemoveBookmark(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		_, err := db.Exec("DELETE FROM bookmarks WHERE id=$1", params["id"])
		drivers.LogFatal(err)

		json.NewEncoder(w).Encode(http.StatusNoContent)
	}
}
