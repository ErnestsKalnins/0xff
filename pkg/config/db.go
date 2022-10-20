package config

import "github.com/spf13/viper"

func init() {
	viper.BindEnv("DSN")
}

// DSN retrieves the database DSN from system env.
func DSN() string {
	return viper.GetString("DSN")
}
