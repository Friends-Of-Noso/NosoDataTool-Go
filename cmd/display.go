package cmd

import (
	"github.com/spf13/cobra"
)

var (
	displayInJSON bool
)

// displayCmd represents the display command
var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "Displays various file contents in text format of JSON format",
}

func init() {
	rootCmd.AddCommand(displayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// displayCmd.PersistentFlags().String("foo", "", "A help for foo")

	displayCmd.PersistentFlags().BoolVarP(&displayInJSON, "json", "j", false, "Outputs in JSON")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// displayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
