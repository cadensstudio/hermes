/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the 10 most trending Google Fonts",
	Long: `Lists the 10 most trending Google Fonts,
providing inspiration for your next project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
