/*
Copyright © 2024 Ethan Holen ethanholen@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/EthanHolen/orthocli/orthoapi"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "orthocli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		orthoResponse, err := orthoapi.MakeRequest("https://orthocal.info/api/gregorian/")

		if err != nil {
			fmt.Fprintf(os.Stderr, "error retrieving orthoResponse: %e\n", err)
			os.Exit(1)
		}

		// fmt.Printf("%d %d   %s\n", orthoResponse.PaschaDistance, orthoResponse.Year, orthoResponse.Titles[0]) // TODO: do some safety checking here on titles
		fmt.Printf(`
🗓️  Date: (%d-%d-%d) %s

🍽️  Fasting: %d %s
🍽️  Feast: %d %s

📚 Readings:

👑 Saints:

`, orthoResponse.Month,
			orthoResponse.Day,
			orthoResponse.Year,
			orthoResponse.Titles[0],
			orthoResponse.FastLevel,
			orthoResponse.FastLevelDescription,
			orthoResponse.FeastLevel,
			orthoResponse.FeastLevelDescription,
		)

	},
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.orthocli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
