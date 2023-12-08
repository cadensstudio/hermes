/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"golang.org/x/text/cases"
  "golang.org/x/text/language"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// extract font file path url from Google Fonts API JSON response
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
		// declare default font or grab from command args
		var fontFamily = "Roboto"
		if len(args) >= 1 && args[0] != "" {
			fontFamily = args[0]
		}

		parsedFontFamily := parseFontFamily(fontFamily)
		fmt.Println(parsedFontFamily)

		fontUrl := getFontUrl(parsedFontFamily)
		donwloadFont(parsedFontFamily, fontUrl)
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

func getFontUrl(fontFamily string) (fontUrl string) {
	key:= viper.Get("GFONTS_KEY")
	url := "https://www.googleapis.com/webfonts/v1/webfonts?key=" + fmt.Sprint(key) + "&family=" + fontFamily + "&capability=WOFF2&capability=VF"

	// Make the GET request
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var font Font
	err = json.Unmarshal(body, &font)

	if err != nil {
		fmt.Println(err)
	}
	// Print the response body
	fontUrl = font.Items[0].Files.Filepath
	fmt.Println(fontUrl)
	return fontUrl
}

func donwloadFont(fontFamily string, url string) {
	filepath := fontFamily + ".woff2"

	// Donwload the font
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	// Create the font file on the local system
	out, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()

	// Write the downloaded file to local file
	_, err = io.Copy(out, res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out.Name() + " successfully downloaded!")
}

func parseFontFamily(fontFamily string) (parsedFontFamily string) {
	fmt.Println(fontFamily)
	// convert font input to lowercase
	parsedFontFamily = cases.Lower(language.Und).String(fontFamily)
	// convert first letter of each word to uppercase
	parsedFontFamily = cases.Title(language.Und).String(parsedFontFamily)
	// replace spaces with + for url formatting
	for _, char := range parsedFontFamily {
		if char == ' ' {
			parsedFontFamily += "+"
		}
	}
	return parsedFontFamily
}