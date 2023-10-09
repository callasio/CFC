package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func quitContest() {
	fmt.Println("Quit contest")
}

var quitCmd = &cobra.Command{
	Use:   "quit",
	Short: "Quit contest",
	Long:  "Quit currently entered contest.",
	Run: func(cmd *cobra.Command, args []string) {
		quitContest()
	},
}

func init() {
	quitCmd.PersistentFlags().BoolP("help", "h", false, "help for enter")
	_ = quitCmd.PersistentFlags().MarkHidden("help")
}
