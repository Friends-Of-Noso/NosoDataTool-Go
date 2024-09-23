package cmd

import (
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/Friends-Of-Noso/NosoData-Go/legacy"
	"github.com/Friends-Of-Noso/NosoData-Go/utils"
)

const (
	cBlockFilenameMask = "%d.blk"
)

// blockCmd represents the block command
var blockCmd = &cobra.Command{
	Use:   "block [number]",
	Short: "This command allows to query blocks by block number",
	Args:  cobra.ExactArgs(1),
	Run:   blockRun,
}

func init() {
	rootCmd.AddCommand(blockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// blockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// blockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func blockRun(cm *cobra.Command, args []string) {
	num, err := strconv.ParseInt(args[0], 10, 64)
	cobra.CheckErr(err)
	filename := fmt.Sprintf(cBlockFilenameMask, num)
	if nosoFolder != "" {
		filename = filepath.Join(nosoFolder, "NOSODATA", "BLOCKS", filename)
	} else {
		filename = filepath.Join(".", "NOSODATA", "BLOCKS", filename)
	}
	block := legacy.LegacyBlock{}
	err = block.ReadFromFile(filename)
	cobra.CheckErr(err)
	fmt.Println("--- Block ---")
	fmt.Println("Number:           ", block.Number)
	fmt.Println("Time Start:       ", time.Unix(block.TimeStart, 0))
	fmt.Println("Time End:         ", time.Unix(block.TimeEnd, 0))
	fmt.Println("Time Total:       ", block.TimeTotal, "seconds")
	fmt.Println("Time Last 20:     ", block.TimeLast20, "seconds")
	fmt.Println("Transaction Count:", block.TransactionsCount)
	fmt.Println("Difficulty:       ", block.Difficulty)
	fmt.Printf("Target Hash:      '%s'\n", block.TargetHash.GetString())
	fmt.Printf("Solution:         '%s'\n", block.Solution.GetString())
	fmt.Printf("Last Block Hash:  '%s'\n", block.LastBlockHash.GetString())
	fmt.Printf("Miner:            '%s'\n", block.Miner.GetString())
	fmt.Println("Fee:             ", utils.ToNoso(block.Fee))
	fmt.Println("Reward:          ", utils.ToNoso(block.Reward))

	if block.TransactionsCount > 0 {
		fmt.Printf("Transactions(%d):\n", block.TransactionsCount)
		var n int32
		for n = 0; n < block.TransactionsCount; n++ {
			fmt.Printf("  OrderID: '%s'\n", block.Transactions[n].OrderID.GetString())
			fmt.Printf("      TransferID:     '%s'\n", block.Transactions[n].TransferID.GetString())
			fmt.Println("      Block:         ", block.Transactions[n].Block)
			fmt.Println("      Order lines:   ", block.Transactions[n].OrderLinesCount)
			fmt.Printf("      Order type:     '%s'\n", block.Transactions[n].OrderType.GetString())
			fmt.Println("      Timestamp:     ", time.Unix(block.Transactions[n].TimeStamp, 0))
			fmt.Printf("      Reference:      '%s'\n", block.Transactions[n].Reference.GetString())
			fmt.Println("      Transfer Index:", block.Transactions[n].TransferIndex)
			fmt.Printf("      Sender:         '%s'\n", block.Transactions[n].Sender.GetString())
			fmt.Printf("      Address:        '%s'\n", block.Transactions[n].Address.GetString())
			fmt.Printf("      Receiver:       '%s'\n", block.Transactions[n].Receiver.GetString())
			fmt.Println("      Fee:           ", utils.ToNoso(block.Transactions[n].AmountFee))
			fmt.Println("      Value:         ", utils.ToNoso(block.Transactions[n].AmountTransfer))
			fmt.Printf("      Signature:      '%s'\n", block.Transactions[n].Signature.GetString())
		}
	} else {
		fmt.Println("No transactions")
	}

	if block.ProofOfStakeRewardCount > 0 {
		fmt.Printf("PoS rewards(%d):\n", block.ProofOfStakeRewardCount)
		fmt.Println("  Amount:", utils.ToNoso(block.ProofOfStakeRewardAmount))
		var n int32
		for n = 0; n < block.ProofOfStakeRewardCount; n++ {
			fmt.Printf("  Address: '%s'\n", block.ProofOfStakeRewardAddresses[n].GetString())
		}
	} else {
		fmt.Println("No PoS rewards")
	}

	if block.MasterNodeRewardCount > 0 {
		fmt.Printf("MN rewards(%d):\n", block.MasterNodeRewardCount)
		fmt.Println("  Amount:", utils.ToNoso(block.MasterNodeRewardAmount))
		var n int32
		for n = 0; n < block.MasterNodeRewardCount; n++ {
			fmt.Printf("  Address: '%s'\n", block.MasterNodeRewardAddresses[n].GetString())
		}
	} else {
		fmt.Println("No MN rewards")
	}
}
