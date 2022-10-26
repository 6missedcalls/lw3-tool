/*
Copyright © 2022 n HERE <EMAIL ADDRESS>

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

func Erc20Command(ctx context.Context) *cobra.Command {
	// erc20Cmd represents the erc20 command
	var erc20Cmd = &cobra.Command{
		Use:   "erc-20",
		Short: "This command will generate a SafeMintable, Enumerable, and Burnable ERC-20 contract",
		Long: `This command will generate a SafeMintable, Enumerable, and Burnable ERC-20 contract
	Usage: scaffold erc-20 -n <name> -t <ticker>`,

		Run: func(cmd *cobra.Command, args []string) {
			var (
				n   string
				t   string
				p   string
				err error
			)

			if n, err = cmd.Flags().GetString("name"); n == "" {
				n = (&prompter.Prompter{
					Message: "Enter n of Contract",
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

			if t, err = cmd.Flags().GetString("ticker"); t == "" {
				t = (&prompter.Prompter{
					Message: "Enter t of Contract",
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

			if p, err = cmd.Flags().GetString("premint"); p == "" {
				p = (&prompter.Prompter{
					Message: "Enter p of Contract",
				}).Prompt()

				if p == "" {
					// CHANGE THIS ERROR TO REFLECT PROPER ERROR
					fmt.Println(e.NilName)
					return
				}
			}
			if err != nil {
				fmt.Println(err)
				return
			}

			// This will create a new contract object
			// This object will be used to generate the contract
			contract, err := contracts.BaseErc20(n, t, p)
			if err != nil {
				fmt.Println(err)
				return
			}

			Path, err := internal.CheckPath()
			if err != nil {
				fmt.Println(err)
				return
			}

			// This will create the .sol file for the contract
			// * Path is hardcoded for now but will be changed later
			g.GenerateSol(Path+n+".sol", contract)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
	erc20Cmd.Flags().StringP("name", "n", "", "n of the contract")
	erc20Cmd.Flags().StringP("ticker", "t", "", "t of the contract")
	erc20Cmd.Flags().StringP("premint", "p", "", "Pre-minted supply of the contract")
	return erc20Cmd
}
