package main

import (
	"context"
	"database/sql"
	"flag"
	"github.com/go-chi/cors"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"

	"github.com/ErnestsKalnins/0xff/internal/environment"
	"github.com/ErnestsKalnins/0xff/internal/feature"
	"github.com/ErnestsKalnins/0xff/internal/project"
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

	projectStore := project.NewStore(db)
	projetService := project.NewService(projectStore)
	projectHandler := project.NewHandler(projetService)

	featureStore := feature.NewStore(db)
	featureService := feature.NewService(featureStore)
	featureHandler := feature.NewHandler(featureService)

	environmentStore := environment.NewStore(db)
	environmentService := environment.NewService(environmentStore)
	environmentHandler := environment.NewHandler(environmentService)

	r := chi.NewRouter()

	r.Use(
		hlog.NewHandler(logger),
		cors.AllowAll().Handler,
	)

	r.Route("/api/v1/projects", func(r chi.Router) {
		r.Get("/", projectHandler.ListProjects)
		r.Post("/", projectHandler.SaveProject)

		r.Route("/{projectId}", func(r chi.Router) {
			r.Get("/", projectHandler.GetProject)
			r.Delete("/", projectHandler.DeleteProject)

			r.Route("/features", func(r chi.Router) {
				r.Get("/", featureHandler.ListFeatures)
				r.Post("/", featureHandler.SaveFeature)

				r.Route("/{featureId}", func(r chi.Router) {
					r.Get("/", featureHandler.GetFeature)
					r.Delete("/", featureHandler.DeleteFeature)
				})
			})

			r.Route("/environments", func(r chi.Router) {
				r.Get("/", environmentHandler.ListEnvironments)
				r.Post("/", environmentHandler.SaveEnvironment)

				r.Route("/{environmentId}", func(r chi.Router) {
					r.Get("/", environmentHandler.GetEnvironment)
					r.Delete("/", environmentHandler.DeleteEnvironment)

					r.Route("/features", func(r chi.Router) {
						r.Get("/", featureHandler.ListFeatureStates)

						r.Route("/{featureId}", func(r chi.Router) {
							r.Put("/", featureHandler.SetFeatureState)
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
