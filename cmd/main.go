package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bochkov/m17go/internal/albums"
	"github.com/bochkov/m17go/internal/gigs"
	"github.com/bochkov/m17go/internal/lib/db"
	"github.com/bochkov/m17go/internal/lib/router"
	"github.com/bochkov/m17go/internal/link"
	"github.com/bochkov/m17go/internal/members"
	"github.com/bochkov/m17go/internal/place"
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

	dbConn, err := db.NewDatabase(host, port, dbname, user, password)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	log.Printf("Connected to %v:%v\n", host, port)

	albums := albums.NewHandler(
		albums.NewService(
			albums.NewRepository(dbConn.GetDB()),
			link.NewRepository(dbConn.GetDB()),
		),
	)
	gigs := gigs.NewHandler(
		gigs.NewService(
			gigs.NewRepository(dbConn.GetDB()),
			place.NewRepository(dbConn.GetDB()),
		),
	)
	members := members.NewHandler(
		members.NewService(
			members.NewRepository(dbConn.GetDB()),
		),
	)

	engine := router.InitRouter(albums, gigs, members)
	srv := &http.Server{Addr: ":5000", Handler: engine}
	notifyCtx, nStop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer nStop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen and serve: %v\n", err)
		}
	}()
	log.Printf("Listening on :5000")
	<-notifyCtx.Done()
	log.Println("shutting down")

	// close HTTP connections
	stopCtx, sStop := context.WithTimeout(context.Background(), 5*time.Second)
	defer sStop()
	if err := srv.Shutdown(stopCtx); err != nil {
		log.Fatalf("shutdown: %v\n", err)
	}
	// close DB connections
	if err := dbConn.Close(); err != nil {
		log.Fatalf("shutdown db conn: %v\n", err)
	}

	log.Println("server is shutdown")
}
