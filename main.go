package main

import (
	"github.com/cadensstudio/hermes/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	cmd.Execute()
}
