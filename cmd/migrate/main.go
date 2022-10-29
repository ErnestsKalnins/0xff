package main

import (
	"database/sql"
	"flag"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/ErnestsKalnins/0xff/pkg/config"
	"github.com/ErnestsKalnins/0xff/pkg/migrate"
)

func main() {
	envFile := flag.String("env-file", "", "Path to env file containing configuration.")

	flag.Parse()

	if *envFile != "" {
		viper.SetConfigFile(*envFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal().
				Err(err).
				Msg("failed to read configuration")
		}
	}

	db, err := sql.Open("sqlite3", config.DSN())
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to open connection to database")
	}

	if err := db.Ping(); err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to ping database")
	}

	if err := migrate.Migrate(db); err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to run migrations")
	}
}
