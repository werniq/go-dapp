package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"go-solidity-gen/logger"
	"os"
)

type Contract struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Inputs []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"inputs"`
	Outputs []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"outputs"`
}

func CreateNewContractInstance(c *cobra.Command, args []string) {
	if args == nil {
		logger.Logger().Printf("error creating new contract instance: %v\n", errors.New("this command should have exactly 1 argument"))
		return
	}
	var err error
	location := fmt.Sprintf("../build/%s.sol", args[0])
	if _, err = os.Stat(fmt.Sprintf(location)); err != nil {
		logger.Logger().Printf("file %s.sol do not exist", args[0])
		return
	}
	data, err := os.ReadFile(location)
	if err != nil {
		logger.Logger().Printf("Error reading file body: %v\n", err)
		return
	}
	var contract Contract
	err = json.Unmarshal(data, &contract)
	fmt.Println(contract)
}
