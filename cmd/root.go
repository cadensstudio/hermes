package cmd

import (
	"os"
	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hermes [COMMAND]",
	Short: "Hermes simplifies font downloading from Google Fonts.",
	Long: `
Tailored for developers and designers, Hermes automates the process
of downloading web-optimized font files from Google Fonts.

Key Features:
- Effortlessly download fonts in the highly efficient WOFF2 format, ideal for web use.
- Automatic retrieval of variable font formats when available, ensuring flexibility in design.
- Empower your web projects with the 'list' command, revealing the top 10 trending Google Fonts.
- Expedites font acquisition, allowing developers to self-host fonts and significantly boost website load speeds.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hermes.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


