package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-solidity-gen/logger"
	"os"
	"os/exec"
)

var (
	generateABICmd = &cobra.Command{
		Use:   "go-dapp",
		Short: "generates abi for contract",
		Long:  "creates new folder, with abi of contract, provided in arguments",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if _, err := os.Stat(fmt.Sprintf("./Contracts/%s", args[0])); err != nil {
				logger.InfoLogger("Contract directory does not exist. Please, run <start> command, to initialize decentralized applciation")
				return
			}
			exec.Command("go get", "-d", "github.com/ethereum/go-ethereum")
			exec.Command("solc", "--abi", "--bin", fmt.Sprintf("%s.sol", args[0]), "-o", "build/")
		},
	}
)

func main() {
	rootCmd := &cobra.Command{Use: "go-dapp"}
	rootCmd.AddCommand(generateABICmd)
}
