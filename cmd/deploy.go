package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-solidity-gen/logger"
	"os/exec"
)

var (
	deployCmd = &cobra.Command{
		Use:        "deploy [flags] [...contractAddress]",
		Aliases:    nil,
		SuggestFor: nil,
		Short:      "deploys smart contract",
		Long:       "Deploys ethereum smart contract to user specified network(mainnet, goerli testnet, etc.). \n Requires public, private signer keys, and contract name. ",
		Args:       cobra.ExactArgs(1),
		Run: func(c *cobra.Command, args []string) {
			// to deploy contract, firstly, I need to generate application binary interface of this contract, and get
			// the bytecode. secondly, I need to get from user his private and public keys.
			// last, I need to ensure, that user have submitted all construct arguments

			// specify dockerfile location
			cmd := exec.Command("docker", "run", "--build-arg", fmt.Sprintf("contractAddress=%s", args[0]))
			err := cmd.Run()
			if err != nil {
				logger.Logger().Printf("error running deploy docker command: %v\n", err)
				return
			}
		},
	}
	contractName = ""
)

func init() {
	rootCmd := &cobra.Command{
		Use: "go-dapp",
	}
	deployCmd.Flags().StringVar(&contractName, "contractName", "s", "Address of contract, that need to be deployed")
	err := deployCmd.MarkFlagRequired("contractName")

	if err != nil {
		logger.Logger().Println("error marking flags required: %v\n", err)
		return
	}

	rootCmd.AddCommand(deployCmd)
	if err := rootCmd.Execute(); err != nil {
		logger.Logger().Printf("error executing deploy command: %v\n", err)
	}
}
