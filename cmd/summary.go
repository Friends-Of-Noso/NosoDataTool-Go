package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/Friends-Of-Noso/NosoData-Go/legacy"
	"github.com/Friends-Of-Noso/NosoData-Go/utils"
	"github.com/spf13/cobra"
)

const (
	cSummaryFilename = "sumary.psk"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Display the contents of a summary file",
	Run:   summaryRun,
}

func init() {
	displayCmd.AddCommand(summaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// summaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// summaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func summaryRun(cmd *cobra.Command, args []string) {
	var filename = cSummaryFilename
	if nosoFolder != "" {
		filename = filepath.Join(nosoFolder, "NOSODATA", filename)
	} else {
		filename = filepath.Join(".", "NOSODATA", filename)
	}
	summary := legacy.LegacySummary{}
	err := summary.ReadFromFile(filename)
	cobra.CheckErr(err)
	if displayInJSON {
		fmt.Println(summary.AsJSON())
		return
	}
	fmt.Println("--- Summary ---")
	for i, a := range summary.Accounts {
		fmt.Println("Position:", i)
		fmt.Printf("    Hash:           '%s'\n", a.Hash.GetString())
		fmt.Printf("    Custom:         '%s'\n", a.Custom.GetString())
		fmt.Println("    Balance:       ", utils.ToNoso(a.Balance))
		fmt.Println("    Score:         ", utils.ToNoso(a.Score))
		fmt.Println("    Last Operation:", utils.ToNoso(a.LastOperation))
	}
}
