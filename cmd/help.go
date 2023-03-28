/*
Copyright © 2023 werniq qniwwwersss@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	allCommandsInfo = `
		Avaliable Commands:
			start       	-> launches decentralized application 
			help		-> lists information about all avaliable commands
			deploy 		-> deploys smart contract to mainnet\testnet, with public and private key, which user specifies
			generateABI 	-> compiles Solidity Smart Contract from folder [Contracts], and generates ABI from it
			test		-> runs tests from [Testing] directory
				`
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "list all available commands",
	Long:  `retrieve info about all commands, that can be ran in this application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(allCommandsInfo)
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
