/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"

	"github.com/6missedcalls/lw3-tool/contracts"
	"github.com/6missedcalls/lw3-tool/internal"
	e "github.com/6missedcalls/lw3-tool/internal/errors"
	g "github.com/6missedcalls/lw3-tool/internal/generate"

	"github.com/Songmu/prompter"
	"github.com/spf13/cobra"
)

func Erc721Command(ctx context.Context) *cobra.Command {
	var erc721Cmd = &cobra.Command{
		Use:   "erc-721",
		Short: "This command will generate a SafeMintable, Enumerable, and Burnable ERC-721 contract",
		Long: `This command will generate a SafeMintable, Enumerable, and Burnable ERC-721 contract
		Usage: scaffold erc-721 -n <name> -t <ticker>`,
		Run: func(cmd *cobra.Command, args []string) {
			// Variables used to store the name and ticker of the contract
			var (
				n   string // Name of the contract
				t   string // Ticker of the contract
				err error  // Error variable
			)

			// This will first check to see if the name flag was used
			// If it was not used, it will prompt the user for the name
			// If it was used, it will store the value of the flag in the n variable
			if n, err = cmd.Flags().GetString("name"); n == "" {
				n = (&prompter.Prompter{
					Message: "Enter Name of Contract",
				}).Prompt()

				if n == "" {
					fmt.Println(e.NilName)
					return
				}
			}
			if err != nil {
				fmt.Println(err)
				return
			}

			// This will first check to see if the ticker flag was used
			// If it was not used, it will prompt the user for the ticker
			// If it was used, it will store the value of the flag in the t variable
			if t, err = cmd.Flags().GetString("ticker"); t == "" {
				t = (&prompter.Prompter{
					Message: "Enter Ticker of Contract",
				}).Prompt()

				if t == "" {
					fmt.Println(e.NilTicker)
					return
				}
			}
			if err != nil {
				fmt.Println(err)
				return
			}

			// This will generate the erc721 contract using the name and ticker
			c, err := contracts.BaseErc721(n, t)
			if err != nil {
				fmt.Println(err)
			}

			Path, err := internal.CheckPath()
			if err != nil {
				fmt.Println(err)
				return
			}

			// This will create the .sol file for the contract
			// * Path is hardcoded for now but will be changed later
			g.GenerateSol(Path+n+".sol", c)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
	erc721Cmd.Flags().StringP("name", "n", "", "Name of the contract")
	erc721Cmd.Flags().StringP("ticker", "t", "", "Ticker of the contract")
	return erc721Cmd
}
