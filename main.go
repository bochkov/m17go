package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	srv := &http.Server{Addr: ":3000", Handler: mux}

	notifyCtx, nStop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer nStop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen and serve: %v", err)
		}
	}()
	log.Printf("Listening on %v:%v", host, port)

	<-notifyCtx.Done()
	log.Printf("shutting down server gracefully")

	// close HTTP connections
	stopCtx, sStop := context.WithTimeout(context.Background(), 5*time.Second)
	defer sStop()
	if err := srv.Shutdown(stopCtx); err != nil {
		log.Fatalf("shutdown: %v", err)
	}
	// close DB connections
	if err := db.Close(); err != nil {
		log.Fatalf("shutdown db conn: %v", err)
	}

	log.Printf("server shutdown properly")
}
