package drivers

import (
	"github.com/lib/pq"

	"database/sql"
	"log"
	"os"
)

// LogFatal error handler
func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ConnectDB connect to database and return instance
func ConnectDB() *sql.DB {
	pgURL, _ := pq.ParseURL(os.Getenv("DATABASE_URL"))
	db, _ := sql.Open("postgres", pgURL)

	err := db.Ping() // check connection to db and log error
	LogFatal(err)

	return db
}
