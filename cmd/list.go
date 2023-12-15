package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
)

type FontList struct {
	Items []struct {
		Family string `json:"family"`
	} `json:"items"`
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the 10 most trending Google Fonts",
	Long: `Lists the 10 most trending Google Fonts,
providing inspiration for your next project.`,
	Run: func(cmd *cobra.Command, args []string) {
		// try grabbing key from .env file, if it exists
		key := viper.Get("GFONTS_KEY")
		if key == nil {	
			// if no .env, grab key from cmd flag
			key = viper.GetString("key")
			if len(fmt.Sprint(key)) < 1 {
				fmt.Println(`Error: required flag "key" not set`)
				os.Exit(1)
			}
		}
		url := "https://www.googleapis.com/webfonts/v1/webfonts?key=" + fmt.Sprint(key) + "&sort=trending"

		// Make the GET request
		res, err := http.Get(url)
		if err != nil {
			fmt.Println("Error: failed to create connection to remote host", err)
			os.Exit(1)
		}
		defer res.Body.Close()

		// check response and handle errors
		if res.StatusCode == 200 {
			// Read the response body
			body, err := io.ReadAll(res.Body)
			if err != nil {
				fmt.Println("Error: Could not read response body", err)
				os.Exit(1)
			}

			// parse the response body into the FontList object struct
			var listResponse FontList
			err = json.Unmarshal(body, &listResponse)
			if err != nil {
				fmt.Println("Error: could not parse json response", err)
				os.Exit(1)
			}

			// Grab the first 10 items
			for i, font := range listResponse.Items {
				if i >= 10 {
					break
				}
				parsedFontFamily := parseFontFamily(font.Family)
				fontUrl := "https://fonts.google.com/?query=" + parsedFontFamily
				fmt.Println(font.Family + ": " + fontUrl)
			}
		} else if res.StatusCode == 400 {
			fmt.Println("Error: Could not complete request")
			os.Exit(1)
			return
		} else {
			fmt.Println("An unexpected error occured")
			os.Exit(1)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().StringVarP(&ApiKey, "key", "k", "", "Your Google Fonts API Key")
	viper.BindPFlag("key", listCmd.PersistentFlags().Lookup("key"))
}
