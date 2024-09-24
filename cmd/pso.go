package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/Friends-Of-Noso/NosoData-Go/legacy"
	"github.com/spf13/cobra"
)

const (
	cPSOFilename = "psos.dat"
)

// psoCmd represents the pso command
var psoCmd = &cobra.Command{
	Use:   "pso",
	Short: "Display the contents of a PSO file",
	Run:   psoRun,
}

func init() {
	displayCmd.AddCommand(psoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// psoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// psoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func psoRun(cmd *cobra.Command, args []string) {
	var filename = cPSOFilename
	if nosoFolder != "" {
		filename = filepath.Join(nosoFolder, "NOSODATA", filename)
	} else {
		filename = filepath.Join(".", "NOSODATA", filename)
	}
	psos := legacy.LegacyPSO{}
	err := psos.ReadFromFile(filename)
	cobra.CheckErr(err)
	if displayInJSON {
		fmt.Println(psos.AsJSON())
		return
	}
	fmt.Println("--- PSO ---")
	fmt.Println("Block:", psos.Block)
	fmt.Printf("MN Locks(%d):\n", psos.MNLockCount)
	for i, mli := range psos.MNLocks {
		fmt.Println("  Position:", i)
		fmt.Printf("      Address: '%s'\n", mli.Address.GetString())
		fmt.Println("       Expire:", mli.Expire, "seconds")
	}
	fmt.Printf("PSO Count(%d):\n", psos.PSOCount)
}
