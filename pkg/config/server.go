package config

import (
	"github.com/spf13/viper"
	"time"
)

func init() {
	viper.BindEnv("SERVER_ADDR")
	viper.BindEnv("SERVER_READ_TIMEOUT")
	viper.BindEnv("SERVER_WRITE_TIMEOUT")
}

func ServerAddr() string {
	return viper.GetString("SERVER_ADDR")
}

func ServerReadTimeout() time.Duration {
	return viper.GetDuration("SERVER_READ_TIMEOUT")
}

func ServerWriteTimeout() time.Duration {
	return viper.GetDuration("SERVER_WRITE_TIMEOUT")
}
