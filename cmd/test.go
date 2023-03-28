package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:     "",
	Short:   "",
	Long:    "",
	Example: "",
	Run: func(cmd *cobra.Command, args []string) {
		Test(cmd, args)
	},
}

func init() {
	rootCmd := &cobra.Command{Use: "go-dapp"}
	rootCmd.AddCommand(testCmd)
}

func Test(cmd *cobra.Command, args []string) {
	fmt.Println("test function is running")
}
