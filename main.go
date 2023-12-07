/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/spf13/viper"
	"hermes/cmd"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	cmd.Execute()
}