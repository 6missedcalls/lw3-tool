/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/6missedcalls/lw3-tool/internal"
	"github.com/spf13/cobra"
)

func VersionCmd(ctx context.Context) *cobra.Command {
	// versionCmd represents the version command
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Here we are using the context to get the version
			v, c, d := internal.GetVersionInfo(ctx)
			gv := runtime.Version()
			os := runtime.GOOS
			arch := runtime.GOARCH

			sb := strings.Builder{}
			sb.WriteString(fmt.Sprintf("Version: %s\n", v))
			sb.WriteString(fmt.Sprintf("Commit: %s\n", c))
			sb.WriteString(fmt.Sprintf("Date: %s\n", d))
			sb.WriteString(fmt.Sprintf("Go Version: %s\n", gv))
			sb.WriteString(fmt.Sprintf("OS: %s\n", os))
			sb.WriteString(fmt.Sprintf("Arch: %s\n", arch))
			fmt.Print(sb.String())
		},
	}
	return versionCmd
}
