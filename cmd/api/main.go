package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// underline syntax as we only use for init() side effects, not directly reference
	// as just importing them causes init() function to run
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

var counts int64

type Config struct {
	// DB *sql.DB
	// // models to type
	// Models data.Models
	Repo   data.Repository
	Client *http.Client
}

func main() {
	log.Println("Starting authentication service")

	// auth with routes
	// connect to DB
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect!")
	}

	// set up config
	app := Config{
		// DB:     conn,
		// Models: data.New(conn),
		Client: &http.Client{},
	}

	srv := &http.Server{
		// alternatively, Addr: ":8080"
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start server
	// add log.Fatal() optionally
	err := srv.ListenAndServe()
	if err != nil {
		// stops execution
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	// environment variable
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Database not ready ...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Waiting 2 seconds ...")
		time.Sleep(2 * time.Second)
		continue
	}
}

func (app *Config) setupRepo(conn *sql.DB) {
	db := data.NewPostgresRepository(conn)
	app.Repo = db
}
