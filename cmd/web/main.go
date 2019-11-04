package main

import (
    "crypto/tls"
	"log"
	"net/http"
	"html/template"
	"flag"
	"os"
	"time"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ECAllen/lets-go/pkg/models/sqlite"
	"github.com/golangcollege/sessions"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	session *sessions.Session
	memories *sqlite.MemoryModel
	templateCache map[string]*template.Template
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
	secret := flag.String("secret", "TODO replace w 32 byte random str" , "Secret key")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	database, err := openDB("./memories.db")
	if err != nil {
		errorLog.Fatal(err)
	}

	defer database.Close()

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	templateCache, err := newTemplateCache("./ui/html")
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		session: session,
		memories: &sqlite.MemoryModel{DB: database},
		templateCache: templateCache,
	}

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}
	

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
		TLSConfig: tlsConfig,
		IdleTimeout: time.Minute,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServeTLS("./tls/cert.pem","./tls/key.pem")
	errorLog.Fatal(err)
}
