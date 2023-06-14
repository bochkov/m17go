package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bochkov/m17go/internal/api"
	"github.com/bochkov/m17go/internal/db"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("M17 backend start == ", time.Now())

	host := os.Getenv("DB_HOST")
	log.Println("DB_HOST: ", host)
	port := os.Getenv("DB_PORT")
	log.Println("DB_PORT: ", port)
	dbname := os.Getenv("DB_NAME")
	log.Println("DB_NAME: ", dbname)
	user := os.Getenv("DB_USER")
	log.Println("DB_USER: ", user)
	password := os.Getenv("DB_PASSWORD")
	if host == "" || dbname == "" || user == "" {
		log.Fatalln("No env DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASSWORD. Exiting.")
	}

	db, err := db.NewDatabase(host, port, dbname, user, password).Connect()
	if err != nil {
		log.Fatalln(err)
	}

	mux := http.NewServeMux()
	api.ConfigureController(db, mux)
	log.Fatal(http.ListenAndServe(":5000", mux))
}
