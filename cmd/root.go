package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cfc",
		Short: "A command line interface for Codeforces.",
		Long: `CFC is a CLI written in Go and Python.
This application is a tool to automate testing and submission for Codeforces.`,
		Run: func(cmd *cobra.Command, args []string) {},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(enterCmd)
	rootCmd.AddCommand(quitCmd)
	rootCmd.AddCommand(submitCmd)
	rootCmd.AddCommand(testCmd)
}
