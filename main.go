/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"hermes/cmd"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	cmd.Execute()
}
