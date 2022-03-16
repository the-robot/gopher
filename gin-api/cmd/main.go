package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"gingo/cmd/http"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use: "gingo",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringP("mode", "m", "dev", "Set the mode")
	http.RegisterCommandRecursive(RootCmd)
}
