package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/Friends-Of-Noso/NosoData-Go/legacy"
	"github.com/Friends-Of-Noso/NosoData-Go/utils"
	"github.com/spf13/cobra"
)

const (
	cWalletFilename = "wallet.pkw"
)

// walletCmd represents the wallet command
var walletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: walletRun,
}

func init() {
	rootCmd.AddCommand(walletCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// walletCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// walletCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func walletRun(cmd *cobra.Command, args []string) {
	var filename = cWalletFilename
	if nosoFolder != "" {
		filename = filepath.Join(nosoFolder, "NOSODATA", filename)
	} else {
		filename = filepath.Join(".", "NOSODATA", "BLOCKS", filename)
	}
	wallet := legacy.LegacyWallet{}
	err := wallet.ReadFromFile(filename)
	cobra.CheckErr(err)
	fmt.Println("--- Wallet ---")
	for i, a := range wallet.Accounts {
		fmt.Println("Position:", i)
		fmt.Printf("    Hash: '%s'\n", a.Hash.GetString())
		fmt.Printf("    Custom:         '%s'\n", a.Custom.GetString())
		fmt.Printf("    Pub key:        '%s'\n", a.PublicKey.GetString())
		fmt.Printf("    Priv key:       '%s'\n", a.PrivateKey.GetString())
		fmt.Println("    Balance:       ", utils.ToNoso(a.Balance))
		fmt.Println("    Pending:       ", utils.ToNoso(a.Pending))
		fmt.Println("    Score:         ", utils.ToNoso(a.Score))
		fmt.Println("    Last Operation:", utils.ToNoso(a.LastOperation))
	}
}
