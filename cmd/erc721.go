/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/6missedcalls/lw3-tool/contracts"
	"github.com/6missedcalls/lw3-tool/internal/generate"

	"github.com/Songmu/prompter"
	"github.com/spf13/cobra"
)

// erc721Cmd represents the erc721 command
var erc721Cmd = &cobra.Command{
	Use:   "erc-721",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// take the BasicContract Method from Contracts
		// and print it to the console

		// here we will prompt the user for some data then call the method to create the contract
		n := prompter.Prompt("Name of the contract", "CLI")    // prompt the user for the name of the contract
		t := prompter.Prompt("Ticker of the contract", "TOOL") // prompt the user for the ticker of the contract

		// call the method to create the contract
		// then print it to the console
		c, err := contracts.BaseErc721(n, t)
		if err != nil {
			fmt.Println(err)
		}

		// write the contract to the file
		err = generate.GenerateSol("erc721.sol", c)

		// write to the file
		fmt.Println("Contract written to file")
	},
}

func init() {
	rootCmd.AddCommand(erc721Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// erc721Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// erc721Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
