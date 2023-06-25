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

	database, err := db.NewDatabase(host, port, dbname, user, password).Connect()
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	mux := http.NewServeMux()
	api.ConfigureController(database, mux)
	srv := &http.Server{Addr: ":5000", Handler: loggingWare(mux)}

	notifyCtx, nStop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer nStop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen and serve: %v\n", err)
		}
	}()
	log.Printf("Listening on %v:%v\n", host, port)

	<-notifyCtx.Done()
	log.Println("shutting down server gracefully")

	// close HTTP connections
	stopCtx, sStop := context.WithTimeout(context.Background(), 5*time.Second)
	defer sStop()
	if err := srv.Shutdown(stopCtx); err != nil {
		log.Fatalf("shutdown: %v\n", err)
	}
	// close DB connections
	if err := database.Close(); err != nil {
		log.Fatalf("shutdown db conn: %v\n", err)
	}

	log.Println("server shutdown properly")
}

func loggingWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		o := &responseWrapper{ResponseWriter: w}
		next.ServeHTTP(o, r)
		log.Printf("%s %s %s %s %d %s %s", r.RemoteAddr, r.Method, r.URL, r.Proto, o.status, r.Referer(), r.UserAgent())
	})
}

type responseWrapper struct {
	http.ResponseWriter
	status      int
	written     int64
	wroteHeader bool
}

func (o *responseWrapper) Write(p []byte) (n int, err error) {
	if !o.wroteHeader {
		o.WriteHeader(http.StatusOK)
	}
	n, err = o.ResponseWriter.Write(p)
	o.written += int64(n)
	return
}

func (o *responseWrapper) WriteHeader(code int) {
	o.ResponseWriter.WriteHeader(code)
	if o.wroteHeader {
		return
	}
	o.wroteHeader = true
	o.status = code
}
