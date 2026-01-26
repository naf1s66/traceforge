package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"traceforge/api/internal/router"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	router.Mount(r)

	addr := getenv("ADDR", ":8080")
	log.Info().Msgf("TraceForge API listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal().Err(err).Msg("server crashed")
	}
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
