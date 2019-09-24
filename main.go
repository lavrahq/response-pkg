package main

import (
	"github.com/lavrahq/response-pkg/sdk"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")

	viper.AddConfigPath("/config")
	viper.AddConfigPath("config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		// fmt.Println("Failed to read config. Terminating process.")

		// os.Exit(1)
	}
}

func main() {
	sdk.FetchPackages()
}
