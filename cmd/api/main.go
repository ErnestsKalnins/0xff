package main

import (
	"context"
	"database/sql"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/spf13/viper"

	"github.com/ErnestsKalnins/0xff/feature/api"
	"github.com/ErnestsKalnins/0xff/feature/data/sqlite"
	"github.com/ErnestsKalnins/0xff/pkg/config"
)

func main() {
	var (
		envFile = flag.String("env-file", "", "Path to env file containing configuration.")
		logger  = zerolog.New(os.Stderr).With().Timestamp().Logger()
	)

	flag.Parse()

	if *envFile != "" {
		viper.SetConfigFile(*envFile)
		if err := viper.ReadInConfig(); err != nil {
			logger.Fatal().
				Err(err).
				Msg("read configuration")
		}
	}

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

	store := sqlite.New(db)
	handler := api.NewHandler(store)

	r := chi.NewRouter()

	r.Use(
		hlog.NewHandler(logger),
		cors.AllowAll().Handler,
	)

	r.Route("/api/v1/projects", func(r chi.Router) {
		r.Get("/", handler.ListProjects)
		r.Post("/", handler.SaveProject)

		r.Route("/{projectId}", func(r chi.Router) {
			r.Get("/", handler.GetProject)
			r.Delete("/", handler.DeleteProject)

			r.Route("/features", func(r chi.Router) {
				r.Get("/", handler.ListFeatures)
				r.Post("/", handler.SaveFeature)

				r.Route("/{featureId}", func(r chi.Router) {
					r.Get("/", handler.GetFeature)
					r.Delete("/", handler.DeleteFeature)
				})
			})

			r.Route("/environments", func(r chi.Router) {
				r.Get("/", handler.ListEnvironments)
				r.Post("/", handler.SaveEnvironment)

				r.Route("/{environmentId}", func(r chi.Router) {
					r.Get("/", handler.GetEnvironment)
					r.Delete("/", handler.DeleteEnvironment)

					r.Route("/features", func(r chi.Router) {
						r.Get("/", handler.ListFeatureStates)

						r.Route("/{featureId}", func(r chi.Router) {
							r.Put("/", handler.SetFeatureState)
						})
					})
				})
			})
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

	logger.Info().
		Str("addr", server.Addr).
		Msg("starting HTTP server")

	if err := server.ListenAndServe(); err != nil {
		logger.Fatal().
			Err(err).
			Msg("serve HTTP")
	}

	<-idleConnsClosed
}
