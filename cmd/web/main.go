package main

import (
	"log"
	"net/http"
	"flag"
	"os"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ECAllen/lets-go/pkg/models/sqlite"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	memories *sqlite.MemoryModel
}

func openDB(dbFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	// TODO does this make sense with sqlite ?
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	database, err := openDB("./memories.db")
	if err != nil {
		errorLog.Fatal(err)
	}

	defer database.Close()

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		memories: &sqlite.MemoryModel{DB: database},
	}

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
