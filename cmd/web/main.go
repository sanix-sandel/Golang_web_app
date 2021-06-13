package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"sanix.net/snippetbox/pkg/models/postgres"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *postgres.SnippetModel
}

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "sanix"
	password = "19972017"
	dbname   = "snippets"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	//Command line flag for the MySQL DSN string
	//dsn := flag.String("dsn", "root:19972017Russi@tcp(127.0.0.1:3306)/snippets?parseTime=true", "MySQL database")
	var dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//flag.Parse()

	//Create a logger for writing information messages. This three parameters:
	//the destination to write the logs to (os.Stdout), a st prefix for message(INFO followed by a tab)
	//and flags to indicate what additional information to include (local data and time). Note that are
	//joined using the bitwise OR operator |.
	//
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &postgres.SnippetModel{DB: db},
	}
	//Initialize a new http.Server strcut
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

	//logger for error message
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
