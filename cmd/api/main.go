package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"

	"github.com/ErnestsKalnins/0xff/internal/feature"
	"github.com/ErnestsKalnins/0xff/pkg/config"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	db, err := sql.Open("sqlite3", config.DSN())
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("open database connection")
	}

	if err := db.Ping(); err != nil {
		logger.Fatal().
			Err(err).
			Msg("ping database")
	}

	store := feature.NewStore(db)
	service := feature.NewService(store)
	handler := feature.NewHandler(service)

	r := chi.NewRouter()

	r.Use(hlog.NewHandler(logger))

	r.Route("/features", func(r chi.Router) {
		r.Get("/", handler.ListFeatures)
		r.Post("/", handler.SaveFeature)

		r.Route("/{featureId}", func(r chi.Router) {
			r.Get("/", handler.GetFeature)
			r.Delete("/", handler.DeleteFeature)
		})
	})

	server := http.Server{
		Addr:         config.ServerAddr(),
		Handler:      r,
		ReadTimeout:  config.ServerReadTimeout(),
		WriteTimeout: config.ServerWriteTimeout(),
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint

		if err := server.Shutdown(context.Background()); err != nil {
			logger.Error().
				Err(err).
				Msg("shutdown HTTP server")
		}
		close(idleConnsClosed)
	}()

	if err := server.ListenAndServe(); err != nil {
		logger.Fatal().
			Err(err).
			Msg("serve HTTP")
	}

	<-idleConnsClosed
}
