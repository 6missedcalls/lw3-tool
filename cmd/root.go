/*
Copyright © 2022 Ian
*/
package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
func BootstrapCmd(ctx context.Context) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "scaffold",
		Short: "A CLI tool to scaffold smart contracts",
		Long: `A CLI tool to scaffold smart contracts
		Usage: scaffold <command>`,
	}
	rootCmd.AddCommand(Erc20Command(ctx))
	rootCmd.AddCommand(Erc721Command(ctx))
	rootCmd.AddCommand(VersionCmd(ctx))
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ctx context.Context) error {
	return BootstrapCmd(ctx).Execute()
}
