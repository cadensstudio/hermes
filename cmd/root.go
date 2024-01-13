package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

// var ApiKey string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hermes",
	Short: "Hermes simplifies the process of downloading Google Font files.",
	Long: `Hermes simplifies the process of downloading web-optimized Google Font files
in the WOFF2 format and generates the necessary CSS code to easily integrate
the downloaded fonts into your project.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
