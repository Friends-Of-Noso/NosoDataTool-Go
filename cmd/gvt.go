package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/Friends-Of-Noso/NosoData-Go/legacy"
	"github.com/spf13/cobra"
)

const (
	cGVTFilename = "gvts.psk"
)

// gvtCmd represents the gvt command
var gvtCmd = &cobra.Command{
	Use:   "gvt",
	Short: "Display the contents of a GVT file",
	Run:   gvtRun,
}

func init() {
	displayCmd.AddCommand(gvtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gvtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gvtCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func gvtRun(cmd *cobra.Command, args []string) {
	var filename = cGVTFilename
	if nosoFolder != "" {
		filename = filepath.Join(nosoFolder, "NOSODATA", filename)
	} else {
		filename = filepath.Join(".", "NOSODATA", filename)
	}
	gvts := legacy.LegacyGVT{}
	err := gvts.ReadFromFile(filename)
	cobra.CheckErr(err)
	if displayInJSON {
		fmt.Println(gvts.AsJSON())
		return
	}
	fmt.Println("--- GVT ---")
	for i, e := range gvts.Entries {
		fmt.Println("Position:", i)
		fmt.Printf("    Number:  '%s'\n", e.Number.GetString())
		fmt.Printf("    Owner:   '%s'\n", e.Owner.GetString())
		fmt.Printf("    Hash:    '%s'\n", e.Hash.GetString())
		fmt.Println("    Control:", e.Control)
	}
}
