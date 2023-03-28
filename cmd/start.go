/*
Copyright © 2023 werniq qniwwwersss@gmail.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"go-solidity-gen/logger"
	"os"
)

var (
	sampleContractTemplate = "" +
		"// SPDX-License-Identifier: MIT\n" +
		"pragma solidity ^0.8.0;\n\n" +
		"contract SimpleToken {\n" +
		"\tstring public name;\n" +
		"\tstring public symbol;\n" +
		"\tuint public totalSupply;\n" +
		"\tmapping(address => uint) public balanceOf;\n\n" +
		"\tconstructor(string memory _name, string memory _symbol, uint _totalSupply) {\n" +
		"\t\tname = _name;\n" +
		"\t\tsymbol = _symbol;\n" +
		"\t\ttotalSupply = _totalSupply;\n" +
		"\t\tbalanceOf[msg.sender] = _totalSupply;\n" +
		"\t}\n" +
		"}"

	testContractSample = `package main

	import (
		"context"
		"fmt"
		"math/big"
		"testing"
		"github.com/ethereum/go-ethereum/accounts/abi/bind"
		"github.com/ethereum/go-ethereum/common"
		"github.com/ethereum/go-ethereum/crypto"
		"github.com/ethereum/go-ethereum/ethclient"
		"github.com/stretchr/testify/assert"
	)

	func TestSimpleToken(t *testing.T) {
		// Connect to a local Ethereum client
		client, err := ethclient.Dial("http://localhost:8545")
		if err != nil {
			t.Fatal(err)
		}
	
		// Generate a new account to use as the contract owner
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		auth := bind.NewKeyedTransactor(privateKey)

		// Deploy the SimpleToken contract
		name := "MyToken"
		symbol := "MTK"
		totalSupply := big.NewInt(1000000)
		_, tx, token, err := DeploySimpleToken(auth, client, name, symbol, totalSupply)
		if err != nil {
			t.Fatal(err)
		}

		// Wait for the contract deployment transaction to be mined
		_, err = bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			t.Fatal(err)
		}

		// Check the contract state
		assert.Equal(t, name, token.Name())
		assert.Equal(t, symbol, token.Symbol())
		assert.Equal(t, totalSupply, token.TotalSupply())
		assert.Equal(t, totalSupply, token.BalanceOf(auth.From))

		// Transfer tokens to a new account
		to := common.HexToAddress("0x1234567890123456789012345678901234567890")
		amount := big.NewInt(100)
		_, err = token.Transfer(auth, to, amount)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, totalSupply.Sub(totalSupply, amount), token.BalanceOf(auth.From))
		assert.Equal(t, amount, token.BalanceOf(to))
	}`
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "initializing new dapp",
	Long:  `This command creates new decentralized application, which you can test and develop`,
	Run: func(cmd *cobra.Command, args []string) {
		Start()
	},
}

func createSampleContract() error {
	fi, err := os.Create("./Contracts/SampleContract.sol")
	if err != nil {
		logger.Logger().Printf("Error creating SampleContract.sol: %v\n", err)
		return err
	}
	_, err = fi.WriteString(sampleContractTemplate)
	if err != nil {
		logger.Logger().Printf("Error inserting data to sample contract: %v\n", err)
		return err
	}
	return nil
}

// CreateContractsDirectory creates directory, where all contract should be located
func CreateContractsDirectory() {
	os.Mkdir("Contracts", os.ModePerm)

	err := createSampleContract()
	if err != nil {
		logger.Logger().Println("error creating sample contract file: %v\n", err)
		return
	}
}

// CreateTestingDirectory creates directory, where all test files should be located
func CreateTestingDirectory() {
	os.Mkdir("Testing", os.ModePerm)
	createTestingFiles()
}

// CreateTestingFiles()
func createTestingFiles() ([]byte, error) {
	file, err := os.Create("Testing/SampleContract_test.go")
	if err != nil {
		logger.Logger().Printf("error creating SampleContract_test.go: %v\n", err)
		return nil, err
	}

	v, err := json.Marshal(testContractSample)
	if err != nil {
		logger.Logger().Printf("error marshalling json data: %v\n", err)
		return nil, err
	}

	_, err = file.WriteString(string(v))
	if err != nil {
		logger.Logger().Printf("error writing test cases for contract: %v\n", err)
		return nil, err
	}
	return v, nil
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.Flags().BoolP("help", "h", false, "")
}

// Start function uses to actually run a command
// it creates contract directory, and sample contract
// creates testing directory, and sample tests
func Start() {
	CreateContractsDirectory()
	CreateTestingDirectory()

	fmt.Println("Successfully created project. Wish you good luck, in creating DaPPs!!")
}
