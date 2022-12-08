package test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("get working directory: %s", err))
	}
	var (
		dirs          = strings.Split(wd, string(filepath.Separator))
		relConfigPath string
	)
	for i := len(dirs) - 1; i >= 0; i-- {
		if dirs[i] == "0xff" {
			break
		}
		relConfigPath += "../"
	}
	relConfigPath += "test.env"
	viper.SetConfigFile(relConfigPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("read test config: %s", err))
	}
}
