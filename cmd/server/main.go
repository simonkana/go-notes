package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/simonkana/go-notes/internal/db"
	"github.com/simonkana/go-notes/internal/notes"

	_ "modernc.org/sqlite"
)

func main() {
	conn, err := sql.Open("sqlite", "file:notes.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(conn)
	mux := http.NewServeMux()

	mux.Handle("/", notes.NewHandler(queries))

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
