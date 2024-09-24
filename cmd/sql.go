package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	dbEngine        string
	allowedDbValues = []string{
		"SQLite",
		"MySQL",
		"MariaDB",
		"PostgreSQL",
	}
)

// sqlCmd represents the sql command
var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "Exports all data to SQL",
	Run:   sqlRun,
}

func init() {
	exportCmd.AddCommand(sqlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sqlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sqlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	sqlCmd.Flags().StringVarP(&dbEngine, "database-engine", "e", "SQLite", "database engine (options: \""+strings.Join(allowedDbValues, "\", \"")+"\")")

}

func sqlRun(cmd *cobra.Command, args []string) {
	// Validate the value provided by the user
	if !contains(allowedDbValues, dbEngine) {
		fmt.Println("invalid value: " + dbEngine + ". Allowed values are: \"" + strings.Join(allowedDbValues, "\", \"") + "\"")
		return
	}

	// If validation passes, print the value
	fmt.Println("Selected database:", dbEngine)
}

// contains checks if a value exists in a list
func contains(list []string, value string) bool {
	for _, v := range list {
		if strings.EqualFold(v, value) {
			return true
		}
	}
	return false
}
