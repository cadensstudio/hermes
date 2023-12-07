/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Font struct {
	Items []struct {
		Files struct {
			Filepath string `json:"regular"`
		} `json:"files"`
	} `json:"items"`
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Download a specific font",
	Long: `Donwload a specific Google font in web optimized WOFF2 format into the current directory by defualt.`,
	Run: func(cmd *cobra.Command, args []string) {
		getFont()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getFont() {
	key:= viper.Get("GFONTS_KEY")
	url := "https://www.googleapis.com/webfonts/v1/webfonts?key=" + fmt.Sprint(key) + "&family=Roboto&capability=WOFF2&capability=VF"

	// Make the GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var font Font
	err = json.Unmarshal(body, &font)

	if err != nil {
		panic(err)
	}
	// Print the response body
	regularURL := font.Items[0].Files.Filepath
	fmt.Println(regularURL)
}