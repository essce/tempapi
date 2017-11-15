package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	handler "github.com/essce/tempapi/http"
	"github.com/essce/tempapi/postgres"
	"github.com/go-chi/chi"
)

func main() {
	pg := postgres.New()

	h := newHandler(&pg)
	s := newServer(":8080", h)

	if err := s.ListenAndServe(); err != nil {
		panic("server exited")
	}

	fmt.Println("ready to serve on :8080")

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)
	select {
	case sig := <-gracefulStop:
		fmt.Printf("received os signal, signal: %s\n", sig.String())
	}

	s.Shutdown(context.Background())
	pg.Close()

}

// ListenAndServe starts a server using the provided TCP.
func ListenAndServe(s *http.Server, keepAlive time.Duration) error {
	return s.ListenAndServe()
}

func newServer(addr string, h http.Handler) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: h,
	}
}

func newHandler(pg *postgres.Postgres) http.Handler {
	h := handler.Handler{
		ReadingStore: pg,
	}

	m := chi.NewMux()
	m.Get("/", h.Version)
	m.Post("/reading", h.InsertReading)
	m.Get("/reading", h.ListReadings)
	return m
}
