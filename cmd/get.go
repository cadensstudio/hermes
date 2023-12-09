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
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// create Axes struct to handle extra key for variable fonts
type Axes struct {
	Tag   string `json:"tag"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

// extract font file path url from Google Fonts API JSON response
type Font struct {
	Items []struct {
		Family   string            `json:"family"`
		Variants []string          `json:"variants"`
		Files    map[string]string `json:"files"`
		Axes     []*Axes           `json:"axes,omitempty"`
	} `json:"items"`
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [FONT FAMILY] [FLAGS]",
	Short: "Download web-optimized font files for a specified font family.",
	Long: `
Download the specified font family in WOFF2 format.
By default, if a single variable format is available, it will be downloaded;
otherwise, each individual font weight file will be downloaded in the current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		} else {
			fontFamily := args[0]
			parsedFontFamily := parseFontFamily(fontFamily)
			fontResponse := getFontUrl(parsedFontFamily)
			if len(fontResponse.Items) >= 1 {
				donwloadFont(fontResponse)
			}
		}
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

func parseFontFamily(fontFamily string) (parsedFontFamily string) {
	// convert font input to lowercase
	fontFamily = cases.Lower(language.Und).String(fontFamily)
	// convert first letter of each word to uppercase
	fontFamily = cases.Title(language.Und).String(fontFamily)
	// replace spaces with + for url formatting
	for _, char := range fontFamily {
		if char == ' ' {
			parsedFontFamily += "+"
		} else {
			parsedFontFamily += string(char)
		}
	}
	return parsedFontFamily
}

func getFontUrl(fontFamily string) (fontResponse Font) {
	key := viper.Get("GFONTS_KEY")
	url := "https://www.googleapis.com/webfonts/v1/webfonts?key=" + fmt.Sprint(key) + "&family=" + fontFamily + "&capability=WOFF2&capability=VF"

	// Make the GET request
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making web request:", err)
		return
	}
	defer res.Body.Close()

	// check response and handle errors
	if res.StatusCode == 200 {
		// Read the response body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		// parse the response body into the Font object struct
		var fontResponse Font
		err = json.Unmarshal(body, &fontResponse)
		if err != nil {
			fmt.Println("Error parsing json response", err)
		}
		return fontResponse
	} else if res.StatusCode == 400 {
		fmt.Println("400 Error: Could not complete request")
		return
	} else if res.StatusCode == 500 {
		fmt.Println("500 Error: Could not find specified font:", fontFamily)
		return
	} else {
		fmt.Println("An unexpected error occured")
		return
	}
}

func donwloadFont(fontResponse Font) {
	var hasVariable bool
	fontFamily := fontResponse.Items[0].Family
	fontFiles := fontResponse.Items[0].Files

	if len(fontResponse.Items[0].Axes) == 0 {
		hasVariable = false
		fmt.Println("Variable font file not available.")
		fmt.Println("Downloading font files individually...")
	} else {
		hasVariable = true
		if len(fontFiles) == 1 {
			fmt.Println("Downloading variable font file...")
		} else {
			fmt.Println("Downloading variable font file(s)...")
		}
	}

	for variant, url := range fontFiles {
		// Make the GET request for each variant
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			continue // Skip to the next variant if an error occurs
		}
		defer res.Body.Close()

		// Create the font file on the local system
		filepath := fmt.Sprintf("%s_%s.woff2", fontFamily, variant)
		out, err := os.Create(filepath)
		if err != nil {
			fmt.Println(err)
			continue // Skip to the next variant if an error occurs
		}
		defer out.Close()

		// Write the downloaded file to the local file
		_, err = io.Copy(out, res.Body)
		if err != nil {
			fmt.Println(err)
			continue // Skip to the next variant if an error occurs
		}
		fmt.Printf("%s successfully downloaded!\n", filepath)
	}
	fmt.Println("\nNext steps: Copy the following CSS rules wherever you would like to use your font!")
	printCssStrings(fontResponse, hasVariable)
}

func printCssStrings(fontResponse Font, hasVariable bool) {
	var newCssString string

	if hasVariable {
		var startWeight string
		var endWeight string
		for _, axis := range fontResponse.Items[0].Axes {
			if axis.Tag == "wght" {
				startWeight = strconv.Itoa(axis.Start)
				endWeight = strconv.Itoa(axis.End)
			}
		}
		for _, variant := range fontResponse.Items[0].Variants {
			newCssString = `
@font-face {
  font-family: '` + fontResponse.Items[0].Family + `';
  font-style: ` + variant + `;
  font-weight: ` + startWeight + `-` + endWeight + `;
  src: url('./` + fontResponse.Items[0].Family + `_` + variant + `.woff2') format('woff2');
}`
			fmt.Println(newCssString)
		}
	} else {
		for _, variant := range fontResponse.Items[0].Variants {
			var fontStyle string
			var fontWeight string
			if variant == "regular" || variant == "italic" {
				fontStyle = variant
				fontWeight = "400"
			} else {
				fontStyle = variant[3:]
				if len(fontStyle) == 0 {
					fontStyle = "regular"
				}
				fontWeight = variant[:3]
			}
			newCssString = `
@font-face {
  font-family: '` + fontResponse.Items[0].Family + `';
  font-style: ` + fontStyle + `;
  font-weight: ` + fontWeight + `;
  src: url('./` + fontResponse.Items[0].Family + `_` + variant + `.woff2') format('woff2');
}`
			fmt.Println(newCssString)
		}
	}
}
